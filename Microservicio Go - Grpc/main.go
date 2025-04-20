package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	// Asegúrate de que este path coincida con el `go_package` o la ubicación de tu archivo `.pb.go`
	invoicer "microservice/Invoicer"
)

type myInvoiceServer struct{
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error){
	response := &invoicer.CreateResponse{
		Success: true,
		Message: "Transacción creada correctamente",
		Transaction: &invoicer.Transaction{
			From: req.From,
			To:   req.To,
			Amount: &invoicer.Amount{
				Amount:   req.Amount.Amount,
				Currency: req.Amount.Currency,
			},
		},
	}

	return response, nil
}



func main(){
	lis, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("Cannot create listener")
	}
	server := grpc.NewServer()
	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(server,service)
	log.Println("Servidor escuchando en puerto tin...")
	err = server.Serve(lis)
	if err != nil{
		log.Fatalf("Muerte eterna")
	}
}