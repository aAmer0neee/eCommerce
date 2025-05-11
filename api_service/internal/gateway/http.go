package gateway

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type Http struct {
	router *gin.Engine
}

func newHttp() *Http {
	router := gin.Default()

	registerRoutes(router)

	return &Http{
		router: router,
	}
}

func (h *Http) Run(addr string) error {
	return h.router.Run(addr)
}

func (h *Http) Shutdown(){
	stop := make(chan os.Signal,1)

	signal.Notify(stop,os.Interrupt, syscall.SIGINT)

	<- stop

}

func registerRoutes(r *gin.Engine) {
	r.GET("/hello", func(ctx *gin.Context) {
		hello(ctx)
	})
}

// func registerMiddlewares()
