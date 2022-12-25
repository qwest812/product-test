package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"test-auth/models"
	"test-auth/service/grpc/product"
)

type Client interface {
	Auth(ctx context.Context, clientID, login, password string) (models.Auth, error)
	CheckAuth(ctx context.Context, clientID, authToken string) (models.AuthStatus, error)
	GetProducts(ctx context.Context) ([]*models.Product, error)
	//FutureStream(ctx context.Context, responses chan *models.FutureEventResponse, events chan *models.FutureEvent) error
	//SpotStream(ctx context.Context, responses chan *models.SpotOrderResponse, orders chan *models.Order) error
}
type client struct {
	mapLastPrice map[string]float64
	grpcClient   product.AuthClient
	log          *logrus.Logger
}

func NewClient(conn grpc.ClientConnInterface, log *logrus.Logger) *client {
	return &client{
		log:        log,
		grpcClient: product.NewAuthClient(conn),
	}
}
func (c *client) Auth(ctx context.Context, clientID, login, password string) (models.Auth, error) {
	resp, err := c.grpcClient.Auth(ctx, &product.LoginRequest{
		ClientID: clientID,
		Login:    login,
		Password: password,
	})
	if err != nil {
		c.log.Error("GRPC", "Auth: %v", err)
	}
	return models.FromAuthGRPCToAuth(resp), nil
}
func (c *client) CheckAuth(ctx context.Context, clientID, authToken string) (models.AuthStatus, error) {
	resp, err := c.grpcClient.CheckAuth(ctx, &product.AuthStruct{
		ClientID:  clientID,
		AuthToken: authToken,
	})
	if err != nil {
		c.log.Error("GRPC", "CheckAuth: %v", err)
	}
	return models.FromAuthStatusToStatus(resp), nil
}

func (c *client) GetProducts(ctx context.Context) ([]*models.Product, error) {
	resp, err := c.grpcClient.Products(ctx, &product.Empty{})
	if err != nil {
		c.log.Error("GRPC", "Products: %v", err)
	}
	return models.FromGRPCProductToProduct(resp), nil
}
