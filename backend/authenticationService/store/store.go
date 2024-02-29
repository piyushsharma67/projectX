package store

import "database/sql"

type Store struct{
	DB *sql.DB
}
func New(dataSourceName string) *Store{

	s:=&Store{}
	db,error:=sql.Open("mysql",dataSourceName)

	if(error!=nil){
		return nil
	}

	err:=db.Ping()
	if(err!=nil){
		return nil
	}

	s.DB=db

	return s

}