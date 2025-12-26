package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Create(ctx context.Context, tx *spanner.ReadWriteTransaction, u *user.User) error {
	m, err := spanner.InsertStruct("User", &database.User{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	})
	if err != nil {
		return errors.WithStack(errors.NewInternalError(err))
	}

	if err := tx.BufferWrite([]*spanner.Mutation{m}); err != nil {
		return errors.WithStack(errors.NewInternalError(err))
	}

	return nil
}

func (r UserRepository) FindByScreenName(
	ctx context.Context, tx *spanner.ReadWriteTransaction, screenName user.ScreenName,
) (*user.User, error) {
	stmt := spanner.Statement{
		SQL: `SELECT ID, ScreenName, UserName, Bio, IsPrivate, CreatedAt FROM User WHERE ScreenName = @screenName LIMIT 1`,
		Params: map[string]any{
			"screenName": string(screenName),
		},
	}

	iter := tx.Query(ctx, stmt)
	defer iter.Stop()

	row, err := iter.Next()
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return nil, errors.WithStack(errors.NewNotFoundError("user"))
		}

		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	u, err := database.UserFromRow(row)
	if err != nil {
		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return u.ToUser(), nil
}

func (r UserRepository) ExistsByScreenName(
	ctx context.Context, tx *spanner.ReadWriteTransaction, screenName user.ScreenName,
) (bool, error) {
	stmt := spanner.Statement{
		SQL: `SELECT 1 FROM User WHERE ScreenName = @screenName LIMIT 1`,
		Params: map[string]any{
			"screenName": string(screenName),
		},
	}

	iter := tx.Query(ctx, stmt)
	defer iter.Stop()

	_, err := iter.Next()
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return false, nil
		}

		return false, errors.WithStack(errors.NewInternalError(err))
	}

	return true, nil
}
