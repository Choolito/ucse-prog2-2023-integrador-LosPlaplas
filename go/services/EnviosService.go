package services

import (
	"errors"
	"fmt"

	"log"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
)

type EnviosInterface interface {
	//metodos
	CrearEnvio(envio *dto.Envio) error
	IniciarViajeEnvio(id string) error
	GenerarParadaEnvio(id string, parada dto.Parada) error
	FinalizarViajeEnvio(id string, paradaDestino dto.Parada) error
	ObtenerEnvio() ([]*dto.Envio, error)
	CambiarEstadoEnvio(envio *dto.Envio, user *dto.User) (bool, error)
	ObtenerPedidosFiltrados(filtro *utils.FiltroPedido) ([]*dto.Pedidos, error)
}

type EnviosService struct {
	enviosRepository   repositories.EnviosRepositoryInterface
	pedidosRepository  repositories.PedidosRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewEnviosService(enviosRepository repositories.EnviosRepositoryInterface, pedidosRepository repositories.PedidosRepositoryInterface,
	camionRepository repositories.CamionRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *EnviosService {
	return &EnviosService{
		enviosRepository:   enviosRepository,
		pedidosRepository:  pedidosRepository,
		camionRepository:   camionRepository,
		productoRepository: productoRepository,
		//llamar desde aca todos los repositorios que necesite
	}
}

//metodos

// Se genere un envio
// Envio pasa a estado "A despachar".
func (es *EnviosService) CrearEnvio(envio *dto.Envio) error {

	// Verificar que el camión no esté vacío
	if envio.IDCamion == "" {
		return fmt.Errorf("el camión no puede estar vacío")
	}

	// Verificar que el camión exista
	camion, err := es.camionRepository.ObtenerCamionPorID(envio.IDCamion)
	if err != nil {
		return fmt.Errorf("no se encontró el camión con el id: %s", envio.IDCamion)
	}

	// Verificar que los pedidos no estén vacíos
	if len(envio.Pedidos) == 0 {
		return fmt.Errorf("los pedidos no pueden estar vacíos")
	}

	// Verificar que el peso total de los pedidos no supere el peso máximo del camión
	pesoMaximo := camion.PesoMaximo
	pedidos := envio.Pedidos
	var pedidosEnvio []*model.Pedidos
	for _, pedidoID := range pedidos {
		pedido, err := es.pedidosRepository.ObtenerPedidoPorID(pedidoID)
		if err != nil {
			return fmt.Errorf("no se encontró el pedido con el id: %s", pedidoID)
		}
		if pedido.EstadoPedido != "Aceptado" {
			return fmt.Errorf("el pedido con el id %s no está en estado 'Aceptado' y no puede ser incluido en el envío", pedidoID)
		}
		pedidosEnvio = append(pedidosEnvio, pedido)
	}

	pesoTotalPedidos := 0
	for _, pedido := range pedidosEnvio {
		for _, producto := range pedido.ListaProductos {
			pesoTotalPedidos += producto.PesoUnitario * producto.Cantidad
		}
	}

	if pesoTotalPedidos > pesoMaximo {
		return fmt.Errorf("el peso total de los pedidos supera el peso máximo del camión")
	}

	// Crear el envío
	envioModel := envio.GetModel()
	err = es.enviosRepository.CrearEnvio(envioModel)
	if err != nil {
		return err
	}
	for _, pedido := range pedidosEnvio {
		_, err := es.pedidosRepository.ActualizarPedidoParaEnviar(pedido.ID.Hex())
		if err != nil {
			return err
		}
	}
	return nil
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

func (enviosService *EnviosService) CambiarEstadoEnvio(envio *dto.Envio, user *dto.User) (bool, error) {
	nuevoEstado := envio.Estado

	if !model.EsUnEstadoEnvioValido(nuevoEstado) {
		return false, errors.New("El estado del envío no es válido")
	}

	envioDB, err := enviosService.enviosRepository.ObtenerEnvioPorID(envio.ID)

	if err != nil {
		return false, err
	}

	//rol validar

	if (nuevoEstado == model.EnRuta && envioDB.Estado != model.ADespachar) || (nuevoEstado == model.Despachado && envioDB.Estado != model.EnRuta) {
		return false, errors.New("El envio no puede pasar del estado " + fmt.Sprint(nuevoEstado) + " si se encuentra en el estado " + fmt.Sprint(envioDB.Estado))
	}

	envioDB.Estado = nuevoEstado
	err = enviosService.enviosRepository.ActualizarEnvio(&envioDB)

	if err != nil {
		return false, err
	}

	if nuevoEstado == model.Despachado {
		enviosService.finalizarViaje(dto.NewEnvio(envioDB))
	}

	return true, nil

}

func (enviosService *EnviosService) finalizarViaje(envio *dto.Envio) (bool, error) {
	//pasar pedidos a estado enviado
	err := enviosService.entregarPedidosDeEnvio(envio)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (enviosService *EnviosService) entregarPedidosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {

		//Descuenta el stock de los productos
		err := enviosService.entregarPedido(&dto.Pedidos{ID: idPedido})

		if err != nil {
			return err
		}
	}
	return nil
}

func (enviosService *EnviosService) entregarPedido(pedidoPorEntregar *dto.Pedidos) error {
	//Primero buscamos el pedido a entregar
	pedido, err := enviosService.pedidosRepository.ObtenerPedidoPorID(pedidoPorEntregar.ID)

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Para enviar
	if pedido.EstadoPedido != model.ParaEnviar {
		return nil
	}

	//Cambia el estado del pedido a Enviado, si es que no estaba ya en ese estado
	if pedido.EstadoPedido != model.Enviado {
		pedido.EstadoPedido = model.Enviado
	}

	//Actualiza el pedido en la base de datos
	return enviosService.pedidosRepository.ActualizarPedido(pedido)
}

func (envioService *EnviosService) ObtenerPedidosFiltrados(filtro *utils.FiltroPedido) ([]*dto.Pedidos, error) {
	if filtro == nil {
		err := errors.New("el filtro no puede ser nulo")
		log.Printf("[service:EnvioService][method:ObtenerPedidosFiltro][reason:INVALID_INPUT][error:%v]", err)
		return nil, err
	}
	pedidoDB, err := envioService.enviosRepository.ObtenerPedidosFiltro(filtro.IdEnvio, filtro.Estado, filtro.FechaCreacionComienzo, filtro.FechaCreacionFin)
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerPedidosFiltro][reason:ERROR][error:%v]", err)
		return nil, err
	}
	var pedidos []*dto.Pedidos
	for _, pedidoDB := range pedidoDB {
		pedido := dto.NewPedidos(pedidoDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}
