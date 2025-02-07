package rediso

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/k0kubun/pp"
	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
)

type (
	storeAndDeleteCache struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func (p *storeAndDeleteCache) StorePost(ctx context.Context, post *posts.Post) error {
	byteData, err := json.Marshal(post)
	if err != nil {
		p.logger.Println("-- ERROR WHILE MARSHALING DATA IN StorePost SERVICE --")
		return err
	}
	if err := p.redisDb.Set(ctx, post.PostId, byteData, time.Hour*24).Err(); err != nil {
		p.logger.Println(err)
		pp.Println(err, "--------------------------------------------------------")
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
