package dbservice

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/shubham103/Appointy_Assignment/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

)

var DbClient *mongo.Client
var userCollection *mongo.Collection
var postCollection *mongo.Collection

// Function to make connect with db and initialize the DbClient, userCollection and postCollection
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

// Funtion to create new user
func CreateUser(w http.ResponseWriter ,user *model.NewUser){
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Inserted one data", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

// Function to get all user in db
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

// Function to find user by ID
func GetUserById(id string)primitive.M {
	
	objID, _ := primitive.ObjectIDFromHex(id)

	var result primitive.M
	err := userCollection.FindOne(context.TODO(), bson.D{{"_id",objID}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result
}

// Funtion to find use by username and password 
func GetUserByUsernamePassword(username, password string)primitive.M{

	var result primitive.M
	err := userCollection.FindOne(context.TODO(), bson.D{{"username",username},{"password",password}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result

}

// Function to update the user
func UpdateUserById(w http.ResponseWriter, UpdatedUser *model.NewUser, id string){
	objID, _ := primitive.ObjectIDFromHex(id)
	updateResult,_:= userCollection.UpdateOne(context.TODO(),  bson.M{"_id": objID},bson.D{{"$set", UpdatedUser }})
	json.NewEncoder(w).Encode(updateResult.MatchedCount)	
}

// ------------------- Below functions are for post db operations


// Function to create  new post
func CreatePost(w http.ResponseWriter ,post *model.NewPost){
	insertResult, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		fmt.Print(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

// Function to delete post by ID
func DeletePostById(w http.ResponseWriter,id string){

	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, _ := postCollection.DeleteMany(context.TODO(), bson.D{{"_id",objId}})
	json.NewEncoder(w).Encode(deleteResult.DeletedCount)
}

// Function to get all post 
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

// Function to get post by Id
func GetPostById(id string) primitive.M{
	objID, _ := primitive.ObjectIDFromHex(id)

	var result primitive.M
	err := postCollection.FindOne(context.TODO(), bson.D{{"_id",objID}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	return result
}

// Function to update post by Id
func UpdatePostById(w http.ResponseWriter, UpdatedPost *model.NewPost,id string)  {
	objID, _ := primitive.ObjectIDFromHex(id)
	updateResult,_:= postCollection.UpdateOne(context.TODO(),  bson.M{"_id": objID},bson.D{{"$set", UpdatedPost }})
	json.NewEncoder(w).Encode(updateResult.MatchedCount)
}