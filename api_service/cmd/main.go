package main

import (
	"github.com/aAmer0neee/eCommerce/api_service/internal/domain"
	"github.com/aAmer0neee/eCommerce/api_service/internal/gateway"
	grpc_client "github.com/aAmer0neee/eCommerce/api_service/internal/grpc_client/user"
	service "github.com/aAmer0neee/eCommerce/api_service/internal/service/user"
	"github.com/aAmer0neee/eCommerce/shared/config_loader"
	"github.com/aAmer0neee/eCommerce/shared/logger"
)

func main() {
	cfg := &domain.Cfg{}
	config_loader.MustLoad(cfg)

	logger := logger.New(cfg.Logger.Level)

	client, err := grpc_client.NewUserClient("localhost:"+cfg.Services.User, cfg.Services.Timeout)

	service := service.NewUserService(client, logger)
	if err != nil {
		logger.Warn("failed connect o user-service", "err", err.Error())
	}

	gateway.New(service).Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
