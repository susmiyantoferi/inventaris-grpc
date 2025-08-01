package main

import (
	"inventaris/app"
	"inventaris/controller"
	"inventaris/inventarispb/produkpb"
	"inventaris/repository"
	"inventaris/service"
	"log"
	"net"
	logging "inventaris/logging"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

func main() {
	
	database := app.DbGrpc()
	validate := validator.New()
	logging := logging.ConsoleLogging{}
	produkRepo := repository.NewProdukRepositoryImpl(database)
	produkService := service.NewProdukServiceImpl(produkRepo, validate, &logging)
	produkGrpcController := controller.NewProdukGRPCServer(produkService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	produkpb.RegisterProdukServiceServer(grpcServer, produkGrpcController)

	log.Println("grpc listrn to server :50051")
	if err = grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}