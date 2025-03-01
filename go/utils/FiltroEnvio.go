package utils

import (
	"time"
)

type FiltroEnvio struct {
	PatenteCamion      string     `json:"patenteCamion"`
	Estado             string     `json:"estado"`
	UltimaParada       string     `json:"ultimaParada"`
	FechaCreacionDesde *time.Time `json:"fechaCreacionDesde"`
	FechaCreacionHasta *time.Time `json:"fechaCreacionHasta"`
}
