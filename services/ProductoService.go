package services

import "github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"

type ProductoInterface interface {
	//metodos
}

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{productoRepository: productoRepository}
}

//CRUD de producto
