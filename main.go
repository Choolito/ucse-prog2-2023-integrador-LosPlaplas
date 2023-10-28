package main

import (
	"log"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/handlers"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/middlewares"
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

	//Productos CRUD
	groupProductos := router.Group("/productos")
	groupProductos.Use(middlewares.CORSMiddleware())

	groupProductos.POST("/", productoHandler.CreateProducto)
	groupProductos.GET("/", productoHandler.GetProductos)
	//Lista de productos con stock menor al mínimo. Se puede filtrar por tipo de producto.
	groupProductos.PUT("/:id", productoHandler.UpdateProducto)
	groupProductos.DELETE("/:id", productoHandler.DeleteProducto)

	//Camiones CRUD
	groupCamiones := router.Group("/camiones")
	groupCamiones.Use(middlewares.CORSMiddleware())

	groupCamiones.POST("/", camionHandler.CreateCamion)
	groupCamiones.GET("/", camionHandler.GetCamiones)
	groupCamiones.PUT("/:id", camionHandler.UpdateCamion)
	groupCamiones.DELETE("/:id", camionHandler.DeleteCamion)

	//Pedidos CRUD
	groupPedidos := router.Group("/pedidos")
	groupPedidos.Use(middlewares.CORSMiddleware())

	groupPedidos.POST("/", pedidosHandler.CreatePedido)
	groupPedidos.GET("/", pedidosHandler.GetPedidos)
	//Se puede filtrar por código de envío, estado, rango de fecha de creación.
	groupPedidos.PUT("/:id", pedidosHandler.UpdatePedido)
	groupPedidos.PUT("/cancelar/:id", pedidosHandler.DeletePedido)
	//Lista pedidos pendientes
	groupPedidos.GET("/pendientes", pedidosHandler.GetPedidosPendientes)
	groupPedidos.PUT("/aceptar/:id", pedidosHandler.UpdatePedidoAceptado)

	//Envios
	groupEnvios := router.Group("/envios")
	groupEnvios.Use(middlewares.CORSMiddleware())

	groupEnvios.POST("/", enviosHandler.CreateEnvio)
	groupEnvios.PUT("/iniciar/:id", enviosHandler.StartTrip)
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
	enviosService = services.NewEnviosService(enviosRepository, pedidosRepository, camionRepository)
	enviosHandler = handlers.NewEnviosHandler(enviosService)
}
