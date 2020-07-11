package db

import (
	"context"
	"log"
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
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://gadiel:dalvigay@mercadolibre.edcuz.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err.Error())
	}

	coll = client.Database("MercadoLibre").Collection("Mutante")
}

// InsertDna inserta en mongoDb un json con el dna y si es mutante o no.
func InsertDna(request models.Request, mutant bool) error {
	_, err := coll.InsertOne(ctx, bson.M{"dna": request, "mutant": mutant})
	if err != nil {
		return err
	}
	return nil
}
