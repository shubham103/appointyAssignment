package dbservice

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"Appointy_Assignment/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

)

var DbClient *mongo.Client
var userCollection *mongo.Collection
var postCollection *mongo.Collection

func ConnectDb() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

	DbClient = client
	userCollection = DbClient.Database("appointy").Collection("user")
	postCollection = DbClient.Database("appointy").Collection("post")
}



func CreateUser(w http.ResponseWriter ,user *model.NewUser){
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Inserted one data", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

func GetAllUsers() []primitive.M {
	var results []primitive.M
	cur, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) {

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	return results
}

func GetUserById(id string)primitive.M {
	
	objID, _ := primitive.ObjectIDFromHex(id)

	var result primitive.M
	err := userCollection.FindOne(context.TODO(), bson.D{{"_id",objID}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result
}


func GetUserByUsernamePassword(username, password string)primitive.M{

	var result primitive.M
	err := userCollection.FindOne(context.TODO(), bson.D{{"username",username},{"password",password}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result

}


func UpdateUserById(w http.ResponseWriter, UpdatedUser *model.NewUser, id string){
	objID, _ := primitive.ObjectIDFromHex(id)
	updateResult,_:= userCollection.UpdateOne(context.TODO(), bson.D{{"_id",objID}}, UpdatedUser)
	json.NewEncoder(w).Encode(updateResult.MatchedCount)	
}

// ------------------- post db functions

func CreatePost(w http.ResponseWriter ,post *model.NewPost){
	insertResult, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		fmt.Print(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}
func DeletePostById(w http.ResponseWriter,id string){

	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, _ := postCollection.DeleteMany(context.TODO(), bson.D{{"_id",objId}})
	json.NewEncoder(w).Encode(deleteResult.DeletedCount)
}
func GetAllPosts()[]primitive.M {
	var results []primitive.M
	cur, err := postCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()){

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	return results
}
func GetPostById(id string) primitive.M{
	objID, _ := primitive.ObjectIDFromHex(id)

	var result primitive.M
	err := postCollection.FindOne(context.TODO(), bson.D{{"_id",objID}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result
}
func UpdatePostById(w http.ResponseWriter, UpdatedPost *model.NewPost,id string)  {
	objID, _ := primitive.ObjectIDFromHex(id)
	updateResult,_:= postCollection.UpdateOne(context.TODO(), bson.D{{"_id",objID}}, UpdatedPost)
	json.NewEncoder(w).Encode(updateResult.MatchedCount)
}