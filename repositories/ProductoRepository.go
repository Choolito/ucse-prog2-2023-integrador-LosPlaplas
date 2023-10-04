package repositories

type ProductoRepositoryInterface interface {
	//metodos
}

type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{
		db: db,
	}
}
