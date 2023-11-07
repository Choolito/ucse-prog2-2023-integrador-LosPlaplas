package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type EnviosInterface interface {
	//metodos
	CrearEnvio(envio *dto.Envio) bool
	IniciarViajeEnvio(id string) bool
	GenerarParadaEnvio(id string, parada dto.Parada) bool
	FinalizarViajeEnvio(id string, paradaDestino dto.Parada) bool
	ObtenerEnvio() []*dto.Envio
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

func (enviosService *EnviosService) CrearEnvio(envio *dto.Envio) bool {
	camion, _ := enviosService.camionRepository.ObtenerCamionPorID(envio.IDCamion)
	pesoMaximo := camion.PesoMaximo

	pedidos := envio.Pedidos
	var pedidosEnvio []*model.Pedidos
	for _, pedido := range pedidos {
		pedidoBuscado, _ := enviosService.pedidosRepository.ObtenerPedidoPorID(pedido)
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
		enviosService.pedidosRepository.ActualizarPedidoParaEnviar(pedido.ID.Hex())
	}

	enviosService.enviosRepository.CrearEnvio(envio.GetModel())

	return true
}

func (enviosService *EnviosService) ObtenerEnvio() []*dto.Envio {
	enviosDB, _ := enviosService.enviosRepository.ObtenerEnvio()

	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(*envioDB)
		envios = append(envios, envio)
	}

	return envios
}

func (enviosService *EnviosService) IniciarViajeEnvio(id string) bool {
	enviosService.enviosRepository.IniciarViajeEnvio(id)
	return true
}

func (enviosService *EnviosService) GenerarParadaEnvio(id string, parada dto.Parada) bool {
	envio, _ := enviosService.enviosRepository.ObtenerEnvioPorID(id)
	if envio.Estado == "En ruta" {
		enviosService.enviosRepository.GenerarParadaEnvio(id, parada.GetModel())
		return true
	}

	return false
}

func (enviosService *EnviosService) FinalizarViajeEnvio(id string, paradaDestino dto.Parada) bool {
	envio, _ := enviosService.enviosRepository.ObtenerEnvioPorID(id)
	if envio.Estado == "En ruta" {
		enviosService.enviosRepository.GenerarParadaEnvio(id, paradaDestino.GetModel())
		enviosService.enviosRepository.FinalizarViajeEnvio(id)
		for _, pedido := range envio.Pedidos {
			enviosService.pedidosRepository.ActualizarPedidoEnviado(pedido)
			pedido, _ := enviosService.pedidosRepository.ObtenerPedidoPorID(pedido)
			for _, producto := range pedido.ListaProductos {
				var idProducto = utils.GetStringIDFromObjectID(producto.IDProducto)
				enviosService.productoRepository.DescontarStock(idProducto, producto.Cantidad)
			}

		}
		return true
	}

	return false
}
