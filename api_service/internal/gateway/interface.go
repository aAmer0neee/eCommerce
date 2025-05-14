package gateway

import service "github.com/aAmer0neee/eCommerce/api_service/internal/service/user"

type Gateway interface {
	Run(addr string) error
	Shutdown()
}

func New(us service.UserService) Gateway {
	return newHttp(us)
}
