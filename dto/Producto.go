package dto

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type Producto struct {
	ID                 string
	CodigoProducto     string
	TipoProducto       model.TipoProducto //tipos de productos son: golosinas, bebidas, cigarrillos, comestibles, higiene y salud.
	Nombre             string
	PrecioUnitario     int
	PesoUnitario       int
	StockMinimo        int
	CantidadEnStock    int
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		ID:                 utils.GetStringIDFromObjectID(producto.ID),
		CodigoProducto:     producto.CodigoProducto,
		TipoProducto:       producto.TipoProducto,
		Nombre:             producto.Nombre,
		PrecioUnitario:     producto.PrecioUnitario,
		PesoUnitario:       producto.PesoUnitario,
		StockMinimo:        producto.StockMinimo,
		CantidadEnStock:    producto.CantidadEnStock,
		FechaCreacion:      producto.FechaCreacion,
		FechaActualizacion: producto.FechaActualizacion,
	}
}

func (producto Producto) GetModel() model.Producto {
	return model.Producto{
		ID:                 utils.GetObjectIDFromStringID(producto.ID),
		CodigoProducto:     producto.CodigoProducto,
		TipoProducto:       producto.TipoProducto,
		Nombre:             producto.Nombre,
		PrecioUnitario:     producto.PrecioUnitario,
		PesoUnitario:       producto.PesoUnitario,
		StockMinimo:        producto.StockMinimo,
		CantidadEnStock:    producto.CantidadEnStock,
		FechaCreacion:      producto.FechaCreacion,
		FechaActualizacion: producto.FechaActualizacion,
	}
}
