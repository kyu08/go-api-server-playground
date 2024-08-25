package usecase

import (
	"context"
	stderrors "errors"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

// 複数のusecaseから呼ばれうる関数はこのファイルのように[entity_name.go]のような形式のファイルに置く

type userHelper_ struct{}

var userHelper userHelper_

func (userHelper_) isUniqueScreenName(
	ctx context.Context,
	userRepository *repository.UserRepository,
	queries *database.Queries,
	screenName user.ScreenName,
) (bool, error) {
	user, err := userRepository.FindByScreenName(ctx, queries, screenName)
	if err != nil && !stderrors.Is(err, repository.ErrFindUserByScreenNameUserNotFound) { // TODO: 対象のユーザーが存在するかどうかのboolを返すメソッドの方が適切
		return false, errors.WithStack(err)
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}
