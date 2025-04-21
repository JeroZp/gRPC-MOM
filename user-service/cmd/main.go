package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/JeroZp/gRPC-MOM/user-service/proto"
	"github.com/JeroZp/gRPC-MOM/user-service/internal/service"
	mompb "github.com/JeroZp/gRPC-MOM/user-service/internal/mom-proto"
)

func main() {
	// 1) Arrancar servidor gRPC de UserService en :50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	usrSvc := service.NewUserService()
	userpb.RegisterUserServiceServer(grpcServer, usrSvc)

	go func() {
		log.Println("UserService gRPC escuchando en :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("serve gRPC: %v", err)
		}
	}()

	// 2) Arrancar worker que consume del MOM
	go startWorker(usrSvc)

	// 3) Mantener el proceso vivo
	select {}
}

func startWorker(usrSvc *service.UserService) {
	// 2.1) Crear canal al MOM usando NewClient
	conn, err := grpc.NewClient(
		"localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("no se pudo conectar a MOM: %v", err)
	}
	defer conn.Close()

	momClient := mompb.NewMOMServiceClient(conn)

	// 2.2) Abrir flujo de Subscribe
	stream, err := momClient.Subscribe(context.Background())
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}
	// enviamos el nombre de la cola que queremos consumir
	if err := stream.Send(&mompb.SubscribeRequest{Queue: "user_ops"}); err != nil {
		log.Fatalf("subscribe send: %v", err)
	}

	// 2.3) Bucle infinito de recepción → procesamiento → ACK
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("subscribe recv err: %v, reintentando en 1s", err)
			time.Sleep(time.Second)
			continue
		}

		// Deserializar payload JSON
		var p struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.Unmarshal(msg.Data, &p); err != nil {
			log.Printf("json unmarshal: %v", err)
			continue
		}

		// Llamada interna a CreateUser
		if _, err := usrSvc.CreateUser(context.Background(), &userpb.CreateUserRequest{
			User: &userpb.User{
				Name:    p.Name,
				Email:   p.Email,
				Credits: 100,
			},
		}); err != nil {
			log.Printf("CreateUser worker error: %v", err)
			continue
		}
		log.Printf("Usuario creado vía MOM: %s", p.Email)

		// Finalmente confirmamos el mensaje con RPC Ack
		if _, err := momClient.Ack(context.Background(), &mompb.AckRequest{
			Queue: "user_ops",
			Id:    msg.GetId(),
		}); err != nil {
			log.Printf("ack error: %v", err)
		}
	}
}