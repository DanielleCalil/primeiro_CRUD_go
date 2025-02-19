package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Email    string             `bson:"email, omitempty"`
	Password string             `bson:"password, omitempty"`
	Name     string             `bson:"name, omitempty"`
	Age      int                `bson:"age, omitempty"`
}
