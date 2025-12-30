package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type UserRepository struct{}

func NewUserRepository() user.UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Create(ctx context.Context, rwtx domain.ReadWriteDB, u *user.User) error {
	return r.apply(rwtx, []*spanner.Mutation{r.fromDomain(u).Insert(ctx)})
}

func (r UserRepository) FindByScreenName(
	ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName,
) (*user.User, error) {
	u, err := FindUserByScreenName(ctx, rtx, screenName.String())
	if err != nil {
		if IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return r.toDomain(u)
}

func (r UserRepository) ExistsByScreenName(
	ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName,
) (bool, error) {
	if _, err := FindUserByScreenName(ctx, rtx, screenName.String()); err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return true, nil
}

func (UserRepository) apply(rwtx domain.ReadWriteDB, m []*spanner.Mutation) error {
	if err := rwtx.BufferWrite(m); err != nil {
		return apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return nil
}

func (UserRepository) fromDomain(u *user.User) *User {
	return &User{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName().String(),
		UserName:   u.UserName().String(),
		Bio:        u.Bio().String(),
		CreatedAt:  u.CreatedAt,
	}
}

func (UserRepository) toDomain(dto *User) (*user.User, error) {
	u, err := user.NewFromDTO(dto.ID, dto.ScreenName, dto.UserName, dto.Bio, dto.CreatedAt)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}
	return u, nil
}
