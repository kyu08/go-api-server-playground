package infrastructure

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/query"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure"
	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure/database/dao"
	"github.com/samber/lo"
)

type UserQuery struct{}

func NewUserQuery() query.UserQuery {
	return &UserQuery{}
}

func (UserQuery) FindByScreenName(
	ctx context.Context, rtx infrastructure.ReadOnlyDB, screenName string,
) (*query.User, error) {
	u, err := dao.FindUserByScreenName(ctx, rtx, screenName)
	if err != nil {
		if dao.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return lo.ToPtr(query.User(*u)), nil
}

func (UserQuery) FindByID(
	ctx context.Context, rtx infrastructure.ReadOnlyDB, userID string,
) (*query.User, error) {
	u, err := dao.FindUser(ctx, rtx, userID)
	if err != nil {
		if dao.IsNotFound(err) {
			return nil, apperrors.WithStack(apperrors.NewNotFoundError("user"))
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return lo.ToPtr(query.User(*u)), nil
}
