package persistence

type KeyValueDBClient interface {
	Connect() error
	Close()
	Set(key string, value []byte, expireInSec int) error
	Get(key string) ([]byte, error)
}
