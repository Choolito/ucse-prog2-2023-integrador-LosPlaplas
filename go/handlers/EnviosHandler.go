package handlers

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/services"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := enviosHandler.enviosService.CrearEnvio(&envio)
	if err != nil {
		if strings.Contains(err.Error(), "el camión no puede estar vacío") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "no se encontró el camión") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "los pedidos no pueden estar vacíos") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "el peso total de los pedidos supera el peso máximo del camión") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Envio creado exitosamente"})
}

func (enviosHandler *EnviosHandler) ObtenerEnvio(c *gin.Context) {
	envios, err := enviosHandler.enviosService.ObtenerEnvio()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, envios)
}

func (enviosHandler *EnviosHandler) IniciarViajeEnvio(c *gin.Context) {
	id := c.Param("id")

	//verificar que no sea nulo

	err := enviosHandler.enviosService.IniciarViajeEnvio(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Envio iniciado en viaje exitosamente"})
}

func (enviosHandler *EnviosHandler) GenerarParadaEnvio(c *gin.Context) {
	id := c.Param("id")

	var parada dto.Parada

	if err := c.Bind(&parada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := enviosHandler.enviosService.GenerarParadaEnvio(id, parada)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Parada exitosamente agregada al envio"})
}

func (EnviosHandler *EnviosHandler) FinalizarViajeEnvio(c *gin.Context) {
	id := c.Param("id")

	var paradaDestino dto.Parada
	if err := c.Bind(&paradaDestino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := EnviosHandler.enviosService.FinalizarViajeEnvio(id, paradaDestino)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Parada exitosamente agregada al envio y finalizada"})
}

func (enviosHandler *EnviosHandler) CambiarEstadoEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio

	err := c.ShouldBindJSON(&envio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operacion, err := enviosHandler.enviosService.CambiarEstadoEnvio(&envio, &user)
	if err != nil || !operacion {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//faltaria trabajar con logging
		//logging.LoggearErrorYResponder(c, "nombre del handler", "Que metodo", err, &user(User no lo tengo aplicado))
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Envio creado exitosamente"})
	//logging.LoggearResultadoYResponder(c, "nombre del handler", "Que metodo", true, &user) Lo mismo aca abajo

}
