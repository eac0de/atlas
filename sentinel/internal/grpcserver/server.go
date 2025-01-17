package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/eac0de/atlas/sentinel/internal/services"
	pb "github.com/eac0de/atlas/sentinel/proto"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

type gprcAuthServer struct {
	pb.UnimplementedAuthServer

	Addr           string
	sessionService *services.SessionService
}

func NewAuthGRPCServer(addr string, sessionService *services.SessionService) *gprcAuthServer {
	return &gprcAuthServer{

		Addr:           addr,
		sessionService: sessionService,
	}
}

func (s *gprcAuthServer) AuthUser(ctx context.Context, req *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	claims, err := s.sessionService.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.AuthUserResponse{UserId: claims.UserID.String()}, nil
}

func (s *gprcAuthServer) Run() {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthServer(server, s)
	log.Printf("gRPC server is running at %s\n", s.Addr)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
