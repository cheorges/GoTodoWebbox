package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection

func init() {
	var dbUrl, dbName, dbCollectionName = loadDbEnv()
	clientOptions := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(dbCollectionName)

	log.Print("Collection instance created!")
}

func loadDbEnv() (string, string, string) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("DB_URI"), os.Getenv("DB_NAME"), os.Getenv("DB_COLLECTION_NAME")
}

func GetAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())

	log.Print(len(results), " Document loaded")
	return results
}

func InsertOneTask(task string) {
	insertResult, err := collection.InsertOne(context.Background(), ToDo{Task: task})

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Inserted a Single Record ", insertResult.InsertedID)
}

func TaskComplete(id string) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result.ModifiedCount, " Modified Document ", _id)
}

func DeleteOneTask(id string) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(d.DeletedCount, " Deleted Document ", _id)
}
