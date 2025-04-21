package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    pb "github.com/JeroZp/gRPC-MOM/user-service/proto"
    "github.com/JeroZp/gRPC-MOM/user-service/internal/service"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
      log.Fatalf("failed to listen: %v", err)
    }
  
    grpcServer := grpc.NewServer()
    // Aquí registras tu UserService concreto:
    pb.RegisterUserServiceServer(grpcServer, service.NewUserService())
  
    log.Println("UserService escuchando en :50051")
    if err := grpcServer.Serve(lis); err != nil {
      log.Fatalf("failed to serve: %v", err)
    }
  }