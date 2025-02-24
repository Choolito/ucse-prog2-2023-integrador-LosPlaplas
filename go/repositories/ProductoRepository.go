package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	//metodos
	CrearProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	CrearProductos(productos []model.Producto) (*mongo.InsertManyResult, error)
	ObtenerProductos() ([]*model.Producto, error)
	ObtenerProductoPorID(id string) (*model.Producto, error)
	ActualizarProducto(id string, producto model.Producto) (*mongo.UpdateResult, error)
	EliminarProducto(id string) error
	DescontarStock(id string, cantidad int) (*mongo.UpdateResult, error)
	ObtenerListaConStockMinimo() ([]*model.Producto, error)
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

// Método para insertar múltiples productos
func (pr *ProductoRepository) CrearProductos(productos []model.Producto) (*mongo.InsertManyResult, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	var documentos []interface{}
	for i := range productos {
		productos[i].FechaCreacion = time.Now()
		productos[i].FechaActualizacion = time.Now()
		documentos = append(documentos, productos[i])
	}

	resultado, err := collection.InsertMany(context.TODO(), documentos)
	return resultado, err
}

func (pr *ProductoRepository) ObtenerProductos() ([]*model.Producto, error) {
	collection := pr.db.GetClient().Database("LosPlaplas").Collection("productos")

	// Agregar filtro para traer solo productos con eliminado: false
	filtroDB := bson.M{"eliminado": false}

	cursor, err := collection.Find(context.TODO(), filtroDB)
	if err != nil {
		return nil, err
	}
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

	// Devolver la lista de productos
	return productos, nil
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
			"codigoProducto":     producto.CodigoProducto,
			"nombre":             producto.Nombre,
			"precioUnitario":     producto.PrecioUnitario,
			"pesoUnitario":       producto.PesoUnitario,
			"stockMinimo":        producto.StockMinimo,
			"cantidadEnStock":    producto.CantidadEnStock,
			"fechaActualizacion": time.Now(),
		},
	}
	resultado, err := collection.UpdateOne(context.Background(), filtro, update)
	return resultado, err
}

func (repo *ProductoRepository) EliminarProducto(id string) error {
	collection := repo.db.GetClient().Database("LosPlaplas").Collection("productos")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filtro := bson.M{"_id": objectID}
	actualizacion := bson.M{"$set": bson.M{"eliminado": true}}

	resultado, err := collection.UpdateOne(context.Background(), filtro, actualizacion)
	if err != nil {
		return err
	}
	if resultado.ModifiedCount == 0 {
		return fmt.Errorf("no se pudo marcar como eliminado el producto con el id: %s", id)
	}

	return nil
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

func (repo *ProductoRepository) ObtenerListaConStockMinimo() ([]*model.Producto, error) {
	var productos []*model.Producto
	collection := repo.db.GetClient().Database("LosPlaplas").Collection("productos")
	filter := bson.M{
		"$expr": bson.M{
			"$lt": []interface{}{"$cantidadEnStock", "$stockMinimo"},
		},
	} // Productos con cantidad en stock menor que el stock mínimo
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var producto model.Producto
		if err := cursor.Decode(&producto); err != nil {
			return nil, err
		}
		productos = append(productos, &producto)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return productos, nil
}
