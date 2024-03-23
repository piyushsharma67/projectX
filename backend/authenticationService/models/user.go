package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
    ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name           string             `bson:"name" json:"name"`
	Email          string             `bson:"email" json:"email"`
	Mobile         string             `bson:"mobile" json:"mobile"`
	Password       string             `bson:"password" json:"password"`
	HashedPassword string             `bson:"hashed_password,omitempty" json:"-"`

}

type UserLoginDetails struct{
	User *User
	Token string
}