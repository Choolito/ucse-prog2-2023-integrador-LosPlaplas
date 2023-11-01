package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type PedidosInterface interface {
	//metodos
	CreatePedido(pedido *dto.Pedidos) bool
	GetPedidos() []*dto.Pedidos
	UpdatePedido(id string, pedido *dto.Pedidos) bool
	DeletePedido(id string) bool
	GetPedidosPendientes() []*dto.Pedidos
	UpdatePedidoAceptado(id string) bool
}

type PedidosService struct {
	pedidosRepository  repositories.PedidosRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewPedidosService(pedidosRepository repositories.PedidosRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *PedidosService {
	return &PedidosService{
		pedidosRepository:  pedidosRepository,
		productoRepository: productoRepository,
	}
}

//CRUD de Pedidos

//Metodo que devuelva []AceptadosElegidos que no superen el pesoMaximo --> Pasa a estado "Para Enviar"

func (ps *PedidosService) CreatePedido(pedido *dto.Pedidos) bool {

	var productosCantidad []dto.ProductoCantidad
	for _, producto := range pedido.ListaProductos {
		productoBuscado, _ := ps.productoRepository.GetProductoForID(producto.IDProducto)
		var ProductoCantidad dto.ProductoCantidad
		ProductoCantidad.IDProducto = producto.IDProducto
		ProductoCantidad.CodigoProducto = productoBuscado.CodigoProducto
		ProductoCantidad.TipoProducto = productoBuscado.TipoProducto
		ProductoCantidad.Nombre = productoBuscado.Nombre
		ProductoCantidad.Cantidad = producto.Cantidad
		ProductoCantidad.PrecioUnitario = productoBuscado.PrecioUnitario
		ProductoCantidad.PesoUnitario = productoBuscado.PesoUnitario
		productosCantidad = append(productosCantidad, ProductoCantidad)
	}
	pedido.ListaProductos = productosCantidad
	ps.pedidosRepository.CreatePedido(pedido.GetModel())
	return true
}

func (ps *PedidosService) GetPedidos() []*dto.Pedidos {
	pedidosDB, _ := ps.pedidosRepository.GetPedidos()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos
}

func (ps *PedidosService) UpdatePedido(id string, pedido *dto.Pedidos) bool {
	ps.pedidosRepository.UpdatePedido(id, pedido.GetModel())

	return true
}

func (ps *PedidosService) DeletePedido(id string) bool {
	resultado, _ := ps.pedidosRepository.DeletePedido(id)
	if resultado.ModifiedCount != 0 {
		return true
	}
	return false
}

func (ps *PedidosService) GetPedidosPendientes() []*dto.Pedidos {
	pedidosDB, _ := ps.pedidosRepository.GetPedidosPendientes()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos
}

func (ps *PedidosService) UpdatePedidoAceptado(id string) bool {
	resultado, _ := ps.pedidosRepository.UpdatePedidoAceptado(id)
	if resultado.ModifiedCount != 0 {
		return true
	}
	return false
}
