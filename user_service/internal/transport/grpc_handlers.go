package transport

import (
	"context"
	"log"

	"github.com/aAmer0neee/eCommerce/gen/user/v1"
	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userv1.UserServiceServer
	userService service.Service
}

func (h *UserHandler) RegisterUser(
	ctx context.Context,
	req *userv1.RegisterUserRequest,
) (*userv1.RegisterUserResponse, error) {

	log.Printf("input: %s %s", req.Email, req.Name)
	if err := validateRegisterInput(req); err != nil {
		return nil, err
	}

	h.userService.Register(&domain.RegisterInput{
		Email: req.Email,
		Name:  req.Name,
	})

	return &userv1.RegisterUserResponse{Id: "pidoras"}, nil
}

func validateRegisterInput(req *userv1.RegisterUserRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "Missing Email")
	}
	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "Missing Name")
	}

	return nil
}

func (h *UserHandler) UnregisterUser(
	context.Context,
	*userv1.UnregisterUserRequest,
) (*userv1.UnregisterUserResponse, error) {
	log.Print("unregister")
	return nil, nil
}

func (h *UserHandler) GetUser(
	context.Context,
	*userv1.GetUserRequest,
) (*userv1.GetUserResponse, error) {
	log.Print("Get")
	return nil, nil
}
