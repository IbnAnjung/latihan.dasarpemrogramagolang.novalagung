package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type student struct {
	Name  string `bson:name`
	Grade int    `bson:Grade`
}

var db *mongo.Database
var err error
var ctx = context.Background()

func connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err.Error())
	}

	db = client.Database("belajar_golang")
}

func insert() {
	connect()
	_, err = db.Collection("student").InsertOne(ctx, student{"Angga", 26})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Evie", 27})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert Success")
}

func find() {
	connect()
	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	defer csr.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	result := []student{}
	for csr.Next(ctx) {
		var row student
		err = csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result > 0) {
		fmt.Println("Name :", result[0].Name)
		fmt.Println("Grade:", result[0].Grade)
	}
}

func main() {
	find()
}
