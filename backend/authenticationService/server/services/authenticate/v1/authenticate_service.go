package authenticate_service

import (
	"authenticationService/models"
	"authenticationService/store"
	"authenticationService/utils"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ValidateIsEmailExists(store *store.Store,email string)error{
	
	var user models.User

	filter := bson.M{"email": email}

	err := store.DB.Collection("user").FindOne(context.Background(),filter).Decode(user)
		
	if(err!=nil){
		return nil
	}else if(user.Email!=""){
		return errors.New("User already exists")
	}

	return nil
}


func SaveUserInDb(store *store.Store,user *models.User) (error){

	hashedPass,err:=utils.HashPassword(user.Password)

	if(err!=nil){
		return err
	}
	user.HashedPassword=hashedPass
	
	// result:=store.DB.Create(&user)
	insertOptions := options.InsertOne().SetBypassDocumentValidation(false)
	fmt.Println("user is",user)
	_,err=store.DB.Collection("users").InsertOne(context.Background(), user, insertOptions)

	if(err!=nil){
		return err
	}
	return nil	
	
}

func FetchUserInDb(store *store.Store,email string) ([]*models.User,error){
	var users []*models.User

	filter := bson.M{"email": email}
	cursor,err:=store.DB.Collection("users").Find(context.Background(),filter)

	if(err!=nil){
		return nil,err
	}

	for cursor.Next(context.Background()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }

	return users,nil	
	
}


func FetchUserByEmailInDb(store *store.Store,email string) (*models.User,error){
	var user *models.User

	filter := bson.M{"email": email}

	err:=store.DB.Collection("users").FindOne(context.Background(),filter).Decode(&user)

	if(err!=nil){
		fmt.Println("error is",err)
		return nil,err
	}

	user.HashedPassword=""

	return user,nil	
	
}

	


