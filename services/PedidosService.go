package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type PedidosInterface interface {
	//metodos
	CrearPedido(pedido *dto.Pedidos) error
	ObtenerPedidos() ([]*dto.Pedidos, error)
	EliminarPedido(id string) error
	ObtenerPedidosPendientes() ([]*dto.Pedidos, error)
	ActualizarPedidoAceptado(id string) error
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

func (ps *PedidosService) CrearPedido(pedido *dto.Pedidos) error {

	var productosCantidad []dto.ProductoCantidad
	for _, producto := range pedido.ListaProductos {
		productoBuscado, _ := ps.productoRepository.ObtenerProductoPorID(producto.IDProducto)
		var ProductoCantidad dto.ProductoCantidad
		ProductoCantidad.IDProducto = producto.IDProducto
		ProductoCantidad.CodigoProducto = productoBuscado.CodigoProducto
		ProductoCantidad.Nombre = productoBuscado.Nombre
		ProductoCantidad.Cantidad = producto.Cantidad
		ProductoCantidad.PrecioUnitario = productoBuscado.PrecioUnitario
		ProductoCantidad.PesoUnitario = productoBuscado.PesoUnitario
		productosCantidad = append(productosCantidad, ProductoCantidad)
	}
	pedido.ListaProductos = productosCantidad
	_, err := ps.pedidosRepository.CrearPedido(pedido.GetModel())
	return err
}

func (ps *PedidosService) ObtenerPedidos() ([]*dto.Pedidos, error) {
	pedidosDB, err := ps.pedidosRepository.ObtenerPedidos()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos, err
}

func (ps *PedidosService) EliminarPedido(id string) error {
	_, err := ps.pedidosRepository.EliminarPedido(id)
	// if resultado.ModifiedCount != 0 {
	// 	return true
	// }
	// return false

	return err
}

func (ps *PedidosService) ObtenerPedidosPendientes() ([]*dto.Pedidos, error) {
	pedidosDB, err := ps.pedidosRepository.ObtenerPedidosPendientes()

	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidosDB {
		pedido := dto.NewPedidos(*pedidoDB)
		pedidos = append(pedidos, pedido)
	}

	return pedidos, err
}

func (ps *PedidosService) ActualizarPedidoAceptado(id string) error {
	_, err := ps.pedidosRepository.ActualizarPedidoAceptado(id)
	// if resultado.ModifiedCount != 0 {
	// 	return true
	// }
	// return false
	return err
}
