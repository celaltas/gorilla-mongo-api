package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Title   string             `bson:"title" json:"title"`
	Author  string             `bson:"author" json:"author"`
	Pages   int                `bson:"pages" json:"pages"`
	Rating  int                `bson:"rating" json:"rating"`
	Genres  []string           `bson:"genres" json:"genres"`
	Reviews []Review           `bson:"reviews" json:"reviews"`
}

type Review struct {
	Name string `bson:"name" json:"name"`
	Body string `bson:"body" json:"body"`
}
