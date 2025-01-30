package dto

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
)

type Camion struct {
	ID                 string
	Patente            string
	PesoMaximo         int
	CostoPorKilometro  int
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		ID:                 utils.GetStringIDFromObjectID(camion.ID),
		Patente:            camion.Patente,
		PesoMaximo:         camion.PesoMaximo,
		CostoPorKilometro:  camion.CostoPorKilometro,
		FechaCreacion:      camion.FechaCreacion,
		FechaActualizacion: camion.FechaActualizacion,
	}
}

func (camion Camion) GetModel() model.Camion {
	return model.Camion{
		ID:                 utils.GetObjectIDFromStringID(camion.ID),
		Patente:            camion.Patente,
		PesoMaximo:         camion.PesoMaximo,
		CostoPorKilometro:  camion.CostoPorKilometro,
		FechaCreacion:      camion.FechaCreacion,
		FechaActualizacion: camion.FechaActualizacion,
	}
}
