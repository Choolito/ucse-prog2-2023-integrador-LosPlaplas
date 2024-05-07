package repositories

import (
	"time"

	"errors"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type EnviosRepositoryInterface interface {
	//metodos
	CrearEnvio(envio model.Envio) error

	ObtenerEnvio() ([]*model.Envio, error)
	ObtenerEnvioPorID(id string) (model.Envio, error)

	ActualizarEnvio(envio *model.Envio) error

	IniciarViajeEnvio(id string) (*mongo.UpdateResult, error)
	GenerarParadaEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error)
	FinalizarViajeEnvio(id string) (*mongo.UpdateResult, error)
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
func (enviosRepository *EnviosRepository) CrearEnvio(envio model.Envio) error {
	collecction := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	envio.FechaCreacion = time.Now()
	envio.FechaActualizacion = time.Now()
	envio.Estado = "A despachar"
	_, err := collecction.InsertOne(context.TODO(), envio)

	return err
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

	envio, err := enviosRepository.ObtenerEnvioPorID(id)
	if err != nil {
		return nil, err // Manejo del error si no se puede obtener el envío.
	}

	if envio.Estado != "En ruta" {
		return nil, errors.New("el envío no está en estado 'En ruta', no se puede generar una parada")
	}

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

func (enviosRepository *EnviosRepository) FinalizarViajeEnvio(id string) (*mongo.UpdateResult, error) {
	// collection := EnviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	// objectID := utils.GetObjectIDFromStringID(id)
	// filter := bson.M{"_id": objectID, "estado": "En ruta"}
	// update := bson.M{"$set": bson.M{"estado": "Despachado", "fechaActualizacion": time.Now()}}
	// resultado, err := collection.UpdateOne(context.Background(), filter, update)
	// return resultado, err

	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID, "estado": "En ruta"}

	// Verificar si el envío está en estado "En ruta" antes de finalizarlo.
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err // Manejo del error si hay un problema en la consulta.
	}

	if count == 0 {
		return nil, errors.New("El envío no está en estado 'En ruta', no se puede finalizar") // Retorna un error si el estado no es el adecuado.
	}

	update := bson.M{"$set": bson.M{"estado": "Despachado", "fechaActualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	return resultado, err
}

func (EnviosRepository *EnviosRepository) ActualizarEnvio(envio model.Envio) error {
	collection := EnviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	filter := bson.M{"_id": envio.ID}

	update := bson.M{"$set": bson.M{
		"estado":             envio.Estado,
		"fechaActualizacion": time.Now(),
		"paradas":            envio.Paradas,
		"pedidos":            envio.Pedidos,
		"camionero":          envio.IDCamion,
	}}
	operacion, err := collection.UpdateOne(context.Background(), filter, update)

	if operacion.MatchedCount == 0 {
		return errors.New("No se encontró el envío")
	}

	return err
}
