package model

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database struct
type Database struct {
	Self   *mongo.Database
	Client *mongo.Client
}

const Millisecond int64 = 1000

//DB instance
var DB *Database

//Init database init
func (db *Database) Init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(viper.GetString("db.addr"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	DB = &Database{
		Client: client,
		Self:   client.Database(viper.GetString("db.name")),
	}
}
func (db *Database) Close() {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
