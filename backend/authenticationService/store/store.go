package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct{
	DB *mongo.Database
}

func New(connectionString string) *Store{
	
	s:=&Store{}
	clientOption:=options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOption)

	if(err!=nil){
		log.Fatal(err)
	}

	s.DB=client.Database("schoolApp")

	return s

}