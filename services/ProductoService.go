package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type ProductoInterface interface {
	//metodos
	CrearProducto(producto *dto.Producto) bool
	ObtenerProductos() []*dto.Producto
	ActualizarProducto(id string, producto *dto.Producto) bool
	EliminarProducto(id string) bool
	ObtenerListaConStockMinimo() []*dto.Producto
	ObtenerListaFiltrada(filtro string) []*dto.Producto
	ObtenerProductoPorID(id string) *dto.Producto
}

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{productoRepository: productoRepository}
}

//CRUD de producto

func (ps *ProductoService) CrearProducto(producto *dto.Producto) bool {
	ps.productoRepository.CrearProducto(producto.GetModel())

	return true
}

func (ps *ProductoService) ObtenerProductos() []*dto.Producto {
	productosDB, _ := ps.productoRepository.ObtenerProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos
}

func (ps *ProductoService) ActualizarProducto(id string, producto *dto.Producto) bool {
	ps.productoRepository.ActualizarProducto(id, producto.GetModel())
	return true
}

func (ps *ProductoService) EliminarProducto(id string) bool {
	ps.productoRepository.EliminarProducto(id)
	return true
}

func (ps *ProductoService) ObtenerListaConStockMinimo() []*dto.Producto {
	productosDB, _ := ps.productoRepository.ObtenerProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		if producto.CantidadEnStock < producto.StockMinimo {
			productos = append(productos, producto)
		}
	}

	return productos
}

func (ps *ProductoService) ObtenerListaFiltrada(filtro string) []*dto.Producto {
	productosDB, _ := ps.productoRepository.ObtenerListaFiltrada(filtro)

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos
}

func (ps *ProductoService) ObtenerProductoPorID(id string) *dto.Producto {
	productoDB, _ := ps.productoRepository.ObtenerProductoPorID(id)

	producto := dto.NewProducto(*productoDB)

	return producto
}
