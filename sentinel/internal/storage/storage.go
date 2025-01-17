package storage

import (
	"context"

	"github.com/eac0de/gophkeeper/shared/pkg/psql"
)

type AuthStorage struct {
	*psql.PSQLStorage
}

func NewAuthStorage(
	ctx context.Context,
	host string,
	port string,
	username string,
	password string,
	dbName string,
) (*AuthStorage, error) {
	storage, err := psql.New(ctx, host, port, username, password, dbName)
	if err != nil {
		return nil, err
	}
	return &AuthStorage{storage}, nil
}
