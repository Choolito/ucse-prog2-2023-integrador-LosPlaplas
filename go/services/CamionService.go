package services

import (
	"fmt"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/repositories"
)

type CamionInterface interface {
	//metodos
	CrearCamion(camion *dto.Camion) error
	ObtenerCamiones() ([]*dto.Camion, error)
	ActualizarCamion(id string, camion *dto.Camion) error
	EliminarCamion(id string) error
	ObtenerCamionPorID(id string) (*dto.Camion, error)
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{
		camionRepository: camionRepository,
	}
}

//CRUD de Camion

//PesoMaximo() --> devuelve el peso maximo del camion

func (cs *CamionService) CrearCamion(camion *dto.Camion) error {
	err := cs.camionRepository.CrearCamion(camion.GetModel())

	return err
}

func (cs *CamionService) ObtenerCamiones() ([]*dto.Camion, error) {
	camionesDB, err := cs.camionRepository.ObtenerCamiones()

	var camiones []*dto.Camion
	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(*camionDB)
		camiones = append(camiones, camion)
	}

	return camiones, err
}

func (cs *CamionService) ObtenerCamionPorID(id string) (*dto.Camion, error) {
	camionDB, err := cs.camionRepository.ObtenerCamionPorID(id)

	camion := dto.NewCamion(*camionDB)

	return camion, err
}

func (cs *CamionService) ActualizarCamion(id string, camion *dto.Camion) error {
	// Verificar si el camión existe
	_, err := cs.camionRepository.ObtenerCamionPorID(id)
	if err != nil {
		return fmt.Errorf("no se encontró el camión con el id: %s", id)
	}

	// Actualizar el camión
	result, err := cs.camionRepository.ActualizarCamion(id, camion.GetModel())
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no se encontró el camión con el id: %s", id)
	}
	return nil
}

func (cs *CamionService) EliminarCamion(id string) error {
	// Verificar si el camión existe
	_, err := cs.camionRepository.ObtenerCamionPorID(id)
	if err != nil {
		return fmt.Errorf("no se encontró el camión con el id: %s", id)
	}

	// Eliminar el camión
	eliminado, err := cs.camionRepository.EliminarCamion(id)
	if err != nil {
		return err
	}
	if eliminado.DeletedCount == 0 {
		return fmt.Errorf("no se pudo eliminar el camión con el id: %s", id)
	}
	return nil
}
