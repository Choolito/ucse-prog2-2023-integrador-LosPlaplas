package handlers

import (
	"net/http"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
	"github.com/gin-gonic/gin"
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

func (handler *ProductoHandler) CrearProducto(c *gin.Context) {
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.productoService.CrearProducto(&producto)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {

	resultado := handler.productoService.ObtenerProductos()

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) ActualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.productoService.ActualizarProducto(id, &producto)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")

	resultado := handler.productoService.EliminarProducto(id)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) ObtenerListaConStockMinimo(c *gin.Context) {

	resultado := handler.productoService.ObtenerListaConStockMinimo()

	c.JSON(http.StatusOK, resultado)
}

func (handler *ProductoHandler) ObtenerListaFiltrada(c *gin.Context) {
	filter := c.Param("filtro")

	resultado := handler.productoService.ObtenerListaFiltrada(filter)

	c.JSON(http.StatusOK, resultado)
}

func (handler *ProductoHandler) ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")

	resultado := handler.productoService.ObtenerProductoPorID(id)

	c.JSON(http.StatusOK, resultado)
}
