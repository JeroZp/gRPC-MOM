package stock_trading_server

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/stock_trading_server/proto"
)

// GET /stock/:symbol - obtener precio actual del stock
func GetStockPrice(c *gin.Context) {
	if StockClient == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "StockClient no está inicializado"})
        return
    }
	symbol := c.Param("symbol")

	grpcReq := &proto.StockRequest{StockSymbol: symbol}

	resp, err := StockClient.GetStockPrice(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stock_symbol": resp.GetStockSymbol(),
		"price":        resp.GetPrice(),
		"timestamp":    resp.GetTimestamp(),
	})
}

// POST /stock - crear un nuevo stock
func CreateStock(c *gin.Context) {
	var req struct {
		StockSymbol   string  `json:"stock_symbol" binding:"required"`
		InitialPrice  float64 `json:"initial_price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &proto.CreateStockRequest{
		StockSymbol:  req.StockSymbol,
		InitialPrice: req.InitialPrice,
	}

	resp, err := StockClient.CreateStock(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.GetSuccess(),
		"message": resp.GetMessage(),
	})
}

// PUT /stock - actualizar precio del stock
func UpdateStockPrice(c *gin.Context) {
	var req struct {
		StockSymbol string  `json:"stock_symbol" binding:"required"`
		NewPrice    float64 `json:"new_price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &proto.UpdateStockRequest{
		StockSymbol: req.StockSymbol,
		NewPrice:    req.NewPrice,
	}

	resp, err := StockClient.UpdateStockPrice(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.GetSuccess(),
		"message": resp.GetMessage(),
	})
}
