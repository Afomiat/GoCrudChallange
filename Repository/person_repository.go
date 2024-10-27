package Repository

import (
	"context"

	"github.com/Afomiat/GoCrudChallange/Domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepoImplement struct {
	collection *mongo.Collection
}

func NewPersonRepoImplement(coll *mongo.Collection) Domain.PersonRepository {	
	return &PersonRepoImplement{
		collection: coll,
	}
}

func (pr *PersonRepoImplement) AddPerson(person Domain.Person) (Domain.Person, error) {
	_, err := pr.collection.InsertOne(context.Background(), person)
	return person, err
}