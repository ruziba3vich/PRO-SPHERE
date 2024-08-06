package storage

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/prosphere/internal/items/models/posts"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
	"github.com/ruziba3vich/prosphere/internal/items/storage/rediso"
)

type PostStorage struct {
	rediso   *rediso.PostsCache
	logger   *log.Logger
	postgres *sql.DB
	builder  sq.StatementBuilderType
}

func NewPostStorage(redis *rediso.PostsCache, postgres *sql.DB, builder sq.StatementBuilderType, logger *log.Logger) repo.PostRepository {
	return &PostStorage{
		rediso:   redis,
		postgres: postgres,
		logger:   logger,
		builder:  builder,
	}
}

func (p *PostStorage) CreatePost(ctx context.Context, req *posts.CreatePostRequest) (*posts.Post, error) {
	tx, err := p.postgres.BeginTx(ctx, nil)
	if err != nil {
		p.logger.Println("Error while starting a transaction")
		return nil, err
	}
	defer tx.Rollback()

	postId := uuid.New().String()
	query, args, err := p.builder.Insert("posts").
		Columns("post_id", "publisher_id", "title", "content").
		Values(postId, req.PublisherId, req.PostTitle, req.PostContent).
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}

	post := &posts.Post{
		PostId:      postId,
		PublisherId: req.PublisherId,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
	}
	result, err := p.rediso.StorePost(ctx, post)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		p.logger.Println("Error while committing transaction:", err.Error())
	}
	return result, nil
}

func (p *PostStorage) GetPostById(ctx context.Context, req *posts.GetPostByIdRequest) (*posts.Post, error) {
	redisPost, _ := p.rediso.GetPostById(ctx, req)
	if redisPost != nil {
		return redisPost, nil
	}

	query, args, err := p.builder.Select("post_id", "publisher_id", "title", "content").
		From("posts").
		Where(sq.Eq{"post_id": req.PostId}).
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	row := p.postgres.QueryRowContext(ctx, query, args...)
	var post posts.Post
	if err := row.Scan(&post.PostId, &post.PublisherId, &post.PostTitle, &post.PostContent); err != nil {
		p.logger.Println(err)
		return nil, err
	}
	return &post, nil
}

func (p *PostStorage) GetPostByPublisherId(ctx context.Context, req *posts.GetPostByPublisherIdRequest) (*posts.GetPostsResponse, error) {
	query, args, err := p.builder.Select("post_id", "publisher_id", "title", "content").
		From("posts").
		Where(sq.Eq{"publisher_id": req.PublisherId}).
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	rows, err := p.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var postsList []*posts.Post
	for rows.Next() {
		var post posts.Post
		if err := rows.Scan(&post.PostId, &post.PublisherId, &post.PostTitle, &post.PostContent); err != nil {
			p.logger.Println(err)
			return nil, err
		}
		postsList = append(postsList, &post)
	}
	if err := rows.Err(); err != nil {
		p.logger.Println(err)
		return nil, err
	}
	return &posts.GetPostsResponse{Posts: postsList}, nil
}

func (p *PostStorage) GetAllPosts(ctx context.Context, req *posts.GetAllPostsRequest) (*posts.GetPostsResponse, error) {
	query, args, err := p.builder.Select("post_id", "publisher_id", "title", "content").
		From("posts").
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	rows, err := p.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var postsList []*posts.Post
	for rows.Next() {
		var post posts.Post
		if err := rows.Scan(&post.PostId, &post.PublisherId, &post.PostTitle, &post.PostContent); err != nil {
			p.logger.Println(err)
			return nil, err
		}
		postsList = append(postsList, &post)
	}
	if err := rows.Err(); err != nil {
		p.logger.Println(err)
		return nil, err
	}
	return &posts.GetPostsResponse{Posts: postsList}, nil
}

func (p *PostStorage) UpdatePost(ctx context.Context, req *posts.UpdatePostRequest) (*posts.Post, error) {
	tx, err := p.postgres.BeginTx(ctx, nil)
	if err != nil {
		p.logger.Println("Error while starting a transaction")
		return nil, err
	}
	defer tx.Rollback()

	queryBuilder := p.builder.Update("posts")

	if len(req.PostTitle) > 0 {
		queryBuilder = queryBuilder.Set("title", req.PostTitle)
	}
	if len(req.PostContent) > 0 {
		queryBuilder = queryBuilder.Set("content", req.PostContent)
	}

	queryBuilder = queryBuilder.Where(sq.Eq{"post_id": req.PostId})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}

	result, err := p.postgres.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	updatedPost, err := p.GetPostById(ctx, &posts.GetPostByIdRequest{PostId: req.PostId})
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}

	redisPost, err := p.rediso.StorePost(ctx, updatedPost)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		p.logger.Println("Error while committing transaction:", err.Error())
	}
	return redisPost, nil
}

func (p *PostStorage) DeletePost(ctx context.Context, req *posts.DeletePostRequest) (*posts.Post, error) {
	tx, err := p.postgres.BeginTx(ctx, nil)
	if err != nil {
		p.logger.Println("Error starting transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := p.builder.Delete("posts").
		Where(sq.Eq{"post_id": req.PostId}).
		ToSql()
	if err != nil {
		p.logger.Println("Error building SQL query:", err)
		return nil, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Println("Error executing SQL query:", err)
		return nil, err
	}

	ra, err := res.RowsAffected()
	if ra == 0 || err != nil {
		p.logger.Println("No rows affected:", err)
		return nil, err
	}

	if err := p.rediso.DeletePost(ctx, req); err != nil {
		p.logger.Println("Error deleting post from Redis:", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		p.logger.Println("Error committing transaction:", err)
		return nil, err
	}

	return nil, nil
}
