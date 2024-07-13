package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Problem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title"`
	Description     string             `bson:"description"`
	Difficulty      string             `bson:"difficulty"`
	Requirements    string             `bson:"requirements"`
	ExampleSolution string             `bson:"exampleSolution"`
}
