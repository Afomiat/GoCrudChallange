package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Age int `json:"age" bson:"age"`
	Hobbies []string `json:"hobbies" bson:"hobbies"`
}