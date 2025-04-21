package router

import (
  "github.com/gin-gonic/gin"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/stock_trading_server"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/transactions_service"

)

func SetupRoutes(r *gin.Engine) {
  r.POST   ("/register", user_service.RegisterUser)
  r.GET    ("/users", user_service.ListUsers)
  r.PUT    ("/user",     user_service.UpdateUser)
  r.DELETE ("/user/:id", user_service.DeleteUser)
  r.GET    ("/stock/:symbol", stock_trading_server.GetStockPrice)   
	r.POST   ("/stock", stock_trading_server.CreateStock)            
	r.PUT    ("/stock", stock_trading_server.UpdateStockPrice)
  r.POST   ("/transaction", transactions_service.CreateTransaction)
        
}
