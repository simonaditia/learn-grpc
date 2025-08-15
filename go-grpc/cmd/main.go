package main

import (
	"log"
	"net"

	"go_grpc/cmd/config"
	"go_grpc/cmd/services"
	productPb "go_grpc/go_grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
