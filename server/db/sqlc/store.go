package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	WordQuery
}
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
