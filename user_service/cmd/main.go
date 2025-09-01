package main

import (
	"github.com/aAmer0neee/eCommerce/shared/config_loader"
	"github.com/aAmer0neee/eCommerce/shared/logger"
	"github.com/aAmer0neee/eCommerce/user_service/internal/config"
	"github.com/aAmer0neee/eCommerce/user_service/internal/service"
	"github.com/aAmer0neee/eCommerce/user_service/internal/storage"
	"github.com/aAmer0neee/eCommerce/user_service/internal/transport"
)

func main() {
	cfg := &config.Cfg{}
	config_loader.MustLoad(cfg)
	logger := logger.New(cfg.Logger.Level)

	storage, _ := storage.New(cfg)
	service := service.New(storage, logger)

	server := transport.New(logger, cfg.GRPC.Port, service)
	go func() {
		if err := server.Run(); err != nil {
			logger.Warn(err.Error())
		}
	}()

	server.Stop()
}
