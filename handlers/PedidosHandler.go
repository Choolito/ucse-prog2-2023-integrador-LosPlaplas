package handlers

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
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
