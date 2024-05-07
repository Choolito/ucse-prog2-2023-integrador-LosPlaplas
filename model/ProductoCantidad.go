package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductoCantidad struct {
	IDProducto     primitive.ObjectID `bson:"idProducto,omitempty"`
	CodigoProducto string             `bson:"codigoProducto"`
	Nombre         string             `bson:"nombre"`
	Cantidad       int                `bson:"cantidad"`
	PrecioUnitario int                `bson:"precioUnitario"`
	PesoUnitario   int                `bson:"pesoUnitario"`
}
