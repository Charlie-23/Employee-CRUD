package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateEmployeeTx(ctx context.Context, arg CreateEmployeeTxParams) (CreateEmployeeTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new storevscode-remote://wsl%2Bubuntu-20.04/home/pzanwar/Go/projects/temp/simplebank/db/sqlc/store_test.go
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
