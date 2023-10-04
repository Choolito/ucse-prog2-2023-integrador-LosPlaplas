package dto

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type Pedidos struct {
	ID                  string
	ListaProductos      []ProductoCantidad
	CiudadDestinoPedido string
	EstadoPedido        string
	FechaCreacion       time.Time
	FechaActualizacion  time.Time
}

func NewPedidos(pedidos model.Pedidos) *Pedidos {
	return &Pedidos{
		ID:                  utils.GetStringIDFromObjectID(pedidos.ID),
		ListaProductos:      listaProductoCantidadDTO(pedidos.ListaProductos),
		CiudadDestinoPedido: pedidos.CiudadDestinoPedido,
		EstadoPedido:        pedidos.EstadoPedido,
		FechaCreacion:       pedidos.FechaCreacion,
		FechaActualizacion:  pedidos.FechaActualizacion,
	}
}

func listaProductoCantidadDTO(listaIngresante []model.ProductoCantidad) []ProductoCantidad {
	var listaSaliente []ProductoCantidad
	for _, producto := range listaIngresante {
		listaSaliente = append(listaSaliente, *NewProductoCantidad(producto))
	}
	return listaSaliente
}

func (pedidos Pedidos) GetModel() model.Pedidos {
	return model.Pedidos{
		ID:                  utils.GetObjectIDFromStringID(pedidos.ID),
		ListaProductos:      listaProductoCantidadModel(pedidos.ListaProductos),
		CiudadDestinoPedido: pedidos.CiudadDestinoPedido,
		EstadoPedido:        pedidos.EstadoPedido,
		FechaCreacion:       pedidos.FechaCreacion,
		FechaActualizacion:  pedidos.FechaActualizacion,
	}
}
func listaProductoCantidadModel(listaIngresante []ProductoCantidad) []model.ProductoCantidad {
	var listaSaliente []model.ProductoCantidad
	for _, producto := range listaIngresante {
		listaSaliente = append(listaSaliente, producto.GetModel())
	}
	return listaSaliente
}
