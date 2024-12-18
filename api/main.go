package main

import (
	"taskflow-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	usuarioHandler := handlers.NewUsuarioHandler()

	// Definir las rutas antes de iniciar el servidor
	r.GET("/api/v1/usuarios", usuarioHandler.ObtenerUsuarios)
	r.POST("/api/v1/usuarios", usuarioHandler.CrearUsuario)

	// Ahora que las rutas están definidas, inicias el servidor
	r.Run(":8080") // Levanta el servidor en el puerto 8080
}
