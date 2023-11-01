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

func (handler *ProductoHandler) CreateProducto(c *gin.Context) {
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.productoService.CreateProducto(&producto)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) GetProductos(c *gin.Context) {

	resultado := handler.productoService.GetProductos()

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) UpdateProducto(c *gin.Context) {
	id := c.Param("id")
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.productoService.UpdateProducto(id, &producto)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) DeleteProducto(c *gin.Context) {
	id := c.Param("id")

	resultado := handler.productoService.DeleteProducto(id)

	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) GetListStockMinimum(c *gin.Context) {
	resultado := handler.productoService.GetListStockMinimum()

	c.JSON(http.StatusOK, resultado)
}

func (handler *ProductoHandler) GetListFiltered(c *gin.Context) {
	filter := c.Param("filtro")

	resultado := handler.productoService.GetListFiltered(filter)

	c.JSON(http.StatusOK, resultado)
}
