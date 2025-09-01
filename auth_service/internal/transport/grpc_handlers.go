package transport

import (
	"context"

	"github.com/aAmer0neee/eCommerce/auth_service/internal/service"
	authv1 "github.com/aAmer0neee/eCommerce/gen/auth/v1"
)

type AuthHandler struct {
	authv1.AuthServiceServer

	authService service.Service
}

func (h *AuthHandler) Authenticate(
	context.Context,
	*authv1.AuthenticateRequest,
) (*authv1.AuthenticateResponse, error) {
	return nil, nil
}

func (h *AuthHandler) Refresh(
	context.Context,
	*authv1.RefreshRequest,
) (*authv1.RefreshResponse, error) {
	return nil, nil
}
