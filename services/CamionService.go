package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
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
	_, err := cs.camionRepository.CrearCamion(camion.GetModel())

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
	_, err := cs.camionRepository.ActualizarCamion(id, camion.GetModel())

	return err
}

func (cs *CamionService) EliminarCamion(id string) error {
	_, err := cs.camionRepository.EliminarCamion(id)

	return err
}
