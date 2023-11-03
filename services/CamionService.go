package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type CamionInterface interface {
	//metodos
	CreateCamion(camion *dto.Camion) bool
	GetCamiones() []*dto.Camion
	UpdateCamion(id string, camion *dto.Camion) bool
	DeleteCamion(id string) bool
	GetCamionForID(id string) *dto.Camion
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

func (cs *CamionService) CreateCamion(camion *dto.Camion) bool {
	cs.camionRepository.CreateCamion(camion.GetModel())

	return true
}

func (cs *CamionService) GetCamiones() []*dto.Camion {
	camionesDB, _ := cs.camionRepository.GetCamiones()

	var camiones []*dto.Camion
	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(*camionDB)
		camiones = append(camiones, camion)
	}

	return camiones
}

func (cs *CamionService) GetCamionForID(id string) *dto.Camion {
	camionDB, _ := cs.camionRepository.GetCamionForID(id)

	camion := dto.NewCamion(*camionDB)

	return camion
}

func (cs *CamionService) UpdateCamion(id string, camion *dto.Camion) bool {
	cs.camionRepository.UpdateCamion(id, camion.GetModel())

	return true
}

func (cs *CamionService) DeleteCamion(id string) bool {
	cs.camionRepository.DeleteCamion(id)

	return true
}
