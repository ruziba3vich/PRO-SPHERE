package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/prosphere/internal/items/config"
	"github.com/ruziba3vich/prosphere/internal/items/http/handlers"
	"github.com/ruziba3vich/prosphere/internal/items/storage"
)

func Run(cfg *config.Config,
	logger *log.Logger) error {

	router := gin.Default()

	// storages

	currencyStorage := storage.NewCurrencyStorage(cfg, logger)

	// handlers

	currencyHandler := handlers.NewCurrencyHandler(currencyStorage, logger)

	// routers

	router.GET("/currency/:ccy", currencyHandler.GetCurrencyByCcy)
	router.GET("/currency", currencyHandler.GetAllCurrenciesHandler)

	// run
	if err := router.Run(cfg.Server.Port); err != nil {
		logger.Fatal(err)
	}
	return nil
}
