package usecase

import (
	"context"
	"fmt"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

// 複数のusecaseから呼ばれうる関数はこのファイルのように[entity_name.go]のような形式のファイルに置く

func isUniqueScreenName(
	ctx context.Context,
	userRepository *repository.UserRepository,
	queries *database.Queries,
	screenName user.ScreenName,
) (bool, error) {
	user, err := userRepository.FindByScreenName(ctx, queries, screenName)
	if err != nil && !database.IsNotFoundError(err) {
		return false, fmt.Errorf("userRepository.FindUserByScreenName: %w", err)
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}
