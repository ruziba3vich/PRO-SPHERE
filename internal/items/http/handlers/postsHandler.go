package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/prosphere/internal/items/models/errors"
	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
)

type (
	PostHandler struct {
		storage repo.PostRepository
		logger  *log.Logger
	}
)

func NewPostHandler(storage repo.PostRepository, logger *log.Logger) *PostHandler {
	return &PostHandler{
		storage: storage,
		logger:  logger,
	}
}

func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var req posts.CreatePostRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Println("Error binding request:", err)
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Invalid request body",
			Err:     err.Error(),
		})
		return
	}

	post, err := h.storage.CreatePost(ctx, &req)
	if err != nil {
		h.logger.Println("Error creating post:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to create post",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, post)
}

func (h *PostHandler) UpdatePost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	if postId == "" {
		h.logger.Println("Post ID is required")
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Post ID is required",
		})
		return
	}

	var req posts.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Println("Error binding request:", err)
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Invalid request body",
			Err:     err.Error(),
		})
		return
	}
	req.PostId = postId
	post, err := h.storage.UpdatePost(ctx, &req)
	if err != nil {
		h.logger.Println("Error updating post:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to update post",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetPostById(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	if postId == "" {
		h.logger.Println("Post ID is required")
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Post ID is required",
		})
		return
	}

	post, err := h.storage.GetPostById(ctx, &posts.GetPostByIdRequest{PostId: postId})
	if err != nil {
		h.logger.Println("Error retrieving post:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to retrieve post",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetAllPosts(ctx *gin.Context) {
	page, limit := ctx.Query("page"), ctx.Query("limit")
	response, err := h.storage.GetAllPosts(ctx, &posts.GetAllPostsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		h.logger.Println("Error retrieving posts:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to retrieve posts",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *PostHandler) GetPostByPublisherId(ctx *gin.Context) {
	publisherId := ctx.Param("publisher_id")
	if publisherId == "" {
		h.logger.Println("Publisher ID is required")
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Publisher ID is required",
		})
		return
	}

	response, err := h.storage.GetPostByPublisherId(ctx, &posts.GetPostByPublisherIdRequest{PublisherId: publisherId})
	if err != nil {
		h.logger.Println("Error retrieving posts by publisher:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to retrieve posts by publisher",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *PostHandler) AddPostView(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	if postId == "" {
		h.logger.Println("Post ID is required")
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Post ID is required",
		})
		return
	}

	req := &posts.AddPostView{PostId: postId}
	post, err := h.storage.AddPostView(ctx, req)
	if err != nil {
		h.logger.Println("Error adding post view:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to add post view",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	if postId == "" {
		h.logger.Println("Post ID is required")
		ctx.JSON(http.StatusBadRequest, errors.ProError{
			Message: "Post ID is required",
		})
		return
	}

	post, err := h.storage.DeletePost(ctx, &posts.DeletePostRequest{PostId: postId})
	if err != nil {
		h.logger.Println("Error deleting post:", err)
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "Failed to delete post",
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}
