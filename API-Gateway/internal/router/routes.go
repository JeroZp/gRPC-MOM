package router

import (
  "github.com/gin-gonic/gin"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/transaction_service"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/notification_service"
)

func SetupRoutes(r *gin.Engine) {
  r.POST   ("/register", user_service.RegisterUser)
  r.GET    ("/users", user_service.ListUsers)
  r.PUT    ("/user",     user_service.UpdateUser)
  r.DELETE ("/user/:id", user_service.DeleteUser)
  r.POST   ("/transfer", transaction_service.Transfer)
  r.POST   ("/notify", notification_service.Notify)
  r.GET    ("/notifications", notification_service.ListNotifications)
}