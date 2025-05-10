package gateway

type Gateway interface {
	Run(addr string) error
	Shutdown() error
}

func New() Gateway {
	return newHttp()
}
