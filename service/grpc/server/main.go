package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"test-auth/service/grpc/product"
	"time"
)

func main() {

	lis, err := net.Listen("tcp", "127.0.0.1:5443")
	if err != nil {
		fmt.Printf("%v", err)
	}

	grpcServer := grpc.NewServer()
	product.RegisterAuthServer(grpcServer, NewServer())

	for {
		err = grpcServer.Serve(lis)
		if err != nil {
			fmt.Printf("Error GRPC Serve: %v", err)
		}
		time.Sleep(time.Second * 5)
	}

}

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	MapAuth map[string]string
	mu      sync.Mutex
	product.UnimplementedAuthServer
}

func (s *Server) Auth(ctx context.Context, loginRequest *product.LoginRequest) (*product.AuthStruct, error) {
	//todo auth service
	token := jwt.New(jwt.SigningMethodEdDSA)
	s.mu.Lock()
	s.MapAuth[loginRequest.ClientID] = token.Raw
	s.mu.Unlock()
	return &product.AuthStruct{
		ClientID:  loginRequest.ClientID,
		AuthToken: token.Raw,
	}, nil
}

func (s *Server) Products(ctx context.Context, req *product.Empty) (*product.ProductsResponse, error) {
	return &product.ProductsResponse{Products: []*product.Product{
		&product.Product{
			Id:    "1",
			Price: 100,
			Qty:   1000,
		},
		&product.Product{
			Id:    "2",
			Price: 200,
			Qty:   2000,
		},
	}}, nil
}

func (s *Server) CheckAuth(ctx context.Context, req *product.AuthStruct) (*product.CheckAuthResponse, error) {
	var val string
	var ok bool
	s.mu.Lock()
	val, ok = s.MapAuth[req.ClientID]
	s.mu.Unlock()
	if !ok {
		return &product.CheckAuthResponse{Status: 100}, fmt.Errorf("Inccorect token: %s for client: %s", req.AuthToken, req.ClientID)
	}
	if val != req.AuthToken {
		if !ok {
			return &product.CheckAuthResponse{Status: 102}, fmt.Errorf("Wrong token: %s for client: %s", req.AuthToken, req.ClientID)
		}
	}
	return &product.CheckAuthResponse{Status: 200}, nil
}
