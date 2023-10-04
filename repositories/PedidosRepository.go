package repositories

type PedidosRepositoryInterface interface {
	//metodos
}

type PedidosRepository struct {
	db DB
}

func NewPedidosRepository(db DB) *PedidosRepository {
	return &PedidosRepository{
		db: db,
	}
}

//CRUD de Pedidos
