package Usecase

import (
	"errors"

	"github.com/Afomiat/GoCrudChallange/Domain"
)

type PersonUsecase struct {
	PersonRepo Domain.PersonRepository
}

func NewPersonUsecase(personRepo Domain.PersonRepository) *PersonUsecase {
	return &PersonUsecase{
		PersonRepo: personRepo,
	}
}

func (pu *PersonUsecase) AddPerson(person Domain.Person) error {	
	if person.Name == "" || person.Age == 0 || person.Hobbies == nil {

		return errors.New("name, Age and Hobbies are required")
	}

	_, err := pu.PersonRepo.AddPerson(person)

	return err
}
func (pu *PersonUsecase) GetPersons() ([]Domain.Person, error) {
	return pu.PersonRepo.GetPersons()
}

func (pu *PersonUsecase) GetPersonById(id string) (Domain.Person, error){
	return pu.PersonRepo.GetPersonById(id)
}

func (sc *PersonUsecase) UpdatePerson(id string, person Domain.Person) error{
	if person.Name == "" || person.Age == 0 || person.Hobbies == nil {
		return errors.New("missing required fields")
	}
	_,err := sc.PersonRepo.UpdatePerson(id, person)
	return err
}

func (pu *PersonUsecase) DeletePerson(id string) (Domain.Person, error){
	return pu.PersonRepo.DeletePerson(id)
}