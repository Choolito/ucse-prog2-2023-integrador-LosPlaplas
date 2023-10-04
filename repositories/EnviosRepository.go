package repositories

type EnviosRepositoryInterface interface {
	//metodos
}

type EnviosRepository struct {
	db DB
}

func NewEnviosRepository(db DB) *EnviosRepository {
	return &EnviosRepository{
		db: db,
	}
}

//metodos
//Generar envio
