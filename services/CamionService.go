package services

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"
)

type CamionInterface interface {
	//metodos
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
