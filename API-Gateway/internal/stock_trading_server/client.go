package stock_trading_server

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/stock_trading_server/proto"
)

var StockClient proto.StockTradingServiceClient

// Esta función se llama desde main.go para inicializar el cliente gRPC
func init() {
	// Establecer la conexión con el microservicio de stock trading
	client, err := grpc.NewClient(
		"localhost:9090",             // Dirección del microservicio
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar al microservicio gRPC: %v", err)
	}

	// Crear el cliente gRPC para StockTradingService
	StockClient = proto.NewStockTradingServiceClient(client)
}
