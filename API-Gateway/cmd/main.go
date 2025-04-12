package main

import (
	"log"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un router Gin
	r := gin.Default()

	// Configura las rutas
	router.SetupRoutes(r)

	log.Println("Iniciando API Gateway en el puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}