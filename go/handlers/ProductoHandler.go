package handlers

import (
	"net/http"
	"strings"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/services"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := producto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.productoService.CrearProducto(&producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Producto creado exitosamente"})
}

func (handler *ProductoHandler) CrearProductos(c *gin.Context) {
	var productos []dto.Producto

	if err := c.ShouldBindJSON(&productos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	for _, producto := range productos {
		if err := producto.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var productosPtr []*dto.Producto
	for i := range productos {
		productosPtr = append(productosPtr, &productos[i])
	}
	err := handler.productoService.CrearProductos(productosPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Productos creados exitosamente"})
}

func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {

	resultado, err := handler.productoService.ObtenerProductos()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, resultado)

}

func (handler *ProductoHandler) ActualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.productoService.ActualizarProducto(id, &producto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto actualizado exitosamente"})

}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")

	err := handler.productoService.EliminarProducto(id)
	if err != nil {
		if strings.Contains(err.Error(), "no se encontró el producto") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	// En lugar de 204 No Content, devuelve 200 con un mensaje
	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado exitosamente"})
}

func (handler *ProductoHandler) ObtenerListaConStockMinimo(c *gin.Context) {
	var filtro utils.FiltroProducto

	// Bindear el JSON del body al filtro
	if err := c.ShouldBindJSON(&filtro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de entrada inválido"})
		return
	}

	// Llamar al servicio con el filtro
	resultado, err := handler.productoService.ObtenerListaConStockMinimo(filtro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultado)
}
func (handler *ProductoHandler) ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")

	resultado, err := handler.productoService.ObtenerProductoPorID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resultado)
}
