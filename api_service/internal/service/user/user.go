package service

import (
	"context"

	"github.com/aAmer0neee/eCommerce/api_service/internal/domain"
	grpc_client "github.com/aAmer0neee/eCommerce/api_service/internal/grpc_client/user"
	"github.com/aAmer0neee/eCommerce/shared/logger"
)

type UserService interface {
	Register(ctx context.Context, user *domain.User)
	Unregister(ctx context.Context)
	Get(ctx context.Context, user *domain.User)
}

type userService struct {
	client *grpc_client.UserClient
	log    logger.Logger
}

func NewUserService(client *grpc_client.UserClient, log logger.Logger) UserService {
	return &userService{
		client: client,
		log:    log,
	}
}

func (us *userService) Register(ctx context.Context, user *domain.User) {
	us.client.Register(ctx, user)
}
func (us *userService) Unregister(ctx context.Context)             {}
func (us *userService) Get(ctx context.Context, user *domain.User) {}
