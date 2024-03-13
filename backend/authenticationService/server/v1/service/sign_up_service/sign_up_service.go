package sign_up_services

import (
	"authenticationService/models"
	"authenticationService/store"
	"authenticationService/utils"
	"errors"

	"gorm.io/gorm"
)

func ValidateIsEmailExists(store *store.Store,email string)error{
	
	var user models.User

	if err := store.DB.First(&user,"Email = ?",email).Error; err != nil {
		// Handle error...
		if(err==gorm.ErrRecordNotFound){
			return nil
		}
	  }
	// User found, email already exists
	return errors.New("user with this email already exists")
	
}


func SaveUserInDb(store *store.Store,user *models.User) (error){

	hashedPass,err:=utils.HashPassword(user.Password)

	if(err!=nil){
		return err
	}
	user.HashedPassword=hashedPass
	
	result:=store.DB.Create(&user)

	if(result.Error!=nil){
		return result.Error
	}
	return nil	
	
}

func FetchUserInDb(store *store.Store,email string) ([]*models.User,error){
	var users []*models.User
	result:=store.DB.First(&users,"Email = ?",email)

	if(result.Error!=nil){
		return nil,result.Error
	}

	return users,nil	
	
}

func DropUsertable(store *store.Store)error{
	if err:=store.DB.Migrator().DropTable(&models.User{});err!=nil{
		return err
	}
	return nil
}

func FetchUserByEmailInDb(store *store.Store,email string) (*models.User,error){
	var users *models.User
	result:=store.DB.First(&users,"Email = ?",email)

	if(result.Error!=nil){
		return nil,result.Error
	}

	users.HashedPassword=""

	return users,nil	
	
}

	
