package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type ProductoInterface interface {
	//metodos
	CreateProducto(producto *dto.Producto) bool
	GetProductos() []*dto.Producto
	UpdateProducto(id string, producto *dto.Producto) bool
	DeleteProducto(id string) bool
	GetListStockMinimum() []*dto.Producto
	GetListFiltered(filtro string) []*dto.Producto
}

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{productoRepository: productoRepository}
}

//CRUD de producto

func (ps *ProductoService) CreateProducto(producto *dto.Producto) bool {
	ps.productoRepository.CreateProducto(producto.GetModel())

	return true
}

func (ps *ProductoService) GetProductos() []*dto.Producto {
	productosDB, _ := ps.productoRepository.GetProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos
}

func (ps *ProductoService) UpdateProducto(id string, producto *dto.Producto) bool {
	ps.productoRepository.UpdateProducto(id, producto.GetModel())
	return true
}

func (ps *ProductoService) DeleteProducto(id string) bool {
	ps.productoRepository.DeleteProducto(id)
	return true
}

func (ps *ProductoService) GetListStockMinimum() []*dto.Producto {
	productosDB, _ := ps.productoRepository.GetProductos()

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		if producto.CantidadEnStock < producto.StockMinimo {
			productos = append(productos, producto)
		}
	}

	return productos
}

func (ps *ProductoService) GetListFiltered(filtro string) []*dto.Producto {
	productosDB, _ := ps.productoRepository.GetListFiltered(filtro)

	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(*productoDB)
		productos = append(productos, producto)
	}

	return productos
}
