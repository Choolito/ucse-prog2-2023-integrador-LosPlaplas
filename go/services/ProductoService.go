package services

import (
	"fmt"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/repositories"
)

type ProductoInterface interface {
	// Métodos
	CrearProducto(producto *dto.Producto) error
	CrearProductos(productos []*dto.Producto) error
	ObtenerProductos() ([]*dto.Producto, error)
	ActualizarProducto(id string, producto *dto.Producto) error
	EliminarProducto(id string) error
	ObtenerListaConStockMinimo() ([]*dto.Producto, error)
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

func (ps *ProductoService) CrearProductos(productos []*dto.Producto) error {
	for _, producto := range productos {
		_, err := ps.productoRepository.CrearProductos([]model.Producto{producto.GetModel()})
		if err != nil {
			return err
		}
	}
	return nil
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
	// Verificar si el producto existe
	_, err := ps.productoRepository.ObtenerProductoPorID(id)
	if err != nil {
		return fmt.Errorf("no se encontró el producto con el id: %s", id)
	}

	// "Eliminar" el producto (marcar como eliminado)
	err = ps.productoRepository.EliminarProducto(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductoService) ObtenerListaConStockMinimo() ([]*dto.Producto, error) {
	// Lógica para obtener productos con stock mínimo
	productos, err := service.productoRepository.ObtenerListaConStockMinimo()
	if err != nil {
		return nil, err
	}
	var productosDTO []*dto.Producto
	for _, producto := range productos {
		productosDTO = append(productosDTO, dto.NewProducto(*producto))
	}
	return productosDTO, nil
}

func (ps *ProductoService) ObtenerProductoPorID(id string) (*dto.Producto, error) {
	productoDB, err := ps.productoRepository.ObtenerProductoPorID(id)

	producto := dto.NewProducto(*productoDB)

	return producto, err
}
