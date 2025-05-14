package transport

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/aAmer0neee/eCommerce/auth_service/internal/service"
	authv1 "github.com/aAmer0neee/eCommerce/gen/auth/v1"
	"github.com/aAmer0neee/eCommerce/shared/logger"
	"google.golang.org/grpc"
)

type Transport struct {
	gRPCServer *grpc.Server
	log        logger.Logger
	port       string
}

func New(log logger.Logger, port string, authService service.Service) *Transport {
	gRPCServer := grpc.NewServer()
	register(gRPCServer, authService)
	return &Transport{
		gRPCServer: gRPCServer,
		log:        log,
		port:       port,
	}
}

func register(s *grpc.Server, authService service.Service) {
	authv1.RegisterAuthServiceServer(s, &AuthHandler{
		authService: authService,
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
