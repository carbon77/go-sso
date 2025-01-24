package auth

import (
	ssov1 "github.com/carbon77/protos/gen/go/sso"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServiceServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("not implemented")
}
