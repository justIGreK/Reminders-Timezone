package main

import (
	"context"
	"log"
	"net"

	"github.com/justIGreK/Reminders-Timezone/cmd/handler"
	"github.com/justIGreK/Reminders-Timezone/internal/repository"
	"github.com/justIGreK/Reminders-Timezone/internal/service"
	"github.com/justIGreK/Reminders-Timezone/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	td := client.NewTimeDiff()
	db := repository.CreateMongoClient(ctx)
	txRepo := repository.NewTimezoneRepository(db)
	txSRV := service.NewTimezoneService(txRepo, td)
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	handler := handler.NewHandler(grpcServer, txSRV)
	handler.RegisterServices()
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
