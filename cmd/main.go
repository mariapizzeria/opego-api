package main

import (
	"log"
	"net"
	"net/http"

	"github.com/mariapizzeria/opego-api/internal/order"
	"github.com/mariapizzeria/opego-api/pkg/configs"
	"github.com/mariapizzeria/opego-api/pkg/db"
	pb "github.com/mariapizzeria/opego-api/services/streaming/pb/proto"
	mygrpc "github.com/mariapizzeria/opego-api/services/streaming/server"
	"google.golang.org/grpc"
)

func main() {
	config := configs.LoadConfig()
	newDB := db.NewDb(config)
	router := http.NewServeMux()
	//gRPC server
	grpcServer := grpc.NewServer()
	streamServer := mygrpc.NewGRPCServer(newDB)
	pb.RegisterStreamServer(grpcServer, streamServer)
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen for gRPC: %v", err)
		}
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
	//repositories
	orderRepository := order.NewRepository(newDB)
	//handlers
	order.NewHandler(router, order.HandlerDeps{
		Repository: orderRepository,
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server is listening")
	server.ListenAndServe()
}
