package user_service

import (	
	"context"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/google/uuid"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service/client"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service/proto"
)

// Función para registrar un nuevo usuario (REST /register)
func RegisterUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear un nuevo cliente gRPC
	grpcReq := &proto.CreateUserRequest{
		User: &proto.User{
			Id:    uuid.New().String(),
			Name:  req.Name,
			Email: req.Email,
			Credits: 1000,
		},
	}

	// Llamar al servicio gRPC para registrar el usuario
	grpcResp, err := client.UserClient.CreateUser(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta al cliente REST
	c.JSON(http.StatusOK, grpcResp.GetUser())
}

// Función para obtener un usuario por ID (REST /user/:id)
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Llamada gRPC para obtener un usuario
	grpcResp, err := client.UserClient.GetUser(context.Background(), &proto.GetUserRequest{Id: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder al cliente con los datos del usuario
	c.JSON(http.StatusOK, grpcResp.GetUser())
}