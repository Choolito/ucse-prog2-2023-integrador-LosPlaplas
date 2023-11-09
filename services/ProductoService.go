package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type ProductoInterface interface {
	//metodos
	CrearProducto(producto *dto.Producto) error
	ObtenerProductos() ([]*dto.Producto, error)
	ActualizarProducto(id string, producto *dto.Producto) error
	EliminarProducto(id string) error
	ObtenerListaConStockMinimo() ([]*dto.Producto, error)
	ObtenerListaFiltrada(filtro string) ([]*dto.Producto, error)
	ObtenerProductoPorID(id string) (*dto.Producto, error)
}

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{productoRepository: productoRepository}
}

//CRUD de producto

func (ps *ProductoService) CrearProducto(producto *dto.Producto) error {
	_, err := ps.productoRepository.CrearProducto(producto.GetModel())
	return err
}

func (ps *ProductoService) ObtenerProductos() ([]*dto.Producto, error) {
	productosDB, err := ps.productoRepository.ObtenerProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos, err
}

func (ps *ProductoService) ActualizarProducto(id string, producto *dto.Producto) error {
	_, err := ps.productoRepository.ActualizarProducto(id, producto.GetModel())
	return err
}

func (ps *ProductoService) EliminarProducto(id string) error {
	_, err := ps.productoRepository.EliminarProducto(id)
	return err
}

func (ps *ProductoService) ObtenerListaConStockMinimo() ([]*dto.Producto, error) {
	productosDB, err := ps.productoRepository.ObtenerProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		if producto.CantidadEnStock < producto.StockMinimo {
			productos = append(productos, producto)
		}
	}

	return productos, err
}

func (ps *ProductoService) ObtenerListaFiltrada(filtro string) ([]*dto.Producto, error) {
	productosDB, err := ps.productoRepository.ObtenerListaFiltrada(filtro)

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos, err
}

func (ps *ProductoService) ObtenerProductoPorID(id string) (*dto.Producto, error) {
	productoDB, err := ps.productoRepository.ObtenerProductoPorID(id)

	producto := dto.NewProducto(*productoDB)

	return producto, err
}
