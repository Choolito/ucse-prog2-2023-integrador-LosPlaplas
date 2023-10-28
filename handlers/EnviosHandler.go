package handlers

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
)

type EnviosHandler struct {
	enviosService services.EnviosInterface
}

func NewEnviosHandler(enviosService services.EnviosInterface) *EnviosHandler {
	return &EnviosHandler{
		enviosService: enviosService,
	}
}

//Post Camionero inicia viaje --> envio estado "En ruta"
//Las paradas las puede ir haciendo y se guarda ciudad y km recorridos desde ultima parada/inicio.
//Metodo generarParadas

func (enviosHandler *EnviosHandler) CreateEnvio(c *gin.Context) {
	var envio dto.Envio

	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := enviosHandler.enviosService.CreateEnvio(&envio)

	c.JSON(http.StatusOK, resultado)
}

func (enviosHandler *EnviosHandler) StartTrip(c *gin.Context) {
	id := c.Param("id")

	//verificar que no sea nulo

	resultado := enviosHandler.enviosService.StartTrip(id)

	c.JSON(http.StatusOK, resultado)
}

func (enviosHandler *EnviosHandler) GenerateStop(c *gin.Context) {
	id := c.Param("id")

	var parada dto.Parada

	//verificar que no sea nulo

	resultado := enviosHandler.enviosService.GenerateStop(id, parada)

	c.JSON(http.StatusOK, resultado)
}
