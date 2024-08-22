package handler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

var ErrCreateUserScreenNameAlreadyUsed = errors.New("the screen name specified is already used")

// TODO: usecase層に分離する
// TODO: usecase層がscreenNameのユニークチェックを行うパターンとdomain serviceがそれを行うパターンの両方を実装してみる

func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.Printf("Received: %v", "CreateUser") // TODO: インターセプター側でログ出力するようにする

	newUser, err := user.New(req.GetScreenName(), req.GetUserName(), req.GetBio())
	if err != nil {
		return nil, fmt.Errorf("user.NewUser: %w", err)
	}

	if err := database.WithTransaction(ctx, s.db, func(queries *database.Queries) error {
		// TODO:そもそもuserService.Create()を呼び出すだけでいいにしたくない？
		// けどそうするとだいぶserviceの責務が増えるかも？
		userService := user.UserService{ // TODO: 別の場所で初期化 or TwitterServerのメソッドとして提供すべき?
			UserRepository: queries,
		}
		isUnique, err := userService.IsUniqueScreenName(ctx, newUser.ScreenName)
		if err != nil {
			return fmt.Errorf("userService.IsUniqueScreenName: %w", err)
		}
		if !isUnique {
			return ErrCreateUserScreenNameAlreadyUsed
		}

		if err := s.userRepository.Create(ctx, queries, newUser); err != nil {
			return fmt.Errorf("s.userRepository.Create: %w", err)
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("database.WithTransaction: %w", err)
	}

	return &api.CreateUserResponse{}, nil
}
