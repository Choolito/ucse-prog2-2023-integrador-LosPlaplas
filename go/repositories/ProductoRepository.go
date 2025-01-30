package repositories

import (
	"context"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	//metodos
	CrearProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	ObtenerProductos() ([]*model.Producto, error)
	ObtenerProductoPorID(id string) (*model.Producto, error)
	ActualizarProducto(id string, producto model.Producto) (*mongo.UpdateResult, error)
	EliminarProducto(id string) (*mongo.DeleteResult, error)

	DescontarStock(id string, cantidad int) (*mongo.UpdateResult, error)

	ObtenerListaFiltrada(filtro string) ([]*model.Producto, error)
}

type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{
		db: db,
	}
}

//CRUD

func (pr *ProductoRepository) CrearProducto(producto model.Producto) (*mongo.InsertOneResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
	producto.FechaCreacion = time.Now()
	producto.FechaActualizacion = time.Now()
	resultado, err := collection.InsertOne(context.TODO(), producto)
	return resultado, err
}

func (pr *ProductoRepository) ObtenerProductos() ([]*model.Producto, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
	filtroDB := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtroDB)

	defer cursor.Close(context.Background())

	var productos []*model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			return nil, err
		}
		productos = append(productos, &producto)
	}
	return productos, err
}

func (pr *ProductoRepository) ObtenerProductoPorID(id string) (*model.Producto, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	var producto model.Producto
	err := collection.FindOne(context.Background(), filtro).Decode(&producto)
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

func (pr *ProductoRepository) ActualizarProducto(id string, producto model.Producto) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"nombre":             producto.Nombre,
			"precioUnitario":     producto.PrecioUnitario,
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}

func (pr *ProductoRepository) EliminarProducto(id string) (*mongo.DeleteResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}

func (pr *ProductoRepository) DescontarStock(id string, cantidad int) (*mongo.UpdateResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$inc": bson.M{"cantidadEnStock": -cantidad, "fechaActualizacion": time.Now()},
	}

	result, err := collection.UpdateOne(context.Background(), filtro, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pr *ProductoRepository) ObtenerListaFiltrada(filtro string) ([]*model.Producto, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
	filtroDB := bson.M{"tipoProducto": filtro}

	cursor, err := collection.Find(context.Background(), filtroDB)

	defer cursor.Close(context.Background())

	var productos []*model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			return nil, err
		}
		productos = append(productos, &producto)
	}
	return productos, err
}
