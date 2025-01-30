package repositories

import (
	"context"
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	//metodos
	CrearCamion(camion model.Camion) error
	ObtenerCamiones() ([]*model.Camion, error)
	ActualizarCamion(id string, camion model.Camion) error
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

func (cr *CamionRepository) ActualizarCamion(id string, camion model.Camion) error {
	collection := cr.db.GetClient().Database("LosPlaplas").Collection("camiones")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"costoPorKilometro":  camion.CostoPorKilometro,
			"fechaActualizacion": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.Background(), filtro, update)
	return err
}

func (cr *CamionRepository) EliminarCamion(id string) (*mongo.DeleteResult, error) {
	collection := cr.db.GetClient().Database("LosPlaplas").Collection("camiones")

	objectID := utils.GetObjectIDFromStringID(id)

	filtro := bson.M{"_id": objectID}

	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}
