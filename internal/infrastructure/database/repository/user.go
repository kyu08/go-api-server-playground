package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/repository"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
)

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Create(ctx context.Context, tx domain.ReadWriteDB, u *user.User) error {
	m, err := spanner.InsertStruct("User", &database.User{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	})
	if err != nil {
		return apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return r.apply(tx, []*spanner.Mutation{m})
}

func (r UserRepository) FindByScreenName(
	ctx context.Context, tx domain.ReadOnlyDB, screenName user.ScreenName,
) (*user.User, error) {
	s := spanner.NewStatement(`
	SELECT ID, ScreenName, UserName, Bio, IsPrivate, CreatedAt FROM User WHERE ScreenName = @screenName LIMIT 1
	`)
	s.Params["screenName"] = string(screenName)

	iter := tx.Query(ctx, s)
	defer iter.Stop()

	row, err := iter.Next()
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	u, err := database.UserFromRow(row)
	if err != nil {
		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return u.ToUser(), nil
}

func (r UserRepository) ExistsByScreenName(
	ctx context.Context, tx *spanner.ReadWriteTransaction, screenName user.ScreenName,
) (bool, error) {
	s := spanner.NewStatement(`
	SELECT 1 FROM User WHERE ScreenName = @screenName LIMIT 1`)
	s.Params["screenName"] = string(screenName)

	iter := tx.Query(ctx, s)
	defer iter.Stop()

	_, err := iter.Next()
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return false, nil
		}

		return false, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return true, nil
}

func (UserRepository) apply(tx domain.ReadWriteDB, m []*spanner.Mutation) error {
	if err := tx.BufferWrite(m); err != nil {
		return apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return nil
}
