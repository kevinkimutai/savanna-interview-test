package ports

import (
	"context"

	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
)

type DBTX interface {
	ExecTx(ctx context.Context, fn func(*queries.Queries) error) error
}
