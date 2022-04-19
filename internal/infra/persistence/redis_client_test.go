package persistence_test

import (
	"testing"
	"time"

	"github.com/fgmaia/chat/internal/infra/persistence"
	"github.com/stretchr/testify/require"
)

func TestRedisClient(t *testing.T) {
	t.Parallel()

	t.Run("when set and get success", func(t *testing.T) {
		t.Parallel()

		client := persistence.NewRedisClient("tcp", "localhost:6379")
		err := client.Connect()
		require.NoError(t, err)

		err = client.Set("Key1", []byte("1234"), 0)
		require.NoError(t, err)

		data, err := client.Get("Key1")
		require.NoError(t, err)
		require.Equal(t, data, []byte("1234"))
	})

	t.Run("when set expire success", func(t *testing.T) {
		t.Parallel()

		client := persistence.NewRedisClient("tcp", "localhost:6379")
		err := client.Connect()
		require.NoError(t, err)

		err = client.Set("Key2", []byte("12345"), 1)
		require.NoError(t, err)

		data, err := client.Get("Key2")
		require.NoError(t, err)
		require.Equal(t, data, []byte("12345"))

		time.Sleep(time.Second * 1)

		data, err = client.Get("Key2")
		require.NoError(t, err)
		require.Nil(t, data)
	})

}
