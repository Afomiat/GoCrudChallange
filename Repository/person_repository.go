package Repository

import (
	"context"

	"github.com/Afomiat/GoCrudChallange/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (pr *PersonRepoImplement) GetPersons() ([]Domain.Person, error) {
	var cursor *mongo.Cursor
	var err error

	cursor, err = pr.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var persons []Domain.Person

	for cursor.Next(context.Background()) {
		var person Domain.Person

		if err := cursor.Decode(&person); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (s *PersonRepoImplement) GetPersonById(id string)(Domain.Person, error){
	var song Domain.Person
	newId, err :=  primitive.ObjectIDFromHex(id)

	if err != nil{
		return Domain.Person{}, err
	}

	err = s.collection.FindOne(context.Background(),bson.M{"_id": newId}).Decode(&song)
	if err != nil{
		return Domain.Person{}, err
	}
	return song, err
}

func (s *PersonRepoImplement) UpdatePerson(id string, person Domain.Person) (Domain.Person, error){
	newId , err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return Domain.Person{}, err
	}

	_, err = s.collection.UpdateOne(context.Background(), bson.M{"_id":newId}, bson.M{"$set": person})
	return person, err
}

func (s *PersonRepoImplement) DeletePerson(id string) (Domain.Person, error){
	newId, err := primitive.ObjectIDFromHex(id) 
	var person Domain.Person

	if err != nil{
		return Domain.Person{}, err
	}

	err = s.collection.FindOneAndDelete(context.Background(), bson.M{"_id": newId}).Decode(&person)
	return person, err

}