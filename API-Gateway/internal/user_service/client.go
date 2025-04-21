package user_service

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // Para credenciales inseguras
	pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto" // Paquete generado por protoc
)

var UserClient pb.UserServiceClient

func init() {
	// Establecer la conexión con el microservicio de usuarios
	client, err := grpc.NewClient(
		"localhost:50051",             // Dirección del microservicio
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar al microservicio gRPC: %v", err)
	}

	// Crear el cliente gRPC
	UserClient = pb.NewUserServiceClient(client)
}