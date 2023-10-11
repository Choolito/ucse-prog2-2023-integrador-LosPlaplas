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
}

type PedidosService struct {
	pedidosRepository repositories.PedidosRepositoryInterface
}

func NewPedidosService(pedidosRepository repositories.PedidosRepositoryInterface) *PedidosService {
	return &PedidosService{
		pedidosRepository: pedidosRepository,
	}
}

//CRUD de Pedidos

//Metodo que devuelva []AceptadosElegidos que no superen el pesoMaximo --> Pasa a estado "Para Enviar"

func (ps *PedidosService) CreatePedido(pedido *dto.Pedidos) bool {
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
	ps.pedidosRepository.DeletePedido(id)

	return true
}
