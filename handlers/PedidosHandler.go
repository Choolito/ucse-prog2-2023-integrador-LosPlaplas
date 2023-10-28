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

func (ph *PedidosHandler) CreatePedido(c *gin.Context) {
	var pedido dto.Pedidos

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ph.pedidosService.CreatePedido(&pedido)

	c.JSON(http.StatusOK, resultado)
}

func (ph *PedidosHandler) GetPedidos(c *gin.Context) {
	pedidos := ph.pedidosService.GetPedidos()
	c.JSON(http.StatusOK, pedidos)
}

func (ph *PedidosHandler) UpdatePedido(c *gin.Context) {
	id := c.Param("id")
	var pedido dto.Pedidos

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := ph.pedidosService.UpdatePedido(id, &pedido)

	c.JSON(http.StatusOK, resultado)
}

// Este delete es un put
func (ph *PedidosHandler) DeletePedido(c *gin.Context) {
	id := c.Param("id")

	resultado := ph.pedidosService.DeletePedido(id)

	c.JSON(http.StatusOK, resultado)
}

func (ph *PedidosHandler) GetPedidosPendientes(c *gin.Context) {
	pedidos := ph.pedidosService.GetPedidosPendientes()
	c.JSON(http.StatusOK, pedidos)
}
func (ph *PedidosHandler) UpdatePedidoAceptado(c *gin.Context) {
	id := c.Param("id")

	resultado := ph.pedidosService.UpdatePedidoAceptado(id)

	c.JSON(http.StatusOK, resultado)
}
