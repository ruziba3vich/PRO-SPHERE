package rediso

import (
	"context"
	"encoding/json"
	"log"

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

func (p *PostsCache) StorePost(ctx context.Context, post *posts.Post) (*posts.Post, error) {
	byteData, err := json.Marshal(post)
	if err != nil {
		p.logger.Println("-- ERROR OCCURED WHILE MARSHALING THE DATA IN StorePost SERVICE --")
		return nil, err
	}
	if err := p.Store(ctx, post.PostId, byteData); err != nil {
		p.logger.Println("-- ERROR RETURNED FROM StorePost SERVICE --")
		return nil, err
	}
	return post, nil
}

func (p *PostsCache) DeletePost(ctx context.Context, req *posts.DeletePostRequest) error {
	if err := p.Delete(ctx, req.PostId); err != nil {
		p.logger.Println("-- ERROR RETURNED FROM StorePost SERVICE --")
		return err
	}
	return nil
}

func (p *PostsCache) GetPostById(ctx context.Context, req *posts.GetPostByIdRequest) (*posts.Post, error) {
	byteData, err := p.redisDb.Get(ctx, req.PostId).Result()
	if err != nil {
		p.logger.Println("-- ERROR OCCURED WHILE GETTING DATA FROM REDIS IN GetPostById SERVICE", err.Error())
		return nil, err
	}
	var response posts.Post
	if err := json.Unmarshal([]byte(byteData), &response); err != nil {
		p.logger.Println("-- ERROR OCCURED WHILE UNMARSHALING DATA IN GetPostById SERVICE", err.Error())
		return nil, err
	}
	return &response, nil
}
