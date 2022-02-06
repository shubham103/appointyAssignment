package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Dob      string             `json:"dob" bson:"dob"`
	Phone    string             `json:"phone" bson:"phone"`
}

type NewUser struct {
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Dob      string             `json:"dob" bson:"dob"`
	Phone    string             `json:"phone" bson:"phone"`
}
