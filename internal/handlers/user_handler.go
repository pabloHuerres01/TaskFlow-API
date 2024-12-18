package handlers

import (
	"net/http"
	"taskflow-api/internal/model"
	"taskflow-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service *services.UsuarioService
}

// NewUsuarioHandler crea una nueva instancia de UsuarioHandler
func NewUsuarioHandler() *UsuarioHandler {
	return &UsuarioHandler{service: services.NewUsuarioService()}
}

// ObtenerUsuarios maneja la solicitud GET para obtener todos los usuarios
func (h *UsuarioHandler) ObtenerUsuarios(c *gin.Context) {
	usuarios, err := h.service.ObtenerUsuarios(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// CrearUsuario maneja la solicitud POST para crear un nuevo usuario
func (h *UsuarioHandler) CrearUsuario(c *gin.Context) {
	var usuario model.Usuario

	// Parsear el cuerpo de la solicitud
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear el usuario
	err := h.service.CrearUsuario(c.Request.Context(), usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado correctamente"})
}

// EliminarUsuario maneja la solicitud DELETE para eliminar un usuario
func (h *UsuarioHandler) EliminarUsuario(c *gin.Context) {
	id := c.Param("id")

	// Llamar al servicio para eliminar el usuario
	err := h.service.EliminarUsuario(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
