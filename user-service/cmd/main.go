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
	// 1) Conectamos con MOM (usando NewClient en v1.63+)
	conn, err := grpc.NewClient(
	  "localhost:50054",
	  grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
	  log.Fatalf("🤖 Error al conectar con MOM: %v", err)
	}
	defer conn.Close()
  
	momClient := mompb.NewMOMServiceClient(conn)
  
	// 2) Abrimos el stream de Subscribe
	stream, err := momClient.Subscribe(context.Background())
	if err != nil {
	  log.Fatalf("🤖 Error en Subscribe: %v", err)
	}
	// Indicamos qué cola queremos
	if err := stream.Send(&mompb.SubscribeRequest{Queue: "user_ops"}); err != nil {
	  log.Fatalf("🤖 Error enviando SubscribeRequest: %v", err)
	}
	log.Println("🤖 Worker suscrito a ‘user_ops’")
  
	// 3) Bucle de mensajes
	for {
	  msg, err := stream.Recv()
	  if err != nil {
		log.Printf("🤖 Error Recv: %v (reintentando en 1s)", err)
		time.Sleep(time.Second)
		continue
	  }
  
	  log.Printf("🤖 Mensaje recibido id=%s", msg.GetId())
  
	  // 4) Procesamos el payload
	  var p struct{ Name, Email string }
	  if err := json.Unmarshal(msg.GetData(), &p); err != nil {
		log.Printf("🤖 JSON inválido: %v", err)
		continue
	  }
  
	  // 5) Creamos el usuario
	  if _, err = usrSvc.CreateUser(context.Background(), &userpb.CreateUserRequest{
		User: &userpb.User{
		  Name:    p.Name,
		  Email:   p.Email,
		  Credits: 1000,
		},
	  }); err != nil {
		log.Printf("🤖 CreateUser error: %v", err)
		continue
	  }
	  log.Printf("✅ Usuario creado vía MOM: %s", p.Email)
  
	  // 6) ACK por RPC separado
	  if _, err := momClient.Ack(context.Background(), &mompb.AckRequest{
		Queue: "user_ops",
		Id:    msg.GetId(),
	  }); err != nil {
		log.Printf("❌ Error en ACK: %v", err)
	  } else {
		log.Printf("✅ ACK enviado para id=%s", msg.GetId())
	  }
	}
}  