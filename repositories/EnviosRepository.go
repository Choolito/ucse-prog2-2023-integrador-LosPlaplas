package repositories

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type EnviosRepositoryInterface interface {
	//metodos
	CreateShipping(envio model.Envio) (*mongo.InsertOneResult, error)
	StartTrip(id string) (*mongo.UpdateResult, error)
	GenerateStop(id string, parada model.Parada) (*mongo.UpdateResult, error)
	GetShippingForID(id string) (model.Envio, error)
	FinishTrip(id string) (*mongo.UpdateResult, error)
}

type EnviosRepository struct {
	db DB
}

func NewEnviosRepository(db DB) *EnviosRepository {
	return &EnviosRepository{
		db: db,
	}
}

// metodos
// Generar envio
// Cambiar a createShipping
func (enviosRepository *EnviosRepository) CreateShipping(envio model.Envio) (*mongo.InsertOneResult, error) {
	collecction := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	envio.FechaCreacion = time.Now()
	envio.FechaActualizacion = time.Now()
	envio.Estado = "A despachar"
	resultado, err := collecction.InsertOne(context.TODO(), envio)

	return resultado, err
}

func (enviosRepository *EnviosRepository) StartTrip(id string) (*mongo.UpdateResult, error) {
	collecction := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"estado": "En ruta", "fechaActualizacion": time.Now()}}
	resultado, err := collecction.UpdateOne(context.Background(), filter, update)

	return resultado, err
}

func (enviosRepository *EnviosRepository) GenerateStop(id string, parada model.Parada) (*mongo.UpdateResult, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$push": bson.M{"paradas": parada, "fechaActualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}

func (enviosRepository *EnviosRepository) GetShippingForID(id string) (model.Envio, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}
	var envio model.Envio
	err := collection.FindOne(context.Background(), filter).Decode(&envio)
	return envio, err
}

func (EnviosRepository *EnviosRepository) FinishTrip(id string) (*mongo.UpdateResult, error) {
	collection := EnviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"estado": "Despachado", "fechaActualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}
