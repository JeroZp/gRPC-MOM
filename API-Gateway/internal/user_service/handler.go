package user_service

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/mom_service"
    pb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
)

type registerReq struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

// RegisterUser encola la creación de usuario en MOM (“user_ops”)
func RegisterUser(c *gin.Context) {
    var req registerReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // serializamos y publicamos
    data, _ := json.Marshal(req)
    resp, err := mom_service.MOMClient.Publish(context.Background(), &pb.PublishRequest{
        Queue: "user_ops",
        Data:  data,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // devolvemos 202 Accepted + ID de encolado
    c.JSON(http.StatusAccepted, gin.H{"enqueued_id": resp.GetId()})
}

// GetUser hace una llamada gRPC directa al UserService para obtener un usuario
func GetUser(c *gin.Context) {
    id := c.Param("id")
    resp, err := UserClient.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp.GetUser())
}

// UpdateUser maneja PUT /user
func UpdateUser(c *gin.Context) {
	var req pb.User
	if err := c.ShouldBindJSON(&req); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
	grpcReq := &pb.UpdateUserRequest{User: &req}
	resp, err := UserClient.UpdateUser(context.Background(), grpcReq)
	if err != nil {
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
	}
	c.JSON(http.StatusOK, resp.GetUser())
  }

// DeleteUser maneja DELETE /user/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := UserClient.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: id})
	if err != nil {
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
  }

// ListUsers llama gRPC directo para obtener todos los usuarios
func ListUsers(c *gin.Context) {
    resp, err := UserClient.ListUsers(context.Background(), &pb.ListUsersRequest{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp.GetUsers())
}