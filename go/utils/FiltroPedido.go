package utils

import (
	"time"

	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"
)

type FiltroPedido struct {
	IdPedidos             []string
	IdEnvio               string
	CodigoProducto        string
	Estado                model.EstadoPedido
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
