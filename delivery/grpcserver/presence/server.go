package presence

import (
	"context"
	"gameapp/contract/goproto/presence"
	"gameapp/pkg"
	presenceservice "gameapp/service/presence"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	presence.UnimplementedPresenceServer
	presenceservice.Service
}

func New(service presenceservice.Service) Server {

	return Server{
		UnimplementedPresenceServer: presence.UnimplementedPresenceServer{},
		Service:                     service,
	}
}
func (s Server) GetPresence(ctx context.Context, request *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {

	req := pkg.MapToGetPresenceRequest(*request)

	res, err := s.Service.GetPresence(context.Background(), *req)

	//TODO handle err
	if err != nil {
		return &presence.GetPresenceResponse{}, err
	}

	return pkg.MapToProGetPresenceResponse(res), nil
}

func (s Server) Start() {

	address := ":8086"

	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	presence.RegisterPresenceServer(grpcServer, s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("couldn't server presence grpc server")
	}

}
