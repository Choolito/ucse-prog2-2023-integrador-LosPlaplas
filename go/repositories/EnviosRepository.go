package repositories

import (
	"time"

	"errors"
	"log"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type EnviosRepositoryInterface interface {
	//metodos
	CrearEnvio(envio model.Envio) error

	ObtenerEnvio() ([]*model.Envio, error)
	ObtenerEnvioPorID(id string) (model.Envio, error)
	ObtenerEnviosFiltrados(filtro *utils.FiltroEnvio) ([]*model.Envio, error)

	ActualizarEnvio(envio *model.Envio) error

	IniciarViajeEnvio(id string) (*mongo.UpdateResult, error)
	GenerarParadaEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error)
	FinalizarViajeEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error)

	ObtenerPedidosFiltro(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedidos, error)
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

func (enviosRepository *EnviosRepository) FinalizarViajeEnvio(id string, parada model.Parada) (*mongo.UpdateResult, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"_id": objectID, "estado": "En ruta"}

	// Verificar si el envío está en estado "En ruta" antes de finalizarlo.
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("el envío no está en estado 'En ruta', no se puede finalizar")
	}

	// 🔹 Primero, guardar la última parada en la base de datos
	_, err = enviosRepository.GenerarParadaEnvio(id, parada)
	if err != nil {
		log.Printf("Error al generar parada para el envío %s: %v", id, err)
		return nil, err
	}

	// 🔹 Luego, actualizar el estado del envío a "Despachado"
	update := bson.M{"$set": bson.M{"estado": "Despachado", "fechaActualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return resultado, nil
}

func (EnviosRepository *EnviosRepository) ActualizarEnvio(envio *model.Envio) error {
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

func (enviosRepository *EnviosRepository) ObtenerPedidosFiltro(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedidos, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")
	filtro := bson.M{}
	if codigoEnvio != "" {
		filtro["codigoEnvio"] = codigoEnvio
	}
	if estado != "" {
		filtro["estadoPedido"] = estado
	}
	if !fechaInicio.IsZero() {
		filtro["fechaCreacion"] = bson.M{"$gte": fechaInicio}
	}
	if !fechaFinal.IsZero() {
		filtro["fechaCreacion"] = bson.M{"$lte": fechaFinal}
	}
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var pedidos []model.Pedidos
	for cursor.Next(context.Background()) {
		var pedido model.Pedidos
		if err := cursor.Decode(&pedido); err != nil {
			log.Printf("Error al decodificar pedido: %v", err)
			continue
		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}

func (enviosRepository *EnviosRepository) ObtenerEnviosFiltrados(filtro *utils.FiltroEnvio) ([]*model.Envio, error) {
	collection := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("envios")

	// Construir filtro dinámico
	filterMongo := bson.M{}

	if filtro.PatenteCamion != "" {
		var camion struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		collectionCamiones := enviosRepository.db.GetClient().Database("LosPlaplas").Collection("camiones")
		err := collectionCamiones.FindOne(context.TODO(), bson.M{"patente": filtro.PatenteCamion}).Decode(&camion)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// No hay camión con esa patente, retornar lista vacía sin error
				return []*model.Envio{}, nil
			}
			return nil, err // Error en la consulta
		}
		filterMongo["idCamion"] = camion.ID
	}
	if filtro.Estado != "" {
		filterMongo["estado"] = filtro.Estado
	}
	if filtro.UltimaParada != "" {
		filterMongo["paradas"] = bson.M{
			"$elemMatch": bson.M{
				"ciudad": filtro.UltimaParada,
			},
		}
	}

	// Filtrar por rango de fechas (desde, hasta o ambas)
	if filtro.FechaCreacionDesde != nil || filtro.FechaCreacionHasta != nil {
		dateFilter := bson.M{}
		if filtro.FechaCreacionDesde != nil {
			dateFilter["$gte"] = *filtro.FechaCreacionDesde
		}
		if filtro.FechaCreacionHasta != nil {
			dateFilter["$lte"] = *filtro.FechaCreacionHasta
		}
		filterMongo["fechaCreacion"] = dateFilter
	}

	// Consultar MongoDB con el filtro construido
	cursor, err := collection.Find(context.TODO(), filterMongo)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var envios []*model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		if err := cursor.Decode(&envio); err != nil {
			return nil, err
		}

		// Verificar si la última parada coincide
		if filtro.UltimaParada != "" && len(envio.Paradas) > 0 {
			// Obtiene la última parada de la lista
			ultimaParada := envio.Paradas[len(envio.Paradas)-1]

			// Verificar si la última parada es la que se busca
			if ultimaParada.Ciudad != filtro.UltimaParada {
				continue // Si la última parada no coincide, saltar este envío
			}
		}

		envios = append(envios, &envio)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return envios, nil
}
