package main

import (
	"inventaris/app"
	"inventaris/controller"
	"inventaris/inventarispb/produkpb"
	"inventaris/repository"
	"inventaris/service"
	"log"
	"net"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

func main() {
	
	database := app.Db()
	validate := validator.New()
	produkRepo := repository.NewProdukRepositoryImpl(database)
	produkService := service.NewProdukServiceImpl(produkRepo, validate)
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