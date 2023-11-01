package repositories

import (
	"context"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	//metodos
	CreateProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	GetProductos() ([]*model.Producto, error)
	UpdateProducto(id string, producto model.Producto) (*mongo.UpdateResult, error)
	DeleteProducto(id string) (*mongo.DeleteResult, error)

	DiscountStock(id string, cantidad int) (*mongo.UpdateResult, error)

	//GetListStockMinimum() ([]*model.Producto, error)
	GetListFiltered(filtro string) ([]*model.Producto, error)

	GetProductoForID(id string) (*model.Producto, error)
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

func (pr *ProductoRepository) CreateProducto(producto model.Producto) (*mongo.InsertOneResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
	producto.FechaCreacion = time.Now()
	producto.FechaActualizacion = time.Now()
	resultado, err := collection.InsertOne(context.TODO(), producto)
	return resultado, err
}

func (pr *ProductoRepository) GetProductos() ([]*model.Producto, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

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

func (pr *ProductoRepository) GetProductoForID(id string) (*model.Producto, error) {
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

func (pr *ProductoRepository) UpdateProducto(id string, producto model.Producto) (*mongo.UpdateResult, error) {
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

func (pr *ProductoRepository) DeleteProducto(id string) (*mongo.DeleteResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}

func (pr *ProductoRepository) DiscountStock(id string, cantidad int) (*mongo.UpdateResult, error) {
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

// func (pr *ProductoRepository) GetListStockMinimum() ([]*model.Producto, error) {
// 	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")
// 	filtro := bson.M{"cantidadEnStock": bson.M{"$lte": "stockMinimo"}}

// 	cursor, err := collection.Find(context.TODO(), filtro)

// 	defer cursor.Close(context.Background())

// 	var productos []*model.Producto
// 	for cursor.Next(context.Background()) {
// 		var producto model.Producto
// 		err := cursor.Decode(&producto)
// 		if err != nil {
// 			return nil, err
// 		}
// 		productos = append(productos, &producto)
// 	}
// 	return productos, err
// }

func (pr *ProductoRepository) GetListFiltered(filtro string) ([]*model.Producto, error) {
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
