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

	router.Use(middlewares.CORSMiddleware())
	//Productos CRUD
	router.POST("/productos", productoHandler.CreateProducto)
	router.GET("/productos", productoHandler.GetProductos)
	//lista stock minimo
	router.GET("/productos/stockminimo", productoHandler.GetListStockMinimum)
	router.GET("/productos/filtrado/:filtro", productoHandler.GetListFiltered)
	router.PUT("/productos/:id", productoHandler.UpdateProducto)
	router.DELETE("/productos/:id", productoHandler.DeleteProducto)

	//Camiones CRUD
	router.POST("/camiones", camionHandler.CreateCamion)
	router.GET("/camiones", camionHandler.GetCamiones)
	router.PUT("/camiones/:id", camionHandler.UpdateCamion)
	router.DELETE("/camiones/:id", camionHandler.DeleteCamion)

	//Pedidos CRUD
	router.POST("/pedidos", pedidosHandler.CreatePedido)
	router.GET("/pedidos", pedidosHandler.GetPedidos)
	//Se puede filtrar por código de envío, estado, rango de fecha de creación.
	router.PUT("/pedidos/:id", pedidosHandler.UpdatePedido)
	router.PUT("/pedidos/cancelar/:id", pedidosHandler.DeletePedido)
	//Lista pedidos pendientes
	router.GET("/pedidos/pendientes", pedidosHandler.GetPedidosPendientes)
	router.PUT("/pedidos/aceptar/:id", pedidosHandler.UpdatePedidoAceptado)

	//Envios
	router.POST("/envios", enviosHandler.CreateShipping)
	router.GET("/envios", enviosHandler.GetShipping)
	router.PUT("/envios/iniciar/:id", enviosHandler.StartTrip)
	router.PUT("/envios/parada/:id", enviosHandler.GenerateStop)
	router.PUT("/envios/finalizar/:id", enviosHandler.FinishTrip)
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
	pedidosService = services.NewPedidosService(pedidosRepository, productoRepository)
	pedidosHandler = handlers.NewPedidosHandler(pedidosService)

	//Envios
	var enviosRepository repositories.EnviosRepositoryInterface
	var enviosService services.EnviosInterface
	enviosRepository = repositories.NewEnviosRepository(database)
	enviosService = services.NewEnviosService(enviosRepository, pedidosRepository, camionRepository)
	enviosHandler = handlers.NewEnviosHandler(enviosService)
}
