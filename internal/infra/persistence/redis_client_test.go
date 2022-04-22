package persistence_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fgmaia/chat/internal/infra/persistence"
	"github.com/stretchr/testify/require"
)

func TestRedisClient(t *testing.T) {
	t.Parallel()

	t.Run("when set and get success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := persistence.NewRedisClient("localhost:6379", "")

		err := client.Set(ctx, "Key1", "1234", persistence.KEEP_TTL)
		require.NoError(t, err)

		data, err := client.Get(ctx, "Key1")
		require.NoError(t, err)
		require.Equal(t, data, "1234")
	})

	t.Run("when set expire success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := persistence.NewRedisClient("localhost:6379", "")

		err := client.Set(ctx, "Key2", "12345", time.Second*2)
		require.NoError(t, err)

		data, err := client.Get(ctx, "Key2")
		require.NoError(t, err)
		require.Equal(t, data, "12345")

		time.Sleep(time.Second * 2)

		data, err = client.Get(ctx, "Key2")
		require.Error(t, err)
		require.Empty(t, data)
	})

	t.Run("when scan success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := persistence.NewRedisClient("localhost:6379", "")

		for i := 1; i <= 20; i++ {
			err := client.Set(ctx, fmt.Sprintf("message-%d", i), fmt.Sprintf("msg%d", i), persistence.KEEP_TTL)
			require.NoError(t, err)
		}

		data, n, err := client.Scan(ctx, 0, "message-*", 20)
		if len(data) < 20 {
			data1, _, err := client.Scan(ctx, n, "message-*", 20)
			require.NoError(t, err)
			data = append(data, data1...)
		}
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(data), 20)
	})

}
