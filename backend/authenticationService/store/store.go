package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct{
	DB *gorm.DB
}


func New(dataSourceName string) *Store{
	
	s:=&Store{}
	db,error:=gorm.Open(mysql.Open(dataSourceName),&gorm.Config{})

	if(error!=nil){
		return nil
	}

	s.DB=db

	return s

}