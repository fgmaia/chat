package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

const (
	KEEP_TTL = time.Duration(-1)
)

type redisClient struct {
	address  string //"localhost:6379"
	password string

	client       *redis.Client
	jsonHandler  *rejson.Handler
	searchClient *redisearch.Client
}

func NewRedisClient(address string, password string) KeyValueDBClient {
	c := &redisClient{
		address:  address,
		password: password,
	}
	c.jsonHandler = rejson.NewReJSONHandler()
	cli := redis.NewClient(&redis.Options{Addr: address})
	c.jsonHandler.SetGoRedisClient(cli)

	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	c.searchClient = redisearch.NewClient(address, "myIndex")

	return c
}

func (r *redisClient) JSONSet(key string, path string, obj interface{}) error {

	if path == "" {
		path = "."
	}

	res, err := r.jsonHandler.JSONSet(key, ".", obj)
	if err != nil {
		return err
	}

	if res.(string) != "OK" {
		return errors.New("failed to set")
	}

	return nil
}

func (r *redisClient) JSONGet(key string, path string, obj interface{}) error {

	if path == "" {
		path = "."
	}

	dataJSON, err := redigo.Bytes(r.jsonHandler.JSONGet(key, "."))
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataJSON, obj)
	return err
}

func (r *redisClient) Close() error {
	return r.client.Close()
}

func (r *redisClient) FlushAll(ctx context.Context) error {
	statusCmd := r.client.FlushAll(ctx)
	return statusCmd.Err()
}

func (r *redisClient) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	statusCmd := r.client.Set(ctx, key, value, expiration)
	return statusCmd.Err()
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	stringCmd := r.client.Get(ctx, key)
	if stringCmd == nil {
		return "", nil
	}
	return stringCmd.Result()
}

func (r *redisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.Scan(ctx, cursor, match, count).Result()
}

func (r *redisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.HScan(ctx, key, cursor, match, count).Result()
}
