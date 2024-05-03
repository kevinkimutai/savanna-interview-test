package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
)

type DBAdapter struct {
	ctx     context.Context
	queries *queries.Queries
}

func NewDB(DBUrl string) *DBAdapter {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, DBUrl)
	if err != nil {
		log.Fatal("Error connecting to db:%w", err)
	}
	// defer conn.Close(ctx)

	queries := queries.New(conn)

	return &DBAdapter{ctx: ctx, queries: queries}

}
