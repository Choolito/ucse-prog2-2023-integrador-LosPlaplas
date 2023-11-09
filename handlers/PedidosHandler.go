package handlers

import (
	"net/http"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
	"github.com/gin-gonic/gin"
)

type PedidosHandler struct {
	pedidosService services.PedidosInterface
}

func NewPedidosHandler(pedidosService services.PedidosInterface) *PedidosHandler {
	return &PedidosHandler{
		pedidosService: pedidosService,
	}
}

//CRUD de Pedidos
//Al eliminar un pedido, no se elimina, se pone como cancelado.
//Si el pedido es aceptado, no puede ser cancelado.
//Se valida stock manualmente, y pasa a estado "Aceptado"

func (ph *PedidosHandler) CrearPedido(c *gin.Context) {
	var pedido dto.Pedidos

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ph.pedidosService.CrearPedido(&pedido)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Pedido creado exitosamente"})
}

func (ph *PedidosHandler) ObtenerPedidos(c *gin.Context) {
	pedidos, err := ph.pedidosService.ObtenerPedidos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, pedidos)
}

func (ph *PedidosHandler) ActualizarPedido(c *gin.Context) {
	id := c.Param("id")
	var pedido dto.Pedidos

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ph.pedidosService.ActualizarPedido(id, &pedido)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido actualizado exitosamente"})
}

// Este delete es un put
func (ph *PedidosHandler) EliminarPedido(c *gin.Context) {
	id := c.Param("id")

	err := ph.pedidosService.EliminarPedido(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido eliminado exitosamente"})
}

func (ph *PedidosHandler) ObtenerPedidosPendientes(c *gin.Context) {
	pedidos, err := ph.pedidosService.ObtenerPedidosPendientes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, pedidos)
}
func (ph *PedidosHandler) ActualizarPedidoAceptado(c *gin.Context) {
	id := c.Param("id")

	err := ph.pedidosService.ActualizarPedidoAceptado(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido actualizado a aceptado exitosamente"})
}
