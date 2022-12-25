package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"test-auth/app"
	"test-auth/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	log := logrus.New()
	log.Debug("CONFIG HTTP", cfg.HTTP.String())
	log.Debug("CONFIG GRPC", cfg.GRPC.String())

	grpcConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.HOST, cfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("init", "failed open grpcConnection, err: %w", err)
	}

	a := app.New(cfg, grpcConnection, log)
	err = a.Process(ctx)
	if err != nil {
		log.Fatal("init", "failed application start, err: %s", err)
	}
}
