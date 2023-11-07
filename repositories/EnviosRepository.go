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
	CrearEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	IniciarViajeEnvio(id string) (*mongo.UpdateResult, error)
	GenerarParadaEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error)
	ObtenerEnvioPorID(id string) (model.Envio, error)
	FinalizarViajeEnvio(id string) (*mongo.UpdateResult, error)
	ObtenerEnvio() ([]*model.Envio, error)
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
func (enviosRepository *EnviosRepository) CrearEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	collecction := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	envio.FechaCreacion = time.Now()
	envio.FechaActualizacion = time.Now()
	envio.Estado = "A despachar"
	resultado, err := collecction.InsertOne(context.TODO(), envio)

	return resultado, err
}

func (enviosRepository *EnviosRepository) ObtenerEnvio() ([]*model.Envio, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	filter := bson.M{}

	cursor, err := collection.Find(context.TODO(), filter)

	defer cursor.Close(context.Background())

	var envios []*model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			return nil, err
		}
		envios = append(envios, &envio)
	}
	return envios, err
}

func (enviosRepository *EnviosRepository) IniciarViajeEnvio(id string) (*mongo.UpdateResult, error) {
	collecction := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID, "estado": "A despachar"}
	update := bson.M{"$set": bson.M{"estado": "En ruta", "fechaActualizacion": time.Now()}}
	resultado, err := collecction.UpdateOne(context.Background(), filter, update)

	return resultado, err
}

func (enviosRepository *EnviosRepository) GenerarParadaEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}

	envio, _ := enviosRepository.ObtenerEnvioPorID(id)

	paradas := envio.Paradas
	paradas = append(paradas, parada)

	update := bson.M{
		"$set": bson.M{"paradas": paradas, "fechaActualizacion": time.Now()},
	}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}

func (enviosRepository *EnviosRepository) ObtenerEnvioPorID(id string) (model.Envio, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID}
	var envio model.Envio
	err := collection.FindOne(context.Background(), filter).Decode(&envio)
	return envio, err
}

func (EnviosRepository *EnviosRepository) FinalizarViajeEnvio(id string) (*mongo.UpdateResult, error) {
	collection := EnviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID, "estado": "En ruta"}
	update := bson.M{"$set": bson.M{"estado": "Despachado", "fechaActualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}
