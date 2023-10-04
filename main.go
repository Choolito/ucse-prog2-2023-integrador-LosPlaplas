package main

import (
	"log"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/handlers"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/services"
	"github.com/gin-gonic/gin"
)

var (
	camionHandler   *handlers.CamionHandler
	pedidosHandler  *handlers.PedidosHandler
	productoHandler *handlers.ProductoHandler
	enviosHandler   *handlers.EnviosHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()

	dependencies()

	mappingRoutes()

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}

func mappingRoutes() {

}

func dependencies() {
	//DB
	var database repositories.DB
	database = repositories.NewMongoDB()

	//Camion
	var camionRepository repositories.CamionRepositoryInterface
	var camionService services.CamionInterface
	camionRepository = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepository)
	camionHandler = handlers.NewCamionHandler(camionService)

	//Producto
	var productoRepository repositories.ProductoRepositoryInterface
	var productoService services.ProductoInterface
	productoRepository = repositories.NewProductoRepository(database)
	productoService = services.NewProductoService(productoRepository)
	productoHandler = handlers.NewProductoHandler(productoService)

	//Pedidos
	var pedidosRepository repositories.PedidosRepositoryInterface
	var pedidosService services.PedidosInterface
	pedidosRepository = repositories.NewPedidosRepository(database)
	pedidosService = services.NewPedidosService(pedidosRepository)
	pedidosHandler = handlers.NewPedidosHandler(pedidosService)

	//Envios
	var enviosRepository repositories.EnviosRepositoryInterface
	var enviosService services.EnviosInterface
	enviosRepository = repositories.NewEnviosRepository(database)
	enviosService = services.NewEnviosService(enviosRepository)
	enviosHandler = handlers.NewEnviosHandler(enviosService, pedidosService, camionService)
}
