package utils

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
)

type FiltroEnvio struct {
	PatenteCamion                 string
	Estado                        model.EstadoEnvio
	UltimaParada                  string
	FechaCreacionDesde            time.Time
	FechaCreacionHasta            time.Time
	FechaUltimaActualizacionDesde time.Time
	FechaUltimaActualizacionHasta time.Time
}
