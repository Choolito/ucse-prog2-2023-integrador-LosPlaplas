package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Patente            string             `bson:"patente"`
	PesoMaximo         int                `bson:"pesoMaximo"`
	CostoPorKilometro  int                `bson:"costoPorKilometro"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}
