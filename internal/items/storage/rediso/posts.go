package rediso

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
)

type (
	PostsCache struct {
		redisDb *redis.Client
		logger  *log.Logger
		storeAndDeleteCache
	}
)

func NewPostCache(redisDb *redis.Client, logger *log.Logger) *PostsCache {
	return &PostsCache{
		redisDb: redisDb,
		logger:  logger,
	}
}

func (p *PostsCache) StorePost(ctx context.Context, post *posts.Post) error {
	byteData, err := json.Marshal(post)
	if err != nil {
		p.logger.Println("-- ERROR OCCURED WHILE MARSHALING THE DATA IN StorePost SERVICE --")
		return err
	}
	if err := p.redisDb.Set(ctx, post.PostId, byteData, time.Hour*24).Err(); err != nil {
		p.logger.Println("-- ERROR RETURNED FROM StorePost SERVICE --")
		return err
	}
	return nil
}

func (p *PostsCache) DeletePost(ctx context.Context, req *posts.DeletePostRequest) error {
	result, err := p.redisDb.Del(ctx, req.PostId).Result()
	if err != nil {
		return err
	}

	if result == 0 {
		p.logger.Printf("post with ID %s does not exist in Redis", req.PostId)
	} else {
		p.logger.Printf("post with ID %s has been deleted from Redis", req.PostId)
	}

	return nil
}

func (p *PostsCache) GetPostById(ctx context.Context, req *posts.GetPostByIdRequest) (*posts.Post, error) {
	byteData, err := p.redisDb.Get(ctx, req.PostId).Result()
	if err != nil {
		if err == redis.Nil {
			p.logger.Println("-- DATA NOT FOUND FROM REDIS --")
			return nil, err
		}
		p.logger.Println("-- ERROR OCCURED WHILE GETTING DATA FROM REDIS IN GetPostById SERVICE", err.Error())
		return nil, err
	}
	var response posts.Post
	if err := json.Unmarshal([]byte(byteData), &response); err != nil {
		p.logger.Println("-- ERROR OCCURED WHILE UNMARSHALING DATA IN GetPostById SERVICE", err.Error())
		return nil, err
	}
	p.logger.Println("-- GOT FROM REDIS --")
	return &response, nil
}
