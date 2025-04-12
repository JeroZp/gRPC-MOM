package user_service

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"  // Para generar un ID único para el usuario
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service/proto" // Paquete generado por protoc
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
			Credits: 100, // Asignar 100 créditos por defecto
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

// Función para obtener la información de un usuario por ID (REST /user/:id)
func GetUser(c *gin.Context) {
	userID := c.Param("id") // Obtener el ID del usuario desde los parámetros de la URL
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	// Llamada gRPC al microservicio para obtener los detalles del usuario
	grpcResp, err := UserClient.GetUser(context.Background(), &proto.GetUserRequest{Id: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver los detalles del usuario
	c.JSON(http.StatusOK, grpcResp.GetUser())
}
