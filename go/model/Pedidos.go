package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedidos struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	ListaProductos      []ProductoCantidad `bson:"listaProductos"`
	CiudadDestinoPedido string             `bson:"ciudadDestinoPedido"`
	EstadoPedido        EstadoPedido       `bson:"estadoPedido"`
	FechaCreacion       time.Time          `bson:"fechaCreacion"`
	FechaActualizacion  time.Time          `bson:"fechaActualizacion"`
}
