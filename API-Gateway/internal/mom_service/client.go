package mom_service

import (
  "log"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

var MOMClient pb.MOMServiceClient

func init() {
	// Establecer la conexión con el microservicio de usuarios
	  client, err := grpc.NewClient(
		  "localhost:50054",             // Dirección del mom
		  grpc.WithTransportCredentials(insecure.NewCredentials()),
	  )
	  if err != nil {
		  log.Fatalf("No se pudo conectar al MOMService: %v", err)
	  }
  
	  // Crear el cliente gRPC
	  MOMClient = pb.NewMOMServiceClient(client)
}