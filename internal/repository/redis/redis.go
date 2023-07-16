package redis

import "github.com/go-redis/redis"

type CouterRepo struct {
	redisClient *redis.Client
}

func NewCouterRepo(redisClient *redis.Client) *CouterRepo {
	return &CouterRepo{
		redisClient: redisClient,
	}
}

func (r *CouterRepo) Add(key string, value int64) (int64, error) {
	return r.redisClient.IncrBy(key, value).Result()
}

func (r *CouterRepo) Sub(key string, value int64) (int64, error) {
	return r.redisClient.DecrBy(key, value).Result()
}

func (r *CouterRepo) Get(key string) (int64, error) {
	return r.redisClient.Get(key).Int64()
}
