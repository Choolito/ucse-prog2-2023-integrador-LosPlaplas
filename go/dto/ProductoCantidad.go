package dto

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
)

type ProductoCantidad struct {
	IDProducto     string `validate:"required"`
	CodigoProducto string
	Nombre         string
	Cantidad       int `validate:"required,gt=0"`
	PrecioUnitario int
	PesoUnitario   int
}

func NewProductoCantidad(productoCantidad model.ProductoCantidad) *ProductoCantidad {
	return &ProductoCantidad{
		IDProducto:     utils.GetStringIDFromObjectID(productoCantidad.IDProducto),
		CodigoProducto: productoCantidad.CodigoProducto,
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
		Nombre:         productoCantidad.Nombre,
		Cantidad:       productoCantidad.Cantidad,
		PrecioUnitario: productoCantidad.PrecioUnitario,
		PesoUnitario:   productoCantidad.PesoUnitario,
	}
}

//Metodo que sirve para crear un ProductoCantidad para un pedido
