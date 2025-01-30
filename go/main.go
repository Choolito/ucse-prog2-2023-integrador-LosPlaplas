package main

import (
	"log"

	//"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/clients"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/handlers"
	//"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/middlewares"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/repositories"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/services"
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

	//authClient := clients.NewAuthClient()
	//authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Uso del middleware para todas las rutas del grupo
	//router.Use(middlewares.CORSMiddleware())
	//router.Use(authMiddleware.ValidateToken)

	//router.Use(middlewares.CORSMiddleware())
	//Productos CRUD
	router.POST("/productos", productoHandler.CrearProducto)
	router.GET("/productos", productoHandler.ObtenerProductos)
	router.GET("/productos/:id", productoHandler.ObtenerProductoPorID)
	router.GET("/productos/stockminimo", productoHandler.ObtenerListaConStockMinimo) //lista stock minimo y ?categoria= filtro
	//Falta usar este
	router.PUT("/productos/:id", productoHandler.ActualizarProducto)
	router.DELETE("/productos/:id", productoHandler.EliminarProducto)

	//Camiones CRUD
	router.POST("/camiones", camionHandler.CrearCamion)
	router.GET("/camiones", camionHandler.ObtenerCamiones)
	router.GET("/camiones/:id", camionHandler.ObtenerCamionPorID)
	router.PUT("/camiones/:id", camionHandler.ActualizarCamion)
	router.DELETE("/camiones/:id", camionHandler.EliminarCamion)

	//Pedidos CRUD
	router.POST("/pedidos", pedidosHandler.CrearPedido)
	router.GET("/pedidos", pedidosHandler.ObtenerPedidos)
	//Se puede filtrar por código de envío, estado, rango de fecha de creación.
	//router.PUT("/pedidos/:id", pedidosHandler.ActualizarPedido)
	router.PUT("/pedidos/cancelar/:id", pedidosHandler.EliminarPedido)
	//Lista pedidos pendientes
	router.GET("/pedidos/pendientes", pedidosHandler.ObtenerPedidosPendientes)
	router.PUT("/pedidos/aceptar/:id", pedidosHandler.ActualizarPedidoAceptado)

	//Envios
	router.POST("/envios", enviosHandler.CrearEnvio)
	router.GET("/envios", enviosHandler.ObtenerEnvio)
	router.PUT("/envios/iniciar/:id", enviosHandler.IniciarViajeEnvio)
	router.PUT("/envios/parada/:id", enviosHandler.GenerarParadaEnvio)
	router.PUT("/envios/finalizar/:id", enviosHandler.FinalizarViajeEnvio)
}

func dependencies() {
	//DB
	//var database repositories.DB
	database := repositories.NewMongoDB()

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
