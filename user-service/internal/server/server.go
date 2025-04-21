package server

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/JeroZp/gRPC-MOM/user-service/proto"
	"github.com/JeroZp/gRPC-MOM/user-service/internal/service"
)

// StartGRPCServer arranca el servidor gRPC en la dirección especificada.
func StartGRPCServer(address string) {
	// Se crea un Listener en el puerto especificado.
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error al abrir el puerto %s: %v", address, err)
	}

	// Se crea una nueva instancia del servidor gRPC.
	grpcServer := grpc.NewServer()

	// Se crea una instancia del servicio de usuario.
	userService := service.NewUserService()

	// Se registra el servicio de usuario en el servidor gRPC.
	pb.RegisterUserServiceServer(grpcServer, userService)

	// Se habilita la reflexión para el servidor gRPC.
	reflection.Register(grpcServer)

	log.Printf("Servidor gRPC de UserService escuchando en %s", address)

	// Se arranca el servidor para que escuche peticiones.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
