package domain

type ShelterRepositoryInterface interface {
	FindAll() ([]*Shelter, error)
	FindBy(value interface{}) (*Shelter, error)
	Save(shelter *Shelter) error
}

type ShelterRepository struct {

}