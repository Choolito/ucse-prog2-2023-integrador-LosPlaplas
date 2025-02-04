package dto

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
)

type Envio struct {
	ID                 string
	IDCamion           string   `validate:"required"`
	Pedidos            []string `validate:"required"`
	Paradas            []Parada
	Estado             model.EstadoEnvio
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		ID:                 utils.GetStringIDFromObjectID(envio.ID),
		IDCamion:           utils.GetStringIDFromObjectID(envio.IDCamion),
		Pedidos:            envio.Pedidos,
		Paradas:            newParadas(envio.Paradas),
		Estado:             envio.Estado,
		FechaCreacion:      envio.FechaCreacion,
		FechaActualizacion: envio.FechaActualizacion,
	}
}

/*
	func listaPedidosDTO(listaIngresante []model.Pedidos) []Pedidos {
		var listaSaliente []Pedidos
		for _, pedido := range listaIngresante {
			listaSaliente = append(listaSaliente, *NewPedidos(pedido))
		}
		return listaSaliente
	}
*/
func newParadas(listaIngresante []model.Parada) []Parada {
	var listaSaliente []Parada
	for _, parada := range listaIngresante {
		listaSaliente = append(listaSaliente, *NewParada(parada))
	}
	return listaSaliente
}

func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		ID:                 utils.GetObjectIDFromStringID(envio.ID),
		IDCamion:           utils.GetObjectIDFromStringID(envio.IDCamion),
		Pedidos:            envio.Pedidos,
		Paradas:            envio.getParadas(),
		Estado:             envio.Estado,
		FechaCreacion:      envio.FechaCreacion,
		FechaActualizacion: envio.FechaActualizacion,
	}
}

/*
	func listaPedidosModel(listaIngresante []Pedidos) []model.Pedidos {
		var listaSaliente []model.Pedidos
		for _, pedido := range listaIngresante {
			listaSaliente = append(listaSaliente, pedido.GetModel())
		}
		return listaSaliente
	}
*/
func (envio Envio) getParadas() []model.Parada {
	var listaSaliente []model.Parada
	for _, parada := range envio.Paradas {
		listaSaliente = append(listaSaliente, parada.GetModel())
	}
	return listaSaliente
}
