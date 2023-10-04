package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type PedidosInterface interface {
	//metodos
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
