package model

type ProductoCantidad struct {
	CodigoProducto string `bson:"codigoProducto"`
	TipoProducto   string `bson:"tipoProducto"`
	Nombre         string `bson:"nombre"`
	Cantidad       int    `bson:"cantidad"`
	PrecioUnitario int    `bson:"precioUnitario"`
	PesoUnitario   int    `bson:"pesoUnitario"`
}
