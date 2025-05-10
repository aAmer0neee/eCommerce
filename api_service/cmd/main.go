package main

import (
	"github.com/aAmer0neee/eCommerce/api_service/internal/domain"
	"github.com/aAmer0neee/eCommerce/api_service/internal/gateway"
	"github.com/aAmer0neee/eCommerce/shared/config"
	"github.com/aAmer0neee/eCommerce/shared/logger"
)

func main() {
	cfg := &domain.ApiCfg{}
	config.MustLoad(cfg)

	logger.New(cfg.Logger.Level)

	gateway.New().Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
