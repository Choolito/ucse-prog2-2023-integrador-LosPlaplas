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
	CreateEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	StartTrip(id string) (*mongo.UpdateResult, error)
	GenerateStop(id string, parada model.Parada) (*mongo.UpdateResult, error)
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
func (enviosRepository *EnviosRepository) CreateEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
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
	update := bson.M{"$push": bson.M{"paradas": parada}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}
