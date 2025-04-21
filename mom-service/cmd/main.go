package main

import (
  "log"
  "net"

  "github.com/JeroZp/gRPC-MOM/mom-service/internal/queue"
  "github.com/JeroZp/gRPC-MOM/mom-service/internal/service"
  pb "github.com/JeroZp/gRPC-MOM/mom-service/proto"
  "google.golang.org/grpc"
)

func main() {
  lis, err := net.Listen("tcp", ":50054")
  if err != nil {
    log.Fatalf("listen: %v", err)
  }
  store, err := queue.NewStore("mom.db")
  if err != nil {
    log.Fatalf("store: %v", err)
  }

  grpcServer := grpc.NewServer()
  pb.RegisterMOMServiceServer(grpcServer, service.NewMOMService(store))

  log.Println("MOMService escuchando en :50054")
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("serve: %v", err)
  }
}
