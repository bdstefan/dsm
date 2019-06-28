package repository

type ShelterRepositoryInterface interface {
	FindAll() ([]*model.Shelter, error)
	FindBy(value interface{}) (*model.Shelter, error)
	Save(shelter *model.Shelter) error
}

type ShelterRepository struct {

}