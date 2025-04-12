package router

import (
	"github.com/gin-gonic/gin"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service"
)

// SetupRoutes define todas las rutas REST
func SetupRoutes(r *gin.Engine) {
	// Rutas REST
	r.POST("/register", user_service.RegisterUser)
	r.GET("/user/:id", user_service.GetUser)
}
