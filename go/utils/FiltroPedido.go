package utils

import (
	"time"
)

type FiltroPedido struct {
	IdPedidos             []string
	IdEnvio               string
	CodigoProducto        string
	Estado                string
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
