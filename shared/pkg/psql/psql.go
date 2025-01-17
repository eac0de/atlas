package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type PSQLStorage struct {
	*pgxpool.Pool
}

func New(
	ctx context.Context,
	host string,
	port string,
	username string,
	password string,
	dbName string,
) (*PSQLStorage, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	storage := PSQLStorage{
		pool,
	}
	err = storage.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &storage, nil
}

func (storage *PSQLStorage) Migrate(ctx context.Context, migrationsDir string, withReset bool) error {
	sqlDB := stdlib.OpenDB(*storage.Config().ConnConfig)
	defer sqlDB.Close()
	if withReset {
		if err := goose.ResetContext(ctx, sqlDB, migrationsDir); err != nil {
			return err
		}
	}
	if err := goose.UpContext(ctx, sqlDB, migrationsDir); err != nil {
		return err
	}
	return nil
}
