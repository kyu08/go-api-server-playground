package domain

import (
	"context"

	"cloud.google.com/go/spanner"
)

// ドメイン層がspannerに依存してしまっており、方針が詳細に依存してしまい本来的には好ましくない。
// しかし依存を完全に排除するいい方法も思いついていないためinterfaceを切って対応する。
//
// spanner固有の構造体をrepositoryの引数として受け取るよりはinterfaceを切っておいた方が
// 後に別DBに移行する場合などでも多少柔軟性が上がると考えてこの形にしている。
type (
	ReadWriteDB interface {
		BufferWrite(ms []*spanner.Mutation) error
	}
	// TODO: cqrs化が完了したらquery packageに移動できるはず
	ReadOnlyDB interface {
		ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
		Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
		ReadUsingIndex(ctx context.Context, table, index string, keys spanner.KeySet, columns []string) (ri *spanner.RowIterator)
		Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
	}
)
