package handlers

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
)

type ProductoHandler struct {
	productoService services.ProductoInterface
}

func NewProductoHandler(productoService services.ProductoInterface) *ProductoHandler {
	return &ProductoHandler{
		productoService: productoService,
	}
}

//CRUD de Producto
