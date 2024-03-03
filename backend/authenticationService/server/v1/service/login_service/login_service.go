package login_services

import (
	"authenticationService/models"
	"authenticationService/store"
)

func FetchUserByEmailInDb(store *store.Store,email string) (*models.User,error){
	var users *models.User
	result:=store.DB.First(&users,"Email = ?",email)

	if(result.Error!=nil){
		return nil,result.Error
	}

	users.HashedPassword=""

	return users,nil	
	
}

