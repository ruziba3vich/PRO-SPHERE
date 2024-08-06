package rediso

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	storeAndDeleteCache struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func (p *storeAndDeleteCache) Store(ctx context.Context, uniqueKey string, data []byte) error {
	if err := p.redisDb.Set(ctx, uniqueKey, data, time.Hour*24).Err(); err != nil {
		p.logger.Println(err)
		return err
	}
	return nil
}

func (p *storeAndDeleteCache) Delete(ctx context.Context, uniqueKey string) error {
	if err := p.redisDb.Del(ctx, uniqueKey).Err(); err != nil {
		p.logger.Println(err)
		return err
	}
	return nil
}
