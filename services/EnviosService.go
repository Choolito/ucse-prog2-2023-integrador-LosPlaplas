package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type EnviosInterface interface {
	//metodos
	CreateEnvio(envio *dto.Envio) bool
	StartTrip(id string) bool
	GenerateStop(id string, parada dto.Parada) bool
}

type EnviosService struct {
	enviosRepository  repositories.EnviosRepositoryInterface
	pedidosRepository repositories.PedidosRepositoryInterface
	camionRepository  repositories.CamionRepositoryInterface
}

func NewEnviosService(enviosRepository repositories.EnviosRepositoryInterface, pedidosRepository repositories.PedidosRepositoryInterface,
	camionRepository repositories.CamionRepositoryInterface) *EnviosService {
	return &EnviosService{
		enviosRepository:  enviosRepository,
		pedidosRepository: pedidosRepository,
		camionRepository:  camionRepository,
		//llamar desde aca todos los repositorios que necesite
	}
}

//metodos

//Se genere un envio
//Envio pasa a estado "A despachar".

func (enviosService *EnviosService) CreateEnvio(envio *dto.Envio) bool {
	camion, _ := enviosService.camionRepository.GetCamionForID(envio.IDCamion)
	pesoMaximo := camion.PesoMaximo

	pedidos := envio.Pedidos
	var pedidosEnvio []*model.Pedidos
	for _, pedido := range pedidos {
		pedidoBuscado, _ := enviosService.pedidosRepository.GetPedidoForID(pedido)
		pedidosEnvio = append(pedidosEnvio, pedidoBuscado)
	}

	pesoTotalPedidos := 0

	for _, pedido := range pedidosEnvio {
		for _, producto := range pedido.ListaProductos {
			pesoTotalPedidos += producto.PesoUnitario * producto.Cantidad
		}
	}

	if pesoTotalPedidos > pesoMaximo {
		return false
	}

	//Pasar a estado "Para enviar"
	for _, pedido := range pedidosEnvio {
		enviosService.pedidosRepository.UpdatePedidoParaEnviar(pedido.ID.Hex())
	}

	enviosService.enviosRepository.CreateEnvio(envio.GetModel())

	return true
}

func (enviosService *EnviosService) StartTrip(id string) bool {
	enviosService.enviosRepository.StartTrip(id)
	return true
}

func (enviosService *EnviosService) GenerateStop(id string, parada dto.Parada) bool {
	enviosService.enviosRepository.GenerateStop(id, parada.GetModel())
	//Falta corroborar que haya llegado a destino
	return true
}
