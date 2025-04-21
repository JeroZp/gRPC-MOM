package user_service

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"  // Para generar un ID único para el usuario
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto" // Paquete generado por protoc
)

// Función para registrar un nuevo usuario (REST /register)
func RegisterUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	// Validar los datos de la solicitud
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear un nuevo ID de usuario único
	userID := uuid.New().String()

	// Construir la solicitud para gRPC
	grpcReq := &proto.CreateUserRequest{
		User: &proto.User{
			Id:      userID,
			Name:    req.Name,
			Email:   req.Email,
			Credits: 1000, // Asignar 100 créditos por defecto
		},
	}

	// Llamada a gRPC al microservicio para crear el usuario
	grpcResp, err := UserClient.CreateUser(context.Background(), grpcReq) // Usar el cliente importado
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta con el nuevo usuario
	c.JSON(http.StatusOK, grpcResp.GetUser())
}

// ListUsers maneja GET /users
func ListUsers(c *gin.Context) {
	resp, err := UserClient.ListUsers(context.Background(), &proto.ListUsersRequest{})
	if err != nil {
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
	}
	c.JSON(http.StatusOK, resp.GetUsers())
  }

// UpdateUser maneja PUT /user
func UpdateUser(c *gin.Context) {
	var req proto.User
	if err := c.ShouldBindJSON(&req); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
	grpcReq := &proto.UpdateUserRequest{User: &req}
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
	resp, err := UserClient.DeleteUser(context.Background(), &proto.DeleteUserRequest{Id: id})
	if err != nil {
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
  }