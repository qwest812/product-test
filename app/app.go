package app

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"test-auth/config"
	client "test-auth/service/grpc"
)

type App struct {
	config *config.Config
	grpc   client.Client
	logger *logrus.Logger
}

func New(config *config.Config, grpcService grpc.ClientConnInterface, logger *logrus.Logger) *App {
	grpcClient := client.NewClient(grpcService, logger)
	return &App{
		config: config,
		grpc:   grpcClient,

		logger: logger,
	}
}

func (a *App) Process(ctx context.Context) error {
	http.HandleFunc("/products", a.Products)
	return http.ListenAndServe(a.config.HTTP.Address, nil)
}

func (a *App) Products(w http.ResponseWriter, req *http.Request) {
	res, err := a.grpc.GetProducts(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("Get products", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		a.logger.Errorf("Encode", err)
	}

}
