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

func (enviosHandler *EnviosHandler) CrearEnvio(c *gin.Context) {
	var envio dto.Envio

	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := enviosHandler.enviosService.CrearEnvio(&envio)

	c.JSON(http.StatusOK, resultado)
}

func (enviosHandler *EnviosHandler) ObtenerEnvio(c *gin.Context) {
	envios := enviosHandler.enviosService.ObtenerEnvio()
	c.JSON(http.StatusOK, envios)
}

func (enviosHandler *EnviosHandler) IniciarViajeEnvio(c *gin.Context) {
	id := c.Param("id")

	//verificar que no sea nulo

	resultado := enviosHandler.enviosService.IniciarViajeEnvio(id)

	c.JSON(http.StatusOK, resultado)
}

func (enviosHandler *EnviosHandler) GenerarParadaEnvio(c *gin.Context) {
	id := c.Param("id")

	var parada dto.Parada

	if err := c.Bind(&parada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := enviosHandler.enviosService.GenerarParadaEnvio(id, parada)

	c.JSON(http.StatusOK, resultado)
}

func (EnviosHandler *EnviosHandler) FinalizarViajeEnvio(c *gin.Context) {
	id := c.Param("id")

	var paradaDestino dto.Parada
	if err := c.Bind(&paradaDestino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := EnviosHandler.enviosService.FinalizarViajeEnvio(id, paradaDestino)

	c.JSON(http.StatusOK, resultado)
}
