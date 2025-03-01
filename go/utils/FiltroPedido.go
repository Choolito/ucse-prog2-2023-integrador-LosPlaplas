package utils

import (
	"time"
)

type FiltroPedido struct {
	IdEnvio               string
	Estado                string
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
