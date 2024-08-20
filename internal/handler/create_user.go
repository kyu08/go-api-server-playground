package handler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

var ErrCreateUserScreenNameAlreadyUsed = errors.New("the screen name specified is already used")

func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.Printf("Received: %v", "CreateUser") // TODO: インターセプター側でログ出力するようにする

	newUser, err := user.New(req.GetScreenName(), req.GetUserName(), req.GetBio())
	if err != nil {
		return nil, fmt.Errorf("user.NewUser: %w", err)
	}

	if err := database.WithTransaction(ctx, s.db, func(queries *database.Queries) error {
		if u, err := s.userRepository.FindByScreenName(ctx, queries, newUser.ScreenName); u != nil {
			return ErrCreateUserScreenNameAlreadyUsed
		} else if err != nil && !errors.Is(err, repository.ErrFindUserByScreenNameUserNotFound) {
			return fmt.Errorf("s.userRepository.FindByScreenName: %w", err) // TODO: 重複チェックは誰がやるべきなんだ... domainService的な存在が必要かも？
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
