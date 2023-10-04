package services

import "github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/repositories"

type EnviosInterface interface {
	//metodos
}

type EnviosService struct {
	enviosRepository repositories.EnviosRepositoryInterface
}

func NewEnviosService(enviosRepository repositories.EnviosRepositoryInterface) *EnviosService {
	return &EnviosService{
		enviosRepository: enviosRepository,
	}
}

//metodos

//Se genere un envio
//Envio pasa a estado "A despachar".
