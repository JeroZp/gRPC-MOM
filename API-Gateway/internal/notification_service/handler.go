package notification_service

import (
  "context"
  "net/http"

  "github.com/gin-gonic/gin"
  userpb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service"
)

type NotifyReq struct {
  UserID  string `json:"user_id" binding:"required"`
  Title   string `json:"title" binding:"required"`
  Message string `json:"message" binding:"required"`
}

func Notify(c *gin.Context) {
  var in NotifyReq
  if err := c.ShouldBindJSON(&in); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  resp, err := NotifClient.Notify(context.Background(), &userpb.NotifyRequest{
    UserId:  in.UserID,
    Title:   in.Title,
    Message: in.Message,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, resp)
}

func ListNotifications(c *gin.Context) {
  // 1) Listar desde NotificationService
  nresp, err := NotifClient.ListNotifications(context.Background(), &userpb.ListNotificationsRequest{})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  // Para cada notificación, obtener datos de usuario
  type out struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Title   string `json:"title"`
    Message string `json:"message"`
  }
  var result []out

  for _, n := range nresp.GetNotifications() {
    // Llamada gRPC a UserService
    uresp, err := user_service.UserClient.GetUser(context.Background(), &userpb.GetUserRequest{Id: n.GetUserId()})
    if err != nil {
      // si falla, aún agregamos la notificación con id + mensaje
      result = append(result, out{
        ID:      n.GetUserId(),
        Title:   n.GetTitle(),
        Message: n.GetMessage(),
      })
      continue
    }
    u := uresp.GetUser()
    result = append(result, out{
      ID:      u.GetId(),
      Name:    u.GetName(),
      Email:   u.GetEmail(),
      Title:   n.GetTitle(),
      Message: n.GetMessage(),
    })
  }

  c.JSON(http.StatusOK, result)
}