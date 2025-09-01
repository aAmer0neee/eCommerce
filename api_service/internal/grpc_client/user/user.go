package grpc_client

import (
	"context"
	"log"

	"time"

	"github.com/aAmer0neee/eCommerce/api_service/internal/domain"
	userv1 "github.com/aAmer0neee/eCommerce/gen/user/v1"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn    *grpc.ClientConn
	client  userv1.UserServiceClient
	timeout time.Duration
}

func NewUserClient(addr string, t time.Duration) (*UserClient, error) {
	//!!! insecure следут добавить tls
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := userv1.NewUserServiceClient(conn)

	return &UserClient{
		conn:    conn,
		client:  client,
		timeout: t,
	}, nil
}

func (uc *UserClient) Register(ctx context.Context, user *domain.User) {

	_, err := uc.client.RegisterUser(ctx, &userv1.RegisterUserRequest{
		Email: user.Email,
		Name:  user.Name,
	})

	log.Printf("connect to %s err : %s", uc.conn.GetState().String(), err.Error())
}

func (uc *UserClient) Unregister(ctx context.Context) {}
func (uc *UserClient) Get(ctx context.Context)        {}
