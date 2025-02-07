package integrations

import (
	"search_service/internal/items/config"

	"go.uber.org/zap"
)

type SearchIntegration struct {
	cfg    *config.Config
	logger *zap.Logger
}

func NewSearchInegreation(cfg *config.Config, logger *zap.Logger) *SearchIntegration {
	return &SearchIntegration{
		cfg:    cfg,
		logger: logger,
	}
}
