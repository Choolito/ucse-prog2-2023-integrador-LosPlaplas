package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	IDCamion           primitive.ObjectID `bson:"idCamion,omitempty"`
	Pedidos            []string
	Paradas            []Parada
	Estado             string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
