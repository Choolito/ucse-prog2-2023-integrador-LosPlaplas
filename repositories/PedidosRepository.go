package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidosRepositoryInterface interface {
	//metodos
	CrearPedido(pedido model.Pedidos) (*mongo.InsertOneResult, error)
	ObtenerPedidos() ([]*model.Pedidos, error)
	ActualizarPedido(pedido model.Pedidos) error
	EliminarPedido(id string) (*mongo.UpdateResult, error)
	ObtenerPedidosPendientes() ([]*model.Pedidos, error)
	ActualizarPedidoAceptado(id string) (*mongo.UpdateResult, error)

	//envio
	ObtenerPedidoPorID(id string) (*model.Pedidos, error)
	ActualizarPedidoParaEnviar(id string) (*mongo.UpdateResult, error)
	ActualizarPedidoEnviado(id string) (*mongo.UpdateResult, error)
}

type PedidosRepository struct {
	db DB
}

func NewPedidosRepository(db DB) *PedidosRepository {
	return &PedidosRepository{
		db: db,
	}
}

//CRUD de Pedidos

func (pr *PedidosRepository) CrearPedido(pedido model.Pedidos) (*mongo.InsertOneResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	pedido.FechaCreacion = time.Now()
	pedido.FechaActualizacion = time.Now()
	pedido.EstadoPedido = "Pendiente"
	resultado, err := collection.InsertOne(context.TODO(), pedido)
	return resultado, err
}

func (pr *PedidosRepository) ObtenerPedidos() ([]*model.Pedidos, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var pedidos []*model.Pedidos
	for cursor.Next(context.Background()) {
		var pedido model.Pedidos
		err := cursor.Decode(&pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, &pedido)
	}
	return pedidos, err
}

func (pr *PedidosRepository) ObtenerPedidoPorID(id string) (*model.Pedidos, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	var pedido model.Pedidos
	err := collection.FindOne(context.Background(), filtro).Decode(&pedido)
	return &pedido, err
}

// Actualizar general
func (pr *PedidosRepository) ActualizarPedido(pedido *model.Pedidos) error {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")

	filtro := bson.M{"_id": pedido.ID}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       pedido.EstadoPedido,
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)

	if resultado.MatchedCount == 0 {
		return errors.New("no se encontró el pedido a actualizar")
	}

	return err
}

func (pr *PedidosRepository) EliminarPedido(id string) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")

	objectID := utils.GetObjectIDFromStringID(id)

	estadosNoCancelables := []string{"Aceptado", "Para enviar", "Enviado"}

	filtro := bson.M{
		"_id":          objectID,
		"estadoPedido": bson.M{"$nin": estadosNoCancelables}, // Validar que el estado no esté en la lista de estados no cancelables
	}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       "Cancelado",
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}

// Lista pedidos pendientes
func (pr *PedidosRepository) ObtenerPedidosPendientes() ([]*model.Pedidos, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	filtro := bson.M{"estadoPedido": "Pendiente"}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var pedidos []*model.Pedidos
	for cursor.Next(context.Background()) {
		var pedido model.Pedidos
		err := cursor.Decode(&pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, &pedido)
	}
	return pedidos, err
}

// Aceptar pedido
func (pr *PedidosRepository) ActualizarPedidoAceptado(id string) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{
		"_id":          objectID,
		"estadoPedido": "Pendiente",
	}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       "Aceptado",
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}

func (pr *PedidosRepository) ActualizarPedidoParaEnviar(id string) (*mongo.UpdateResult, error) {
	collecction := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       "Para enviar",
			"fechaActualizacion": time.Now(),
		},
	}

	resultado, err := collecction.UpdateOne(context.Background(), filtro, update)

	return resultado, err
}

func (pr *PedidosRepository) ActualizarPedidoEnviado(id string) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       "Enviado",
			"fechaActualizacion": time.Now(),
		},
	}

	resultado, err := collection.UpdateOne(context.Background(), filtro, update)

	return resultado, err
}
