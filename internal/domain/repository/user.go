package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
)

// TODO: domain層がgo/spannerに依存している。
// これを許容するならその旨をドキュメントに残す。
// あるいは依存しない方法を考える。（たぶんtx用のI/Fを切るくらいしかない気はしている）
type UserRepository interface {
	Create(ctx context.Context, tx *spanner.ReadWriteTransaction, u *user.User) error
	FindByScreenName(ctx context.Context, tx *spanner.ReadWriteTransaction, screenName user.ScreenName) (*user.User, error)
}
