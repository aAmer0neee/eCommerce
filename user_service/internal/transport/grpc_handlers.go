package transport

import (
	"context"
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

	if err := validateInput(req); err != nil {
		return nil, err
	}

	h.userService.Register(&domain.RegisterInput{
		Email: req.Email,
		Name:  req.Name,
	})
	return &userv1.RegisterUserResponse{Id: "pidoras"}, nil
}

func validateInput(req *userv1.RegisterUserRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "Missing Email")
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "Missing Password")
	}

	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "Missing Name")
	}

	return nil
}
