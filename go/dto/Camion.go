package dto

import (
	"fmt"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"github.com/go-playground/validator/v10"
)

type Camion struct {
	ID                 string    `json:"id"`
	Patente            string    `json:"patente" validate:"required"`
	PesoMaximo         int       `json:"pesoMaximo" validate:"required,gt=0"`
	CostoPorKilometro  int       `json:"costoPorKilometro" validate:"required,gt=0"`
	FechaCreacion      time.Time `json:"fechaCreacion"`
	FechaActualizacion time.Time `json:"fechaActualizacion"`
}

func (c *Camion) Validate() error {
	err := utils.Validate.Struct(c)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Patente":
				return fmt.Errorf("la patente del camión es obligatoria")
			case "PesoMaximo":
				return fmt.Errorf("el peso máximo es obligatorio y debe ser un número positivo")
			case "CostoPorKilometro":
				return fmt.Errorf("el costo por kilómetro es obligatorio y debe ser un número positivo")
			}
		}
	}
	return nil
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
