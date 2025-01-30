package handlers

import (
	"net/http"
	"strings"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/services"
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

	err := handler.productoService.CrearProducto(&producto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Producto creado exitosamente"})

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado exitosamente"})

}

func (handler *ProductoHandler) ObtenerListaConStockMinimo(c *gin.Context) {
	// Obtener la lista de productos con stock mínimo
	resultado, err := handler.productoService.ObtenerListaConStockMinimo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener parámetro de consulta
	categoria := c.Query("categoria")

	// Filtrar la lista de productos si el parámetro de consulta está presente
	var productosFiltrados []dto.Producto
	if categoria != "" {
		categoria = strings.ToLower(categoria)
		for _, producto := range resultado {
			if strings.ToLower(string(producto.TipoProducto)) == categoria {
				productosFiltrados = append(productosFiltrados, *producto)
			}
		}
	} else {
		for _, producto := range resultado {
			productosFiltrados = append(productosFiltrados, *producto)
		}
	}

	c.JSON(http.StatusOK, productosFiltrados)
}

func (handler *ProductoHandler) ObtenerListaFiltrada(c *gin.Context) {
	filter := c.Param("filtro")

	resultado, err := handler.productoService.ObtenerListaFiltrada(filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
