package gateway

type Gateway interface {
	Run(addr string) error
	Shutdown()
}

func New() Gateway {
	return newHttp()
}
