package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	// Asegúrate de que este path coincida con el `go_package` o la ubicación de tu archivo `.pb.go`
	invoicer "microservice/Invoicer"
	"database/sql"
	_ "github.com/lib/pq"
)



type myInvoiceServer struct{
	invoicer.UnimplementedInvoicerServer
	db *sql.DB
}

func (s myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	_, err := s.db.Exec(`
		INSERT INTO transactions (from_user, to_user, amount, currency)
		VALUES ($1, $2, $3, $4)
	`, req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	if err != nil {
		log.Println("❌ Error al insertar en la base de datos:", err)
		return &invoicer.CreateResponse{
			Success: false,
			Message: "Error al guardar la transacción",
		}, nil
	}

	return &invoicer.CreateResponse{
		Success: true,
		Message: "Transacción guardada y creada correctamente",
		Transaction: &invoicer.Transaction{
			From: req.From,
			To:   req.To,
			Amount: &invoicer.Amount{
				Amount:   req.Amount.Amount,
				Currency: req.Amount.Currency,
			},
		},
	}, nil
}




func main() {
	connStr := "postgres://Usuario:Contraseñajejejej@localhost:5432/Microservicios?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Cannot create listener")
	}

	server := grpc.NewServer()
	service := &myInvoiceServer{db: db} // ← Aquí pasas la conexión
	invoicer.RegisterInvoicerServer(server, service)

	log.Println("Servidor escuchando en puerto tin...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Muerte eterna")
	}
}
