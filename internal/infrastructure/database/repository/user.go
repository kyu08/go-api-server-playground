package repository

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
)

type UserRepository struct {
	queries *database.Queries
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		queries: nil,
	}
}

func (r *UserRepository) SetQueries(q *database.Queries) {
	r.queries = q
}

func (r UserRepository) Create(ctx context.Context, u *user.User) error {
	p := database.CreateUserParams{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}
	if _, err := r.queries.CreateUser(ctx, p); err != nil {
		// FIXME: 正確にはID, ScreenNameの重複エラーが返ってくる場合もあり、それらの場合はPreconditionErrorにすべきだがサボってInternalにしている
		return errors.WithStack(errors.NewInternalError(err))
	}

	return nil
}

func (r UserRepository) FindByScreenName(ctx context.Context, screenName user.ScreenName,
) (*user.User, error) {
	u, err := r.queries.FindUserByScreenName(ctx, string(screenName))
	if err != nil {
		if database.IsNotFoundFromDB(err) {
			return nil, errors.WithStack(errors.NewNotFoundError("user"))
		}

		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return u.ToUser(), nil
}
