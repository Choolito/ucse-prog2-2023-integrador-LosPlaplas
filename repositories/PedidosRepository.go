package repositories

import (
	"context"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidosRepositoryInterface interface {
	//metodos
	CreatePedido(pedido model.Pedidos) (*mongo.InsertOneResult, error)
	GetPedidos() ([]*model.Pedidos, error)
	UpdatePedido(id string, pedido model.Pedidos) (*mongo.UpdateResult, error)
	DeletePedido(id string) (*mongo.UpdateResult, error)
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

func (pr *PedidosRepository) CreatePedido(pedido model.Pedidos) (*mongo.InsertOneResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")
	pedido.FechaCreacion = time.Now()
	pedido.FechaActualizacion = time.Now()
	pedido.EstadoPedido = "Pendiente"
	resultado, err := collection.InsertOne(context.TODO(), pedido)
	return resultado, err
}

func (pr *PedidosRepository) GetPedidos() ([]*model.Pedidos, error) {
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

func (pr *PedidosRepository) UpdatePedido(id string, pedido model.Pedidos) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"listaProductos":      pedido.ListaProductos,
			"ciudadDestinoPedido": pedido.CiudadDestinoPedido,
			"fechaActualizacion":  time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}

func (pr *PedidosRepository) DeletePedido(id string) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("pedidos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"estadoPedido":       "Cancelado",
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}
