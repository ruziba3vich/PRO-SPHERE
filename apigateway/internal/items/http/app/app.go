/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-09-21 19:53:48
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2025-01-07 16:26:18
 * @FilePath: /sfere_backend/internal/items/http/app/app.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/auth"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/feeds"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/searching"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/config"
	_ "github.com/projects/pro-sphere-backend/apigateway/internal/items/http/app/docs"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/http/handlers"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/http/middleware"

	"go.uber.org/zap"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpService struct {
	searchingClient searching.SearchingServiceClient
	feedCatClient   feeds.CategoriesServiceClient
	feedsClient     feeds.FeedsServiceClient
	feedItemsClient feeds.FeedItemsServiceClient
	authClient      auth.AuthenticationClient
	userClient      auth.UserManagementClient
	logger          *zap.Logger
	cfg             *config.Config
}

func NewHttpService(searchingClient searching.SearchingServiceClient,
	feedCatClient feeds.CategoriesServiceClient,
	feedsClient feeds.FeedsServiceClient,
	feedItemsClient feeds.FeedItemsServiceClient,
	authClient auth.AuthenticationClient,
	userClient auth.UserManagementClient,
	logger *zap.Logger, cfg *config.Config) *HttpService {

	return &HttpService{
		feedCatClient:   feedCatClient,
		feedsClient:     feedsClient,
		feedItemsClient: feedItemsClient,
		searchingClient: searchingClient,
		authClient:      authClient,
		userClient:      userClient,

		cfg:    cfg,
		logger: logger,
	}
}

// @title           Pro-Sphere
// @version         1.0
// @description     This is pro-sphere browser's APIs
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @name Authorization
// @in header
func (h *HttpService) Run() error {

	router := gin.Default()

	router.Use(middleware.CORS())

	router.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))

	searchHandler := handlers.NewSearchHandler(h.searchingClient, h.logger)

	adminHandler := handlers.NewAdminHandler(h.feedsClient, h.feedItemsClient, h.feedCatClient, h.logger)

	authHandler := handlers.NewAuthHandler(h.cfg, h.logger, h.authClient)

	userHandler := handlers.NewUserHandler(h.cfg, h.logger, h.userClient)
	// currencyRouter := router.Group("/currency")

	// currencyRouter.GET("/:ccy", currencyHandler.GetCurrencyByCcy)
	// currencyRouter.GET("", currencyHandler.GetAllCurrenciesHandler)

	// postsRouter := router.Group("/posts")

	// postsRouter.POST("", postHandler.CreatePost)
	// postsRouter.POST("/add/view/:post_id", postHandler.AddPostView)
	// postsRouter.GET("/:post_id", postHandler.GetPostById)
	// postsRouter.GET("", postHandler.GetAllPosts)
	// postsRouter.GET("/publisher/:publisher_id", postHandler.GetPostByPublisherId)
	// postsRouter.PUT("/:post_id", postHandler.UpdatePost)
	// postsRouter.DELETE("/:post_id", postHandler.DeletePost)

	authRouter := router.Group("/v1/auth/oauth")
	{
		authRouter.GET("/start", authHandler.StartAndRedirectToProID)
		authRouter.GET("/callback", authHandler.HandleCallBack)
		authRouter.GET("/tokens", authHandler.GetTokensByCode)
	}

	authAdminRouter := router.Group("/v1/auth/oauth/admin")
	{
		authAdminRouter.GET("/callback", authHandler.HandleCallBackForAdmin)
	}

	authMiddleware := middleware.AuthMiddleware(h.cfg.ApiKey)

	userRouter := router.Group("v1/users", authMiddleware)
	{
		userRouter.POST("", userHandler.CreateUser)
		userRouter.GET("/me", userHandler.GetUserByID)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.PUT("/:id", userHandler.UpdateUserByID)
		userRouter.GET("", userHandler.GetAllUsers)
	}

	adminRouter := router.Group("/v1/admin")
	{
		feeds := adminRouter.Group("/feeds")
		{
			categories := feeds.Group("/categories")
			{
				categories.POST("/", adminHandler.CreateFeedCategory)
				categories.GET("/:id", adminHandler.GetFeedCategoryByID)
				categories.PUT("/:id", adminHandler.UpdateFeedCategory)
				categories.DELETE("/:id", adminHandler.DeleteFeedCategory)
				categories.GET("/all", adminHandler.GetAllFeedCategories)
				categories.GET("/icon", adminHandler.ServeFeedCategoryIcon)
			}
			feed := feeds.Group("/feed")
			{
				feed.POST("/", adminHandler.CreateFeed)
				feed.GET("/:id", adminHandler.GetFeedByID)
				feed.PUT("/:id", adminHandler.UpdateFeed)
				feed.DELETE("/:id", adminHandler.DeleteFeed)
				feed.GET("/all", adminHandler.GetAllFeeds)
				feed.POST("/refresh")
				// feed.GET("/fetch/:id", adminHandler.FetchFeedItems)
				content := feeds.Group("/content")
				{
					content.POST("", adminHandler.AddFeedContent)
					content.GET("", adminHandler.GetFeedContent)
					content.GET("/all", adminHandler.GetAllFeedContent)
					content.DELETE("/:id", adminHandler.DeleteFeedContent)
					content.PUT("/:id", adminHandler.UpdateFeedContent)
				}
			}
			items := feeds.Group("/items")
			{
				items.POST("/", adminHandler.CreateFeedItem)
				items.GET("/:id", adminHandler.GetFeedItemByID)
				items.PUT("/:id", adminHandler.UpdateFeedItem)
				items.DELETE("/:id", adminHandler.DeleteFeedItem)
				items.GET("/all/:feed_id", adminHandler.GetAllFeedItemsByFeed)
			}
		}
	}

	searchRouter := router.Group("/v1/search")
	{
		searchRouter.GET("", searchHandler.Search)
		searchRouter.GET("/youtube", searchHandler.SearchVideo)
		searchRouter.GET("/images", searchHandler.SearchImages)
	}

	// run
	if err := router.Run(h.cfg.Server.Port); err != nil {
		h.logger.Panic("Server failed to run", zap.Error(err))
	}
	return nil
}
