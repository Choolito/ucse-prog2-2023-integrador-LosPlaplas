package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	IDCamion           primitive.ObjectID `bson:"idCamion,omitempty"`
	Pedidos            []string           `bson:"pedidos"`
	Paradas            []Parada           `bson:"paradas"`
	Estado             string             `bson:"estado"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}
