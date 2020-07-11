package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gadielMa/test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var coll *mongo.Collection
var ctx context.Context

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION")))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer cancel()

	coll = client.Database("MercadoLibre").Collection("Mutante")
}

// InsertDna - inserta en mongoDb un json con el dna y si es mutante o no.
func InsertDna(mutant models.Mutant, isMutant bool) error {
	_, err := coll.InsertOne(ctx, bson.M{"dna": mutant.Dna, "mutant": isMutant})
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

// GetHumans - Trae todos los humanos de la base
func GetHumans() int {
	var total int

	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	for cur.Next(context.TODO()) {
		total++
	}
	return total
}

// GetMutants - Trae solo los mutantes de la base
func GetMutants() int {
	var total int

	cur, err := coll.Find(context.TODO(), bson.D{{"mutant", true}})
	if err != nil {
		log.Fatal(err.Error())
	}
	for cur.Next(context.TODO()) {
		total++
	}
	return total
}
