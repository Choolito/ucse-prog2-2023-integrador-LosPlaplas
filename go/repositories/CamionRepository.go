package repositories

import (
	"context"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	//metodos
	CrearCamion(camion model.Camion) error
	ObtenerCamiones() ([]*model.Camion, error)
	ActualizarCamion(id string, camion model.Camion) (*mongo.UpdateResult, error)
	EliminarCamion(id string) (*mongo.DeleteResult, error)
	//envios
	ObtenerCamionPorID(id string) (*model.Camion, error)
}

type CamionRepository struct {
	db DB
}

func NewCamionRepository(db DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

//CRUD de Camion

func (cr *CamionRepository) CrearCamion(camion model.Camion) error {
	collection := cr.db.GetClient().Database("LosPlaplas").Collection("camiones")
	camion.FechaCreacion = time.Now()
	camion.FechaActualizacion = time.Now()
	_, err := collection.InsertOne(context.TODO(), camion)
	return err
}

func (cr *CamionRepository) ObtenerCamiones() ([]*model.Camion, error) {
	collection := cr.db.GetClient().Database("LosPlaplas").Collection("camiones")
	filtro := bson.M{}

	cursor, err := collection.Find(context.Background(), filtro)

	defer cursor.Close(context.Background())

	var camiones []*model.Camion
	for cursor.Next(context.Background()) {
		var camion model.Camion
		err := cursor.Decode(&camion)
		if err != nil {
			return nil, err
		}
		camiones = append(camiones, &camion)
	}
	return camiones, err
}

func (cr *CamionRepository) ObtenerCamionPorID(id string) (*model.Camion, error) {
	collection := cr.db.GetClient().Database("LosPlaplas").Collection("camiones")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	var camion model.Camion
	err := collection.FindOne(context.Background(), filtro).Decode(&camion)

	return &camion, err
}

func (repo *CamionRepository) ActualizarCamion(id string, camion model.Camion) (*mongo.UpdateResult, error) {
	collection := repo.db.GetClient().Database("LosPlaplas").Collection("camiones")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"costoPorKilometro":  camion.CostoPorKilometro,
			"fechaActualizacion": time.Now(),
		},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *CamionRepository) EliminarCamion(id string) (*mongo.DeleteResult, error) {
	collection := repo.db.GetClient().Database("LosPlaplas").Collection("camiones")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return nil, err
	}
	return result, nil
}
