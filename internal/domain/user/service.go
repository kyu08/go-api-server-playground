package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// TODO: repositoryのinterfaceを再度作成する
// TODO: 重複チェックを行う database層やdatabase/sqlパッケージに依存せずに普通のエラーとnotFoundエラーの区別できる...?
// ↑そのチェックもinterfaceに含めればいけるか！

type UserService struct {
	UserRepository UserRepository
}

func (u UserService) IsUniqueScreenName(ctx context.Context, screenName ScreenName) (bool, error) {
	user, err := u.UserRepository.FindUserByScreenName_(ctx, screenName.String())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("u.UserRepository.FindUserByScreenName: %w", err)
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}
