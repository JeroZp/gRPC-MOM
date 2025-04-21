package notification_service

import (
  "log"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

var NotifClient pb.NotificationServiceClient

func init() {
  // Establecer la conexión con el microservicio de usuarios
	client, err := grpc.NewClient(
		"localhost:50053",             // Dirección del microservicio
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar al microservicio gRPC: %v", err)
	}

	// Crear el cliente gRPC
	NotifClient = pb.NewNotificationServiceClient(client)
}