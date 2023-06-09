package main

import (
	context "context"
	"net"

	"github.com/NishanthSpShetty/pholo/proto"

	interceptors "github.com/NishanthSpShetty/grpc-interceptors"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPholoServer
}

// Echo implements proto.PholoServer
func (*server) Echo(ctx context.Context, req *proto.EchoMessage) (*proto.EchoMessage, error) {
	return &proto.EchoMessage{
		Message: req.GetMessage(),
	}, nil
}

// Healthz implements proto.PholoServer
func (*server) Healthz(context.Context, *proto.HealthZ) (*proto.HealthZ, error) {
	return &proto.HealthZ{}, nil
}

// Ping implements proto.PholoServer
func (*server) Ping(context.Context, *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Message: "pong",
	}, nil
}

func startGRPC(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "server.Start")
	}
	interceptor := interceptors.NewInterceptor("aos", log.Logger)
	opts := []grpc.ServerOption{interceptor.Get()}
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterPholoServer(grpcServer, &server{})

	log.Info().Str("address", addr).Msg("starting grpc server")
	return grpcServer.Serve(lis)
}

func main() {
	err := startGRPC("localhost:8090")
	if err != nil {
		log.Error().Err(err)
		return
	}
}
