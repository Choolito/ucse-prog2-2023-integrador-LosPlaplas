package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type EnviosInterface interface {
	//metodos
	CreateShipping(envio *dto.Envio) bool
	StartTrip(id string) bool
	GenerateStop(id string, parada dto.Parada) bool
	FinishTrip(id string, paradaDestino dto.Parada) bool
	GetShipping() []*dto.Envio
}

type EnviosService struct {
	enviosRepository   repositories.EnviosRepositoryInterface
	pedidosRepository  repositories.PedidosRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
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

func (enviosService *EnviosService) CreateShipping(envio *dto.Envio) bool {
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

	enviosService.enviosRepository.CreateShipping(envio.GetModel())

	return true
}

func (enviosService *EnviosService) GetShipping() []*dto.Envio {
	enviosDB, _ := enviosService.enviosRepository.GetShipping()

	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(*envioDB)
		envios = append(envios, envio)
	}

	return envios
}

func (enviosService *EnviosService) StartTrip(id string) bool {
	enviosService.enviosRepository.StartTrip(id)
	return true
}

func (enviosService *EnviosService) GenerateStop(id string, parada dto.Parada) bool {
	envio, _ := enviosService.enviosRepository.GetShippingForID(id)
	if envio.Estado == "En ruta" {
		enviosService.enviosRepository.GenerateStop(id, parada.GetModel())
		return true
	}

	return false
}

func (enviosService *EnviosService) FinishTrip(id string, paradaDestino dto.Parada) bool {
	envio, _ := enviosService.enviosRepository.GetShippingForID(id)
	if envio.Estado == "En ruta" {
		enviosService.enviosRepository.GenerateStop(id, paradaDestino.GetModel())
		enviosService.enviosRepository.FinishTrip(id)
		for _, pedido := range envio.Pedidos {
			enviosService.pedidosRepository.UpdatePedidoEnviado(pedido)
			pedido, _ := enviosService.pedidosRepository.GetPedidoForID(pedido)
			for _, producto := range pedido.ListaProductos {
				var idProducto = utils.GetStringIDFromObjectID(producto.IDProducto)
				enviosService.productoRepository.DiscountStock(idProducto, producto.Cantidad)
			}

		}
		return true
	}

	return false
}
