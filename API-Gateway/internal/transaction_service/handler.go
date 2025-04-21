package transaction_service

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

// TransferReq define cómo llega la solicitud REST de transferencia
type TransferReq struct {
	FromId string `json:"from_id" binding:"required"`
	ToId   string `json:"to_id"   binding:"required"`
	Amount int32  `json:"amount"  binding:"required,gt=0"`
}

// Transfer maneja POST /transfer y llama al TransactionService gRPC
func Transfer(c *gin.Context) {
	var req TransferReq
	// 1. Parsear y validar JSON de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Construir el mensaje gRPC
	grpcReq := &pb.TransferRequest{
		FromId: req.FromId,
		ToId:   req.ToId,
		Amount: req.Amount,
	}

	// 3. Llamada gRPC al TransactionService
	grpcResp, err := TxClient.Transfer(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4. Devolver directamente la respuesta gRPC en JSON
	c.JSON(http.StatusOK, grpcResp)
}