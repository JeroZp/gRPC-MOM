package transactions_service

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/transactions_service/proto"
)

// Crear una transacción
func CreateTransaction(c *gin.Context) {
	var req struct {
		Amount   int64  `json:"amount" binding:"required"`
		Currency string `json:"currency" binding:"required"`
		From     string `json:"from" binding:"required"`
		To       string `json:"to" binding:"required"`
	}

	// Vincular la solicitud JSON con la estructura
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el objeto CreateRequest para enviar a gRPC
	grpcReq := &proto.CreateRequest{
		Amount: &proto.Amount{
			Amount:   req.Amount,
			Currency: req.Currency,
		},
		From: req.From,
		To:   req.To,
	}

	// Llamar al servicio gRPC para crear la transacción
	resp, err := TransactionClient.Create(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con el resultado de la transacción creada
	c.JSON(http.StatusOK, gin.H{
		"success":   resp.GetSuccess(),
		"message":   resp.GetMessage(),
		"transaction": gin.H{
			"from":     resp.GetTransaction().GetFrom(),
			"to":       resp.GetTransaction().GetTo(),
			"amount":   resp.GetTransaction().GetAmount().GetAmount(),
			"currency": resp.GetTransaction().GetAmount().GetCurrency(),
		},
	})
}
