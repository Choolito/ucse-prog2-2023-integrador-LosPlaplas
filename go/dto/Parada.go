package dto

import "github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/go/model"

type Parada struct {
	Ciudad       string
	KmRecorridos int
}

func NewParada(parada model.Parada) *Parada {
	return &Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}

func (parada Parada) GetModel() model.Parada {
	return model.Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}
