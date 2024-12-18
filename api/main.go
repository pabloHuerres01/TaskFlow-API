package main

import (
	"taskflow-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Crear una instancia del handler
	usuarioHandler := handlers.NewUsuarioHandler()

	// Configurar las rutas
	r.GET("/api/v1/usuarios", usuarioHandler.ObtenerUsuarios)
	r.POST("/api/v1/usuarios", usuarioHandler.CrearUsuario)
	r.DELETE("/api/v1/usuarios/:id", usuarioHandler.EliminarUsuario) // Nueva ruta DELETE

	// Iniciar el servidor
	r.Run(":8080") // Levanta el servidor en el puerto 8080
}
