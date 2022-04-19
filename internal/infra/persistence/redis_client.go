package persistence

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

type redisClient struct {
	network string //"tcp"
	address string //"redis:6379"
	options []redis.DialOption

	conn      redis.Conn
	connected bool
}

func NewRedisClient(network string, address string, options ...redis.DialOption) KeyValueDBClient {

	return &redisClient{
		network: network,
		address: address,
		options: options}
}

func (r *redisClient) Connect() error {
	conn, err := redis.Dial(r.network, r.address)
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r *redisClient) Close() {
	r.conn.Close()
}

func (r *redisClient) Set(key string, value []byte, expireInSec int) error {

	if err := r.Connect(); err != nil {
		return err
	}
	defer r.conn.Close()

	_, err := r.conn.Do("SET", key, []byte(value))
	if err != nil {
		return err
	}

	if expireInSec > 0 {
		r.conn.Do("EXPIRE", key, expireInSec)
	}

	return err
}

func (r *redisClient) Get(key string) ([]byte, error) {

	var data []byte

	if err := r.Connect(); err != nil {
		return data, err
	}
	defer r.conn.Close()

	data, err := redis.Bytes(r.conn.Do("GET", key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return data, err
	}

	return data, nil
}

func (r *redisClient) Flush(key string) ([]byte, error) {

	var data []byte

	if err := r.Connect(); err != nil {
		return data, err
	}

	defer r.conn.Close()

	data, err := redis.Bytes(r.conn.Do("DEL", key))
	if err != nil {
		return data, err
	}

	return data, nil
}
