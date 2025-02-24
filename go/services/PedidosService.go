package services

import (
	"fmt"
	"log"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/repositories"
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
		productoBuscado, err := ps.productoRepository.ObtenerProductoPorID(producto.IDProducto)
		if err != nil || productoBuscado == nil {
			log.Printf("error: Producto con ID %s no encontrado o está eliminado", producto.IDProducto)
			return fmt.Errorf("el producto con ID %s no existe o está eliminado", producto.IDProducto)
		}

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
	pedido, err := ps.pedidosRepository.ObtenerPedidoPorID(id)
	if err != nil {
		return err
	}

	estadosNoCancelables := []string{"Aceptado", "Para enviar", "Enviado"}
	for _, estado := range estadosNoCancelables {
		if string(pedido.EstadoPedido) == estado {
			return fmt.Errorf("no se puede eliminar el pedido con el id: %s porque está en estado %s", id, estado)
		}
	}

	eliminado, err := ps.pedidosRepository.EliminarPedido(id)
	if err != nil {
		return err
	}
	if !eliminado {
		return fmt.Errorf("no se pudo eliminar el pedido con el id: %s", id)
	}
	return nil
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
	resultado, err := ps.pedidosRepository.ActualizarPedidoAceptado(id)
	if err != nil {
		return err
	}
	if resultado.ModifiedCount == 0 {
		return fmt.Errorf("no se pudo actualizar el pedido con el id: %s", id)
	}
	return nil
}
