package dto

import "github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"

type ProductoCantidad struct {
	CodigoProducto string
	TipoProducto   string
	Nombre         string
	Cantidad       int
	PrecioUnitario int
	PesoUnitario   int
}

func NewProductoCantidad(productoCantidad model.ProductoCantidad) *ProductoCantidad {
	return &ProductoCantidad{
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
		CodigoProducto: productoCantidad.CodigoProducto,
		TipoProducto:   productoCantidad.TipoProducto,
		Nombre:         productoCantidad.Nombre,
		Cantidad:       productoCantidad.Cantidad,
		PrecioUnitario: productoCantidad.PrecioUnitario,
		PesoUnitario:   productoCantidad.PesoUnitario,
	}
}

