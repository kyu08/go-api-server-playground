package database

import (
	"context"
	_ "embed"
	"errors"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/apstndb/spanemuboost"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	tcspanner "github.com/testcontainers/testcontainers-go/modules/gcloud/spanner"
)

//go:embed schema/schema.sql
var schemaDDL string
var ErrSpannerEmulatorNotInitialized = errors.New("spanner emulator is not initialized")

func GetSpannerClient(spannerEmulator *tcspanner.Container) (*spanner.Client, func(), error) {
	if spannerEmulator == nil {
		return nil, nil, apperrors.NewInternalError(ErrSpannerEmulatorNotInitialized)
	}

	clients, clientsTeardown, err := spanemuboost.NewClients(context.Background(), spannerEmulator,
		spanemuboost.EnableDatabaseAutoConfigOnly(),
		spanemuboost.WithRandomDatabaseID(),
		spanemuboost.WithSetupDDLs(parseDDLs(schemaDDL)),
		spanemuboost.WithSetupDMLs([]spanner.Statement{}), // 必要になったら外から受け取るなどする。
	)
	if err != nil {
		return nil, nil, apperrors.WithStack(err)
	}

	return clients.Client, clientsTeardown, nil
}

// parseDDLs splits a DDL string into individual statements.
func parseDDLs(ddl string) []string {
	statements := strings.Split(ddl, ";")

	var result []string

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			result = append(result, stmt)
		}
	}

	return result
}
