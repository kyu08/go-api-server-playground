package repository

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

var ErrFindUserByScreenNameUserNotFound = errors.NewPreconditionError("user not found")

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Create(ctx context.Context, queries *database.Queries, u *user.User) error {
	p := database.CreateUserParams{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}
	if _, err := queries.CreateUser(ctx, p); err != nil {
		// FIXME: 正確にはID, ScreenNameの重複エラーが返ってくる場合もあり、それらの場合はPreconditionErrorにすべきだがサボってInternalにしている
		return errors.WithStack(errors.NewInternalError(err))
	}

	return nil
}

func (UserRepository) FindByScreenName(
	ctx context.Context,
	queries *database.Queries,
	screenName user.ScreenName,
) (*user.User, error) {
	u, err := queries.FindUserByScreenName(ctx, string(screenName))
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return nil, errors.WithStack(ErrFindUserByScreenNameUserNotFound)
		}

		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return u.ToUser(), nil
}
