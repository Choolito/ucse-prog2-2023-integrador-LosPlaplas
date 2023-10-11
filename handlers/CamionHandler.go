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

func (ch *CamionHandler) CreateCamion(c *gin.Context) {
	var camion dto.Camion

	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ch.camionService.CreateCamion(&camion)

	c.JSON(http.StatusOK, resultado)
}

func (ch *CamionHandler) GetCamiones(c *gin.Context) {
	camiones := ch.camionService.GetCamiones()

	c.JSON(http.StatusOK, camiones)
}

func (ch *CamionHandler) UpdateCamion(c *gin.Context) {
	id := c.Param("id")
	var camion dto.Camion

	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ch.camionService.UpdateCamion(id, &camion)

	c.JSON(http.StatusOK, resultado)
}

func (ch *CamionHandler) DeleteCamion(c *gin.Context) {
	id := c.Param("id")

	resultado := ch.camionService.DeleteCamion(id)

	c.JSON(http.StatusOK, resultado)
}
