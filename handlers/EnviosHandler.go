package handlers

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
)

type EnviosHandler struct {
	enviosService  services.EnviosInterface
	pedidosService services.PedidosInterface
	camionService  services.CamionInterface
}

func NewEnviosHandler(enviosService services.EnviosInterface, pedidosService services.PedidosInterface, camionService services.CamionInterface) *EnviosHandler {
	return &EnviosHandler{
		enviosService:  enviosService,
		pedidosService: pedidosService,
		camionService:  camionService,
	}
}

//Generar envio, en base a pedidos pendientes y se valida manualmente el stock --> estado = aceptado
//Duda

//ejemplo
/*func crearEnvio...{

	PesoMaximo := camionService.PesoMaximo()
	pedidos := pedidosService.AceptadosElegidos(PesoMaximo)

	envio := enviosServices.CrearEnvio(pedidos)

}*/

//Post Camionero inicia viaje --> envio estado "En ruta"
//Las paradas las puede ir haciendo y se guarda ciudad y km recorridos desde ultima parada/inicio.
//Metodo generarParadas
