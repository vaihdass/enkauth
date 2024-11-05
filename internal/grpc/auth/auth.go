package auth

import (
	"github.com/vaihdass/enkproto/gen/go/sso/v1"
	"google.golang.org/grpc"
)

func Register(gRPC *grpc.Server, server sso.AuthServiceServer) {
	sso.RegisterAuthServiceServer(gRPC, server)
}

type Handler struct {
	sso.UnimplementedAuthServiceServer
}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}
