package transaction_service

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

var TxClient pb.TransactionServiceClient

func init() {
	// Creamos un cliente gRPC usando la dirección y puerto de tu TransactionService
	client, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar a TransactionService: %v", err)
	}

	// Inicializamos el stub generado por protoc
	TxClient = pb.NewTransactionServiceClient(client)
}
