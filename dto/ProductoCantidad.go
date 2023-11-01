package dto

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type ProductoCantidad struct {
	IDProducto     string
	CodigoProducto string
	TipoProducto   string
	Nombre         string
	Cantidad       int
	PrecioUnitario int
	PesoUnitario   int
}

func NewProductoCantidad(productoCantidad model.ProductoCantidad) *ProductoCantidad {
	return &ProductoCantidad{
		IDProducto:     utils.GetStringIDFromObjectID(productoCantidad.IDProducto),
		CodigoProducto: productoCantidad.CodigoProducto,
		TipoProducto:   productoCantidad.TipoProducto,
		Nombre:         productoCantidad.Nombre,
		Cantidad:       productoCantidad.Cantidad,
		PrecioUnitario: productoCantidad.PrecioUnitario,
		PesoUnitario:   productoCantidad.PesoUnitario,
	}
}

func (productoCantidad ProductoCantidad) GetModel() model.ProductoCantidad {
	return model.ProductoCantidad{
		IDProducto:     utils.GetObjectIDFromStringID(productoCantidad.IDProducto),
		CodigoProducto: productoCantidad.CodigoProducto,
		TipoProducto:   productoCantidad.TipoProducto,
		Nombre:         productoCantidad.Nombre,
		Cantidad:       productoCantidad.Cantidad,
		PrecioUnitario: productoCantidad.PrecioUnitario,
		PesoUnitario:   productoCantidad.PesoUnitario,
	}
}
