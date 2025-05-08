package gateway

import "github.com/gin-gonic/gin"

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

func (h *Http) Shutdown() error {
	panic("not implemented")
}

func registerRoutes(r *gin.Engine) {
	r.GET("/hello", func(ctx *gin.Context) {
		hello(ctx)
	})
}

// func registerMiddlewares()
