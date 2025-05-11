package transport

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	userv1 "github.com/aAmer0neee/eCommerce/gen/user/v1"
	"github.com/aAmer0neee/eCommerce/shared/logger"
	"github.com/aAmer0neee/eCommerce/user_service/internal/service"
	"google.golang.org/grpc"
)

type Transport struct {
	gRPCServer *grpc.Server
	log        logger.Logger
	port       string
}

func New(log logger.Logger, port string, userService service.Service) *Transport {
	gRPCServer := grpc.NewServer()
	register(gRPCServer, userService)
	return &Transport{
		gRPCServer: gRPCServer,
		log:        log,
		port:       port,
	}

}

func register(s *grpc.Server, userService service.Service) {
	userv1.RegisterUserServiceServer(s, &UserHandler{
		userService: userService,
	})
}

func (t *Transport) Run() error {
	listener, err := net.Listen("tcp", ":"+t.port)
	if err != nil {
		return err
	}

	t.log.Info("SERVER RUNNING", "port", t.port)

	if err := t.gRPCServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (t *Transport) Stop() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT)

	<-stop
	t.log.Info("SERVER STOPPING", "addr", t.port)
	t.gRPCServer.GracefulStop()
}


