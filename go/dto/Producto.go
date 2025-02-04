package dto

import (
	"fmt"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"github.com/go-playground/validator/v10"
)

type Producto struct {
	ID                 string             `json:"id"`
	CodigoProducto     string             `json:"codigoProducto" validate:"required"`
	TipoProducto       model.TipoProducto `json:"tipoProducto" validate:"required"`
	Nombre             string             `json:"nombre" validate:"required"`
	PrecioUnitario     int                `json:"precioUnitario" validate:"required,gt=0"`
	PesoUnitario       int                `json:"pesoUnitario" validate:"required,gt=0"`
	StockMinimo        int                `json:"stockMinimo" validate:"required,gt=0"`
	CantidadEnStock    int                `json:"cantidadEnStock" validate:"required,gt=0"`
	FechaCreacion      time.Time          `json:"fechaCreacion"`
	FechaActualizacion time.Time          `json:"fechaActualizacion"`
}

func (p *Producto) Validate() error {
	err := utils.Validate.Struct(p)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "CodigoProducto":
				return fmt.Errorf("el código del producto es obligatorio")
			case "TipoProducto":
				return fmt.Errorf("el tipo de producto es obligatorio")
			case "Nombre":
				return fmt.Errorf("el nombre del producto es obligatorio")
			case "PrecioUnitario":
				return fmt.Errorf("el precio unitario es obligatorio y debe ser un número positivo")
			case "PesoUnitario":
				return fmt.Errorf("el peso unitario es obligatorio y debe ser un número positivo")
			case "StockMinimo":
				return fmt.Errorf("el stock mínimo es obligatorio y debe ser un número positivo")
			case "CantidadEnStock":
				return fmt.Errorf("la cantidad en stock es obligatoria y debe ser un número positivo")
			}
		}
	}
	return nil
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
