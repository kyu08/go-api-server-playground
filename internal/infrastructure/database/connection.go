package database

import (
	"context"
	_ "embed"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/apstndb/spanemuboost"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

//go:embed schema/schema.sql
var schemaDDL string

// NewEmulatorWithClient starts a Spanner emulator and returns a client connected to it.
// The returned teardown function should be called when the client is no longer needed.
func NewEmulatorWithClient(ctx context.Context) (*spanner.Client, func(), error) {
	ddls := parseDDLs(schemaDDL)

	_, clients, teardown, err := spanemuboost.NewEmulatorWithClients(ctx,
		spanemuboost.WithSetupDDLs(ddls),
	)
	if err != nil {
		return nil, nil, errors.WithStack(errors.NewInternalError(err))
	}

	return clients.Client, teardown, nil
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
