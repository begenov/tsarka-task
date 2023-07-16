package redis

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
)

func TestCouterRepo_Add(t *testing.T) {
	s, err := miniredis.Run()
	require.NoError(t, err)

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	repo := NewCouterRepo(client)

	result, err := repo.Add("counter", 5)

	require.NoError(t, err)
	require.Equal(t, int64(5), result)

	client.Close()
	s.Close()
}

func TestCouterRepo_Sub(t *testing.T) {
	s, err := miniredis.Run()
	require.NoError(t, err)

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	repo := NewCouterRepo(client)

	err = client.Set("counter", "10", 0).Err()
	require.NoError(t, err)

	result, err := repo.Sub("counter", 3)

	require.NoError(t, err)
	require.Equal(t, int64(7), result)

	client.Close()
	s.Close()
}

func TestCouterRepo_Get(t *testing.T) {

	s, err := miniredis.Run()
	require.NoError(t, err)

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	repo := NewCouterRepo(client)

	err = client.Set("counter", "15", 0).Err()
	require.NoError(t, err)

	result, err := repo.Get("counter")

	require.NoError(t, err)
	require.Equal(t, int64(15), result)

	client.Close()
	s.Close()
}
