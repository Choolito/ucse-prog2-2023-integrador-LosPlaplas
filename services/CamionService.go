package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/dto"
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type CamionInterface interface {
	//metodos
	CrearCamion(camion *dto.Camion) bool
	ObtenerCamiones() []*dto.Camion
	ActualizarCamion(id string, camion *dto.Camion) bool
	EliminarCamion(id string) bool
	ObtenerCamionPorID(id string) *dto.Camion
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

func (cs *CamionService) CrearCamion(camion *dto.Camion) bool {
	cs.camionRepository.CrearCamion(camion.GetModel())

	return true
}

func (cs *CamionService) ObtenerCamiones() []*dto.Camion {
	camionesDB, _ := cs.camionRepository.ObtenerCamiones()

	var camiones []*dto.Camion
	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(*camionDB)
		camiones = append(camiones, camion)
	}

	return camiones
}

func (cs *CamionService) ObtenerCamionPorID(id string) *dto.Camion {
	camionDB, _ := cs.camionRepository.ObtenerCamionPorID(id)

	camion := dto.NewCamion(*camionDB)

	return camion
}

func (cs *CamionService) ActualizarCamion(id string, camion *dto.Camion) bool {
	cs.camionRepository.ActualizarCamion(id, camion.GetModel())

	return true
}

func (cs *CamionService) EliminarCamion(id string) bool {
	cs.camionRepository.EliminarCamion(id)

	return true
}
