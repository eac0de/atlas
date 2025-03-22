package storage

import (
	"context"

	"github.com/eac0de/atlas/shared/pkg/psql"
)

type NexusStorage struct {
	*psql.PSQLStorage
}

func NewNexusStorage(
	ctx context.Context,
	host string,
	port string,
	username string,
	password string,
	dbName string,
) (*NexusStorage, error) {
	storage, err := psql.New(ctx, host, port, username, password, dbName)
	if err != nil {
		return nil, err
	}
	return &NexusStorage{storage}, nil
}
