package gateway

import (
	"net/http"

	"github.com/aAmer0neee/eCommerce/api_service/internal/domain"
	service "github.com/aAmer0neee/eCommerce/api_service/internal/service/user"
	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello",
		"code":    http.StatusOK,
	})
}

type UserHandler struct {
	service service.UserService
}

func newUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}
func (uh *UserHandler) registerUserHandler(ctx *gin.Context) {
	request := &RegisterUserRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, RegisterUserResponse{
			Message: ResponseMessage{
				Code:    http.StatusBadRequest,
				Message: "incorrect input",
			},
		})
		return
	}

	uh.service.Register(ctx.Request.Context(), &domain.User{
		Email: request.Email,
		Name:  request.Name,
	})
}

func (uh *UserHandler) getUserHandler(ctx *gin.Context) {
}

func (uh *UserHandler) unregisterUserHandler(ctx *gin.Context) {}
