package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto     string             `bson:"codigoProducto"`
	TipoProducto       TipoProducto       `bson:"tipoProducto"`
	Nombre             string             `bson:"nombre"`
	PrecioUnitario     int                `bson:"precioUnitario"`
	PesoUnitario       int                `bson:"pesoUnitario"`
	StockMinimo        int                `bson:"stockMinimo"`
	CantidadEnStock    int                `bson:"cantidadEnStock"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}
