package service

import (
	"context"

	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
)

type (
	PostService struct {
		storage repo.PostRepository
	}
)

func NewPostService(storage repo.PostRepository) repo.PostRepository {
	return &PostService{
		storage: storage,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *posts.CreatePostRequest) (*posts.Post, error) {
	return s.storage.CreatePost(ctx, req)
}

func (s *PostService) UpdatePost(ctx context.Context, req *posts.UpdatePostRequest) (*posts.Post, error) {
	return s.storage.UpdatePost(ctx, req)
}

func (s *PostService) GetPostById(ctx context.Context, req *posts.GetPostByIdRequest) (*posts.Post, error) {
	return s.storage.GetPostById(ctx, req)
}

func (s *PostService) GetAllPosts(ctx context.Context, req *posts.GetAllPostsRequest) (*posts.GetPostsResponse, error) {
	return s.storage.GetAllPosts(ctx, req)
}

func (s *PostService) GetPostByPublisherId(ctx context.Context, req *posts.GetPostByPublisherIdRequest) (*posts.GetPostsResponse, error) {
	return s.storage.GetPostByPublisherId(ctx, req)
}

func (s *PostService) AddPostView(ctx context.Context, req *posts.AddPostView) (*posts.Post, error) {
	return s.storage.AddPostView(ctx, req)
}

func (s *PostService) DeletePost(ctx context.Context, req *posts.DeletePostRequest) (*posts.Post, error) {
	return s.storage.DeletePost(ctx, req)
}
