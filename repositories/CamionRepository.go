package repositories

type CamionRepositoryInterface interface {
	//metodos
}

type CamionRepository struct {
	db DB
}

func NewCamionRepository(db DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

//CRUD de Camion
