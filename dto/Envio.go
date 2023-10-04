package dto

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type Envio struct {
	ID                 string
	IDCamion           string
	Pedidos            []string
	Paradas            []Parada
	Estado             string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		ID:                 utils.GetStringIDFromObjectID(envio.ID),
		IDCamion:           utils.GetStringIDFromObjectID(envio.IDCamion),
		Pedidos:            envio.Pedidos,
		Paradas:            listaParadasDTO(envio.Paradas),
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
func listaParadasDTO(listaIngresante []model.Parada) []Parada {
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
		Paradas:            listaParadasModel(envio.Paradas),
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
func listaParadasModel(listaIngresante []Parada) []model.Parada {
	var listaSaliente []model.Parada
	for _, parada := range listaIngresante {
		listaSaliente = append(listaSaliente, parada.GetModel())
	}
	return listaSaliente
}
