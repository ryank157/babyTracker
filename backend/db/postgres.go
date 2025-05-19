package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	db *pgxpool.Pool
}

func NewPostgresDB(connectionString string) (*PostgresDB, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDB{db: dbpool}, nil
}

func (p *PostgresDB) Close() error {
	p.db.Close()
	return nil
}

// ExecContext executes a query without returning any rows.
func (p *PostgresDB) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return p.db.Exec(ctx, query, args...)
}

func (p *PostgresDB) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return p.db.Query(ctx, query, args...)
}

func (p *PostgresDB) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return p.db.QueryRow(ctx, query, args...)
}

func (p *PostgresDB) Queries() *Queries {
	return New(p)
}

func (p *PostgresDB) Ping(ctx context.Context) error {
	return p.db.Ping(ctx)
}
