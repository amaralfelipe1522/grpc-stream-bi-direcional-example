package main

import (
	"log"
	"net"

	"github.com/amaralfelipe1522/grpc-stream-bi-direcional-example/pb"
	"github.com/amaralfelipe1522/grpc-stream-bi-direcional-example/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Registrando o serviço no servidor
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	// Modo Reflection para o client conseguir entender quais são os métodos existentes
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
}
