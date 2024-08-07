package poststorage

import (
	"context"
	"database/sql"
	"log"
	"strconv"

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
		p.logger.Println("Error while starting a transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	postId := uuid.New().String()
	query, args, err := p.builder.Insert("posts").
		Columns(
			"post_id",
			"publisher_id",
			"post_title",
			"post_category",
			"post_short_content",
			"post_content",
			"post_source",
			"imported_data",
			"views").
		Values(
			postId,
			req.PublisherId,
			req.PostTitle,
			req.PostCategory,
			req.PostShortContent,
			req.PostContent,
			req.PostSource,
			req.ImportedData,
			0).
		ToSql()
	if err != nil {
		p.logger.Println("Error building SQL query:", err)
		return nil, err
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Println("Error executing SQL query:", err)
		return nil, err
	}

	post := posts.Post{
		PostId:            postId,
		PublisherId:       req.PublisherId,
		PostTitle:         req.PostTitle,
		PostCategory:      req.PostCategory,
		PostShortContent:  req.PostShortContent,
		PostContent:       req.PostContent,
		PostSource:        req.PostSource,
		ImportedData:      req.ImportedData,
		PostViews:         0,
		PostFeaturedImage: "demo_image" + postId,
		Deleted:           false,
	}

	if p.rediso != nil {
		if err := p.rediso.StorePost(ctx, &post); err != nil {
			p.logger.Println("Error storing post in cache:", err)
			return nil, err
		}
	} else {
		p.logger.Println("Redis cache is not initialized")
	}

	if err := tx.Commit(); err != nil {
		p.logger.Println("Error while committing transaction:", err)
		return nil, err
	}
	return &post, nil
}

func (p *PostStorage) GetPostById(ctx context.Context, req *posts.GetPostByIdRequest) (*posts.Post, error) {
	redisPost, _ := p.rediso.GetPostById(ctx, req)
	if redisPost != nil {
		p.logger.Println("-- GOT FROM REDIS --")
		return redisPost, nil
	}

	query, args, err := p.builder.Select(
		"post_id",
		"publisher_id",
		"post_title",
		"post_category",
		"post_short_content",
		"post_content",
		"post_source",
		"imported_data",
		"views",
		"deleted").
		From("posts").
		Where(
			sq.Eq{"post_id": req.PostId, "deleted": false},
		).
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	row := p.postgres.QueryRowContext(ctx, query, args...)
	var post posts.Post
	if err := row.Scan(
		&post.PostId,
		&post.PublisherId,
		&post.PostTitle,
		&post.PostCategory,
		&post.PostShortContent,
		&post.PostContent,
		&post.PostSource,
		&post.ImportedData,
		&post.PostViews,
		&post.Deleted,
	); err != nil {
		p.logger.Println(err)
		return nil, err
	}
	if err := p.rediso.StorePost(ctx, &post); post.Deleted || err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostStorage) GetPostByPublisherId(ctx context.Context, req *posts.GetPostByPublisherIdRequest) (*posts.GetPostsResponse, error) {
	query, args, err := p.builder.Select(
		"post_id",
		"publisher_id",
		"post_title",
		"post_category",
		"post_short_content",
		"post_content",
		"post_source",
		"imported_data",
		"views").
		From("posts").
		Where(
			sq.Eq{"publisher_id": req.PublisherId, "deleted": false}).
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
		if err := rows.Scan(
			&post.PostId,
			&post.PublisherId,
			&post.PostTitle,
			&post.PostCategory,
			&post.PostShortContent,
			&post.PostContent,
			&post.PostSource,
			&post.ImportedData,
			&post.PostViews,
		); err != nil {
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
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		p.logger.Println("-- ERROR WHILE CONVERTING THE VALUE --", err.Error())
		return nil, err
	}
	page, err := strconv.Atoi(req.Page)
	if err != nil {
		p.logger.Println("-- ERROR WHILE CONVERTING THE VALUE --", err.Error())
		return nil, err
	}

	query, args, err := p.builder.Select(
		"post_id",
		"publisher_id",
		"post_title",
		"post_category",
		"post_short_content",
		"post_content",
		"post_source",
		"imported_data",
		"views",
	).
		From("posts").
		Where(sq.Eq{"deleted": false}).
		Limit(uint64(limit)).
		Offset(uint64(page)).
		ToSql()
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	// pp.Println(query)
	rows, err := p.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var postsList []*posts.Post
	for rows.Next() {
		var post posts.Post
		if err := rows.Scan(
			&post.PostId,
			&post.PublisherId,
			&post.PostTitle,
			&post.PostCategory,
			&post.PostShortContent,
			&post.PostContent,
			&post.PostSource,
			&post.ImportedData,
			&post.PostViews); err != nil {
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

	post, err := p.GetPostById(ctx, &posts.GetPostByIdRequest{PostId: req.PostId})
	if err != nil {
		p.logger.Println(err)
		return nil, err
	}

	queryBuilder := p.builder.Update("posts")

	if len(req.PostTitle) > 0 {
		queryBuilder = queryBuilder.Set("post_title", req.PostTitle)
		post.PostTitle = req.PostTitle
	}
	if len(req.PostContent) > 0 { //
		queryBuilder = queryBuilder.Set("post_content", req.PostContent)
		post.PostContent = req.PostContent
	}
	if len(req.PostCategory) > 0 {
		queryBuilder = queryBuilder.Set("post_category", req.PostCategory)
		post.PostCategory = req.PostCategory
	}
	if len(req.PostShortContent) > 0 {
		queryBuilder = queryBuilder.Set("post_short_content", req.PostShortContent)
		post.PostShortContent = req.PostShortContent
	}
	if len(req.ImportedData) > 0 {
		queryBuilder = queryBuilder.Set("imported_data", req.ImportedData)
		post.ImportedData = req.ImportedData
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

	if err := p.rediso.StorePost(ctx, post); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		p.logger.Println("Error while committing transaction:", err.Error())
	}
	return post, nil
}

func (p *PostStorage) DeletePost(ctx context.Context, req *posts.DeletePostRequest) (*posts.Post, error) {
	tx, err := p.postgres.BeginTx(ctx, nil)
	post, err := p.GetPostById(ctx, &posts.GetPostByIdRequest{PostId: req.PostId})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if err != nil {
		p.logger.Println("Error starting transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := p.builder.Update("posts").
		Set("deleted", true).
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

	return post, nil
}

func (p *PostStorage) AddPostView(ctx context.Context, req *posts.AddPostView) (*posts.Post, error) {
	tx, err := p.postgres.BeginTx(ctx, nil)
	if err != nil {
		p.logger.Println("Error starting transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	post, err := p.GetPostById(ctx, &posts.GetPostByIdRequest{PostId: req.PostId})
	if err != nil {
		p.logger.Println("Error fetching post by ID:", err)
		return nil, err
	}
	post.PostViews++

	query, args, err := p.builder.Update("posts").
		Set("views", post.PostViews).
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
		p.logger.Println("-- NO ROWS AFFECTED --")
		return nil, err
	}

	if err := p.rediso.StorePost(ctx, post); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		p.logger.Println("Error committing transaction:", err)
		return nil, err
	}

	return post, nil
}
