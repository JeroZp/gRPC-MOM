package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/JeroZp/gRPC-MOM/notification-service/proto"
	"github.com/JeroZp/gRPC-MOM/notification-service/internal/service"
)

func main() {
	// Abrimos el socket TCP en el puerto 50053
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("no se pudo escuchar en el puerto 50053: %v", err)
	}

	// Creamos el servidor gRPC
	grpcServer := grpc.NewServer()

	pb.RegisterNotificationServiceServer(grpcServer, &service.Server{})

	log.Println("NotificationService escuchando en :50053")

	// Iniciamos el loop de servir peticiones
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error al servir gRPC: %v", err)
	}
}
