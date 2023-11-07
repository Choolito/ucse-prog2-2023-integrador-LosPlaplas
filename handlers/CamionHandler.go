package handlers

import (
	"net/http"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ch.camionService.CrearCamion(&camion)

	c.JSON(http.StatusOK, resultado)
}

func (ch *CamionHandler) ObtenerCamiones(c *gin.Context) {
	camiones := ch.camionService.ObtenerCamiones()

	c.JSON(http.StatusOK, camiones)
}

func (ch *CamionHandler) ObtenerCamionPorID(c *gin.Context) {
	id := c.Param("id")

	camion := ch.camionService.ObtenerCamionPorID(id)

	c.JSON(http.StatusOK, camion)
}

func (ch *CamionHandler) ActualizarCamion(c *gin.Context) {
	id := c.Param("id")
	var camion dto.Camion

	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ch.camionService.ActualizarCamion(id, &camion)

	c.JSON(http.StatusOK, resultado)
}

func (ch *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")

	resultado := ch.camionService.EliminarCamion(id)

	c.JSON(http.StatusOK, resultado)
}
