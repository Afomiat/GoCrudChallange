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