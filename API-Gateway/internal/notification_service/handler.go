package notification_service

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

type notifyReq struct {
    UserID  string `json:"user_id" binding:"required"`
    Title   string `json:"title" binding:"required"`
    Message string `json:"message" binding:"required"`
}

// Notify invoca directamente al NotificationService gRPC
func Notify(c *gin.Context) {
    var req notifyReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    resp, err := NotifClient.Notify(context.Background(), &pb.NotifyRequest{
        UserId:  req.UserID,
        Title:   req.Title,
        Message: req.Message,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}

// ListNotifications llama gRPC directo para obtener todas las notificaciones
func ListNotifications(c *gin.Context) {
    resp, err := NotifClient.ListNotifications(context.Background(), &pb.ListNotificationsRequest{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp.GetNotifications())
}
