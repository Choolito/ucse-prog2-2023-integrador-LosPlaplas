package handlers

import (
	"net/http"
	"strings"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/services"
	"github.com/gin-gonic/gin"
)

type CamionHandler struct {
	camionService services.CamionInterface
}

func NewCamionHandler(camionService services.CamionInterface) *CamionHandler {
	return &CamionHandler{
		camionService: camionService,
	}
}

//CRUD de Camion

func (ch *CamionHandler) CrearCamion(c *gin.Context) {
	var camion dto.Camion

	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := camion.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ch.camionService.CrearCamion(&camion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Camión creado exitosamente"})
}

func (ch *CamionHandler) ObtenerCamiones(c *gin.Context) {
	//user

	camiones, err := ch.camionService.ObtenerCamiones()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, camiones)
}

func (ch *CamionHandler) ObtenerCamionPorID(c *gin.Context) {
	id := c.Param("id")

	camion, err := ch.camionService.ObtenerCamionPorID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, camion)
}

func (ch *CamionHandler) ActualizarCamion(c *gin.Context) {
	id := c.Param("id")
	var camion dto.Camion

	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := ch.camionService.ActualizarCamion(id, &camion)
	if err != nil {
		if strings.Contains(err.Error(), "no se encontró el camión") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Camión actualizado exitosamente"})
}

func (ch *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")

	err := ch.camionService.EliminarCamion(id)
	if err != nil {
		if strings.Contains(err.Error(), "no se encontró el camión") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"mensaje": "Camión eliminado exitosamente"})
}
