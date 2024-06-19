package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rishabhcodes65536/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb+srv://Rishabh:.e7vNcgwJMSz_sn@cluster0.p3sj2tx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB

func init(){
	//client option
	clientOption := options.Client().ApplyURI(CONNECTION_STRING)

	//connect to mongoDB

	//context is passed coz the machine is different it may be miles away TODO is used when we dont know what to do
	client,err := mongo.Connect(context.TODO(), clientOption)

	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance more as like reference

	fmt.Println("Collection instance is ready")
}

//MONGO DB HELPERS


//INSERT 1 RECORD

//i lower case is helper we wont export this

func insertOneMovie(movie model.Netflix){
	insertResult,err := collection.InsertOne(context.TODO(),movie)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id", insertResult.InsertedID)
}

//update 1 movie


func updateOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	updateFilter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result,err := collection.UpdateOne(context.Background(), updateFilter, update)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Ypdated 1 movie in db with id", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount,err := collection.DeleteOne(context.Background(), filter)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted movie in db with count: ", deleteCount)
}

//delete all movies
func deleteAllMovie() int64{
	deleteResult,err := collection.DeleteOne(context.Background(), bson.D{{}}, nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted movie in db with count: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}


//read all document from the collection

func getAllMovies() []primitive.M{
	cursor,err := collection.Find(context.Background(),bson.D{{}})


	if err !=nil {
		log.Fatal(err)
	}

	var movies []primitive.M  
	// this is almost same as bson.M just reliable in here bson.M working for mw

	for cursor.Next(context.TODO()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}


//Actual controllers

func GetMyAllMovies (w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie (w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count:= deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}