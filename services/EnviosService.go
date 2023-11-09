package services

import (
	"errors"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
)

type EnviosInterface interface {
	//metodos
	CrearEnvio(envio *dto.Envio) error
	IniciarViajeEnvio(id string) error
	GenerarParadaEnvio(id string, parada dto.Parada) error
	FinalizarViajeEnvio(id string, paradaDestino dto.Parada) error
	ObtenerEnvio() ([]*dto.Envio, error)
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

func (enviosService *EnviosService) CrearEnvio(envio *dto.Envio) error {
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
		return errors.New("El peso total de los pedidos supera el peso maximo del camion")
	}

	//Pasar a estado "Para enviar"
	for _, pedido := range pedidosEnvio {
		enviosService.pedidosRepository.ActualizarPedidoParaEnviar(pedido.ID.Hex())
	}

	_, err := enviosService.enviosRepository.CrearEnvio(envio.GetModel())

	return err
}

func (enviosService *EnviosService) ObtenerEnvio() ([]*dto.Envio, error) {
	enviosDB, err := enviosService.enviosRepository.ObtenerEnvio()

	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(*envioDB)
		envios = append(envios, envio)
	}

	return envios, err
}

func (enviosService *EnviosService) IniciarViajeEnvio(id string) error {
	_, err := enviosService.enviosRepository.IniciarViajeEnvio(id)
	return err
}

func (enviosService *EnviosService) GenerarParadaEnvio(id string, parada dto.Parada) error {
	_, err := enviosService.enviosRepository.GenerarParadaEnvio(id, parada.GetModel())

	return err
}

func (enviosService *EnviosService) FinalizarViajeEnvio(id string, paradaDestino dto.Parada) error {
	_, err := enviosService.enviosRepository.FinalizarViajeEnvio(id)
	if err != nil {
		return err
	}
	envio, _ := enviosService.enviosRepository.ObtenerEnvioPorID(id)
	enviosService.enviosRepository.GenerarParadaEnvio(id, paradaDestino.GetModel())
	for _, pedido := range envio.Pedidos {
		enviosService.pedidosRepository.ActualizarPedidoEnviado(pedido)
		pedido, _ := enviosService.pedidosRepository.ObtenerPedidoPorID(pedido)
		for _, producto := range pedido.ListaProductos {
			var idProducto = utils.GetStringIDFromObjectID(producto.IDProducto)
			enviosService.productoRepository.DescontarStock(idProducto, producto.Cantidad)
		}
	}
	return nil
}
