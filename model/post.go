package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Author    string `json:"authorid" bson:"authorid"`
	PostedOn  string `json:"postedon" bson:"postedon"`
	Title     string `json:"title" bson:"title"`
	Body      string `json:"body" bson:"body"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
}


type NewPost struct {
	Author    string `json:"authorid" bson:"authorid"`
	PostedOn  string `json:"postedon" bson:"postedon"`
	Title     string `json:"title" bson:"title"`
	Body      string `json:"body" bson:"body"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
}

