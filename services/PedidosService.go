package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type PedidosInterface interface {
	//metodos
	CrearPedido(pedido *dto.Pedidos) bool
	ObtenerPedidos() []*dto.Pedidos
	ActualizarPedido(id string, pedido *dto.Pedidos) bool
	EliminarPedido(id string) bool
	ObtenerPedidosPendientes() []*dto.Pedidos
	ActualizarPedidoAceptado(id string) bool
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

func (ps *PedidosService) CrearPedido(pedido *dto.Pedidos) bool {

	var productosCantidad []dto.ProductoCantidad
	for _, producto := range pedido.ListaProductos {
		productoBuscado, _ := ps.productoRepository.ObtenerProductoPorID(producto.IDProducto)
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
	ps.pedidosRepository.CrearPedido(pedido.GetModel())
	return true
}

func (ps *PedidosService) ObtenerPedidos() []*dto.Pedidos {
	pedidosDB, _ := ps.pedidosRepository.ObtenerPedidos()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos
}

func (ps *PedidosService) ActualizarPedido(id string, pedido *dto.Pedidos) bool {
	ps.pedidosRepository.ActualizarPedido(id, pedido.GetModel())

	return true
}

func (ps *PedidosService) EliminarPedido(id string) bool {
	resultado, _ := ps.pedidosRepository.EliminarPedido(id)
	if resultado.ModifiedCount != 0 {
		return true
	}
	return false
}

func (ps *PedidosService) ObtenerPedidosPendientes() []*dto.Pedidos {
	pedidosDB, _ := ps.pedidosRepository.ObtenerPedidosPendientes()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos
}

func (ps *PedidosService) ActualizarPedidoAceptado(id string) bool {
	resultado, _ := ps.pedidosRepository.ActualizarPedidoAceptado(id)
	if resultado.ModifiedCount != 0 {
		return true
	}
	return false
}
