package persistence

import (
	"context"
	"time"
)

type KeyValueDBClient interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)
	JSONSet(key string, path string, obj interface{}) error
	JSONGet(key string, path string, obj interface{}) error
	FlushAll(ctx context.Context) error
	Close() error
}
