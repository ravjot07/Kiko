package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Problem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title           string             `bson:"title" json:"title"`
	Description     string             `bson:"description" json:"description"`
	Difficulty      string             `bson:"difficulty" json:"difficulty"`
	Requirements    string             `bson:"requirements" json:"requirements"`
	ExampleSolution string             `bson:"exampleSolution" json:"exampleSolution"`
	Tests           []Test             `bson:"tests" json:"tests"`
}

type Test struct {
	Description string `bson:"description" json:"description"`
	TestCode    string `bson:"testCode" json:"testCode"`
}
