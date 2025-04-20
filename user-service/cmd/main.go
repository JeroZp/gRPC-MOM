package main

import (
    "log"
    "github.com/JeroZp/gRPC-MOM/user-service/internal/server"
)

func main() {
    log.Println("Iniciando microservicio de usuarios...")
    server.StartGRPCServer(":50051")
}