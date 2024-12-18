package handlers

import (
	"net/http"
	models "taskflow-api/internal/model"
	"taskflow-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service *services.UsuarioService
}

func NewUsuarioHandler() *UsuarioHandler {
	return &UsuarioHandler{service: services.NewUsuarioService()}
}

func (h *UsuarioHandler) ObtenerUsuarios(c *gin.Context) {
	usuarios, err := h.service.ObtenerUsuarios(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

/*
	func (h *UsuarioHandler) CrearUsuario(c *gin.Context) {
		var usuario models.Usuario

		// Parsear el cuerpo de la petición
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
*/
func (h *UsuarioHandler) CrearUsuario(c *gin.Context) {
	var usuario models.Usuario

	// Parsear el cuerpo de la petición
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
