package main

import (
	"context"
	"fmt"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	UnimplementedPaymentServiceServer
}

// Implementación de ProcessPayment
func (s *server) ProcessPayment(ctx context.Context, req *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	fmt.Printf("Processing payment: %s for user: %s\n", req.GetPaymentId(), req.GetUserId())

	// Simulando procesamiento de pago
	response := &ProcessPaymentResponse{
		PaymentId: req.GetPaymentId(),
		Status:    "Success",
		Message:   "Payment processed successfully.",
	}
	return response, nil
}

// Implementación de GetPaymentStatus
func (s *server) GetPaymentStatus(ctx context.Context, req *GetPaymentStatusRequest) (*GetPaymentStatusResponse, error) {
	fmt.Printf("Getting status for payment: %s\n", req.GetPaymentId())

	// Simulando el estado del pago
	response := &GetPaymentStatusResponse{
		PaymentId: req.GetPaymentId(),
		Status:    "Completed",
	}
	return response, nil
}

func main() {
	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()
	RegisterPaymentServiceServer(grpcServer, &server{})

	// Registrar reflexión (opcional, útil para herramientas de desarrollo)
	reflection.Register(grpcServer)

	// Iniciar el servidor gRPC
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	fmt.Println("Payment service listening on port 50052")
	grpcServer.Serve(lis)
}
