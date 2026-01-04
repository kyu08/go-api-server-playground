package queryimpl

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/model"
	"github.com/kyu08/go-api-server-playground/internal/query"
	"github.com/samber/lo"
)

type UserQuery struct{}

func NewUserQuery() query.UserQuery {
	return &UserQuery{}
}

func (UserQuery) FindByScreenName(
	ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName,
) (*query.User, error) {
	u, err := model.FindUserByScreenName(ctx, rtx, screenName.String())
	if err != nil {
		if database.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return lo.ToPtr(query.User(*u)), nil
}

func (UserQuery) FindByID(
	ctx context.Context, rtx domain.ReadOnlyDB, userID domain.ID[user.User],
) (*query.User, error) {
	u, err := model.FindUser(ctx, rtx, userID.String())
	if err != nil {
		if database.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return lo.ToPtr(query.User(*u)), nil
}

func (UserQuery) ExistsByScreenName(
	ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName,
) (bool, error) {
	if _, err := model.FindUserByScreenName(ctx, rtx, screenName.String()); err != nil {
		if database.IsNotFound(err) {
			return false, nil
		}
		return false, apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return true, nil
}
