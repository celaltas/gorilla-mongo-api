package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo-api/configs"
	"mongo-api/models"
	"mongo-api/responses"
	"net/http"
	"time"
)

var bookCollection *mongo.Collection = configs.GetCollection(configs.DB, "bookstore", "books")
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func GetBook(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	cursor, err := bookCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var results []models.Book
	for cursor.Next(ctx) {
		var book models.Book
		if err = cursor.Decode(&book); err != nil {
			json.NewEncoder(w).Encode(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		}
		results = append(results, book)
	}
	json.NewEncoder(w).Encode(responses.Response{Status: http.StatusOK, Message: "success", Data: results})

}
func GetBookbyTitle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	cursor := bookCollection.FindOne(ctx, filter)
	var result models.Book
	if err := cursor.Decode(&result); err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	} else {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusOK, Message: "success", Data: result})
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	newBook := models.Book{
		ID:      primitive.NewObjectID(),
		Title:   book.Title,
		Pages:   book.Pages,
		Rating:  book.Rating,
		Genres:  book.Genres,
		Reviews: book.Reviews,
	}

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(responses.Response{Status: http.StatusOK, Message: "success", Data: result})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", bson.D{{"pages", book.Pages}, {"rating", book.Rating}}}}
	result, err := bookCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(responses.Response{Status: http.StatusOK, Message: "success", Data: result})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	result, err := bookCollection.DeleteOne(ctx, filter)
	if err != nil {
		json.NewEncoder(w).Encode(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(responses.Response{Status: http.StatusOK, Message: "success", Data: result})
}
