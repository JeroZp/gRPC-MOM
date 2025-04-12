package client

import (
	"log"
	"google.golang.org/grpc"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service/proto"
)

var UserClient proto.UserServiceClient

// Inicializa el cliente gRPC para el microservicio de usuarios
func init() {
	// Conexión con el microservicio gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el microservicio gRPC: %v", err)
	}
	UserClient = proto.NewUserServiceClient(conn)
}
