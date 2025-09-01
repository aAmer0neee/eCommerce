package gateway

import (
	"os"
	"os/signal"
	"syscall"

	service "github.com/aAmer0neee/eCommerce/api_service/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Http struct {
	router *gin.Engine
}

func newHttp(us service.UserService) *Http {
	router := gin.Default()

	registerRoutes(router, us)

	return &Http{
		router: router,
	}
}

func (h *Http) Run(addr string) error {
	return h.router.Run(addr)
}

func (h *Http) Shutdown() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT)

	<-stop

}

func registerRoutes(r *gin.Engine, userService service.UserService) {
	{
		handler := newUserHandler(userService)
		user := r.Group("/user")

		user.GET("get", handler.getUserHandler)
		user.POST("register", handler.registerUserHandler)
		user.POST("unregister", handler.unregisterUserHandler)
	}

	r.GET("/hello", func(ctx *gin.Context) {
		hello(ctx)
	})
}

// func registerMiddlewares()
