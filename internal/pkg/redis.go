package redisCl

import (
	"github.com/go-redis/redis/v8"
	"github.com/k0kubun/pp"
	"github.com/ruziba3vich/prosphere/internal/items/config"
)

func NewRedisDB(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: "",
		DB:       0,
	})
	pp.Println(rdb)
	return rdb, nil
}
