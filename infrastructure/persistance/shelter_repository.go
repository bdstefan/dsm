package persistance

import (
	"github.com/bdstefan/dsm/app/domain/model"
)

type ShelterRepository struct {

}

func (shelterRepository *ShelterRepository) FindAll() ([]*model.Shelter, error) {
	return nil, nil
}

func (shelterRepository *ShelterRepository) FindBy(value interface{}) (*model.Shelter, error) {
	return nil, nil
}

func (shelterRepository *ShelterRepository) Save(shelter *model.Shelter) error {
	return nil
}
