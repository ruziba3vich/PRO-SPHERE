package app

import (
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/prosphere/internal/items/config"
	"github.com/ruziba3vich/prosphere/internal/items/http/handlers"
	"github.com/ruziba3vich/prosphere/internal/items/service"
	"github.com/ruziba3vich/prosphere/internal/items/storage"
	currencystorage "github.com/ruziba3vich/prosphere/internal/items/storage/currency"
	poststorage "github.com/ruziba3vich/prosphere/internal/items/storage/posts"
	"github.com/ruziba3vich/prosphere/internal/items/storage/rediso"
	redisCl "github.com/ruziba3vich/prosphere/internal/pkg"
)

func Run(cfg *config.Config,
	logger *log.Logger) error {

	router := gin.Default()

	postgres, err := storage.ConnectDB(cfg)

	// --------
	if err != nil {
		logger.Fatalln(err)
	}

	redis, err := redisCl.NewRedisDB(cfg)
	if err != nil {
		logger.Fatalln(err)
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// handlers

	currencyHandler := handlers.NewCurrencyHandler(service.New(
		currencystorage.NewCurrencyStorage(cfg, logger),
	), logger)

	postHandler := handlers.NewPostHandler(
		service.NewPostService(
			poststorage.NewPostStorage(rediso.NewPostCache(redis, logger), postgres, sqrl, logger),
		),
		logger)

	// routers

	currencyRouter := router.Group("/currency")

	currencyRouter.GET("/:ccy", currencyHandler.GetCurrencyByCcy)
	currencyRouter.GET("", currencyHandler.GetAllCurrenciesHandler)

	postsRouter := router.Group("/posts")

	postsRouter.POST("", postHandler.CreatePost)
	postsRouter.POST("/add/view/:post_id", postHandler.AddPostView)
	postsRouter.GET("/:post_id", postHandler.GetPostById)
	postsRouter.GET("", postHandler.GetAllPosts)
	postsRouter.GET("/publisher/:publisher_id", postHandler.GetPostByPublisherId)
	postsRouter.PUT("/:post_id", postHandler.UpdatePost)
	postsRouter.DELETE("/:post_id", postHandler.DeletePost)

	// run
	if err := router.Run(cfg.Server.Port); err != nil {
		logger.Fatal(err)
	}
	return nil
}
