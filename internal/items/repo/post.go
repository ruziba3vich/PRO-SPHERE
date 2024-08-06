package repo

import (
	"context"

	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
)

type (
	PostRepository interface {
		CreatePost(context.Context, *posts.CreatePostRequest) (*posts.Post, error)
		GetPostById(context.Context, *posts.GetPostByIdRequest) (*posts.Post, error)
		GetPostByPublisherId(context.Context, *posts.GetPostByPublisherIdRequest) (*posts.GetPostsResponse, error)
		GetAllPosts(context.Context, *posts.GetAllPostsRequest) (*posts.GetPostsResponse, error)
		UpdatePost(context.Context, *posts.UpdatePostRequest) (*posts.Post, error)
		DeletePost(context.Context, *posts.DeletePostRequest) (*posts.Post, error)
	}
)
