package commandimpl

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/model"
)

type UserRepository struct{}

func NewUserRepository() user.UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Create(ctx context.Context, rwtx domain.ReadWriteDB, u *user.User) error {
	return r.apply(rwtx, []*spanner.Mutation{r.fromDomain(u).Insert(ctx)})
}

func (r UserRepository) FindByID(
	ctx context.Context, rtx domain.ReadOnlyDB, userID domain.ID[user.User],
) (*user.User, error) {
	u, err := model.FindUser(ctx, rtx, userID.String())
	if err != nil {
		if database.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return r.toDomain(u)
}

func (r UserRepository) FindByScreenName(
	ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName,
) (*user.User, error) {
	u, err := model.FindUserByScreenName(ctx, rtx, screenName.String())
	if err != nil {
		if database.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return r.toDomain(u)
}

func (UserRepository) apply(rwtx domain.ReadWriteDB, m []*spanner.Mutation) error {
	if err := rwtx.BufferWrite(m); err != nil {
		return apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return nil
}

func (UserRepository) fromDomain(u *user.User) *model.User {
	return &model.User{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName().String(),
		UserName:   u.UserName().String(),
		Bio:        u.Bio().String(),
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (UserRepository) toDomain(dto *model.User) (*user.User, error) {
	u, err := user.NewFromDTO(
		dto.ID,
		dto.ScreenName,
		dto.UserName,
		dto.Bio,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}
	return u, nil
}
