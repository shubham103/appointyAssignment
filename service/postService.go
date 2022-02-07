package service

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/shubham103/Appointy_Assignment/model"
	"github.com/shubham103/Appointy_Assignment/dbservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Funtion to get post model object wiht valuse from query parameter
func getPostFormData(r *http.Request)(model.NewPost){
	temp := model.NewPost{}
	if author:=r.FormValue("authorid"); len(author)>0{
		temp.Author = author
	}
	if postedon:=r.FormValue("postedon"); len(postedon)>0{
		temp.PostedOn = postedon
	}
	if title:=r.FormValue("title"); len(title)>0{
		temp.Title = title
	}
	if body:=r.FormValue("body"); len(body) >0{
		temp.Body = body
	}
	if thumbnail:=r.FormValue("thumbnail"); len(thumbnail)>0{
		temp.Thumbnail = thumbnail
	}

	return temp


}

// Function to validate author id before creating new post
func isValidAuthor(authorid string)bool{

	result := dbservice.GetUserById(authorid)

	if len(result["email"].(string))>0{
		return true
	}
	return false

}

// Function to create new post 
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPost := getPostFormData(r)
	if isValidAuthor(newPost.Author){
		dbservice.CreatePost(w,&newPost)
	} else {
		fmt.Fprintf(w, "invalid author id")
	}


	
}

// Function to delete post 
func PostDelete(w http.ResponseWriter, r *http.Request){
	path := getURLFields(r.URL.Path)
	if len(path)==2{
		dbservice.DeletePostById(w,path[1])
	}
}

// Function to decide GET and PUT for '/post/' url
func PostSubRouter(w http.ResponseWriter, r *http.Request) {

	switch  r.Method {
	case "GET":
			path := getURLFields(r.URL.Path)
			if len(path)==1{
				getAllPost(w,r)
			}
	case "PUT":
		updatePostById(w,r)
		
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

// Function to get app posts
func getAllPost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	results:= dbservice.GetAllPosts()
	json.NewEncoder(w).Encode(results)
}

// Functiont to merge data from url and existing data in db
func mergePostUrlDataDbData(PostDataFromUrl model.NewPost,PostDataFromDb primitive.M)model.NewPost  {
	if len(PostDataFromUrl.Author) == 0{
		PostDataFromUrl.Author = PostDataFromDb["author"].(string)
	}
	if len(PostDataFromUrl.PostedOn) == 0{
		PostDataFromUrl.PostedOn = PostDataFromDb["postedon"].(string)
	}
	if len(PostDataFromUrl.Title) == 0{
		PostDataFromUrl.Title = PostDataFromDb["title"].(string)
	}
	if len(PostDataFromUrl.Body) == 0{
		PostDataFromUrl.Body = PostDataFromDb["body"].(string)
	}
	if len(PostDataFromUrl.Thumbnail) == 0{
		PostDataFromUrl.Thumbnail = PostDataFromDb["thumbnail"].(string)
	}

	return PostDataFromUrl 
} 

// Function to update the post 
func updatePostById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	objId := r.FormValue("_id")
	if len(objId) >0{
		// first get the data from url
		PostDataFromUrl := getPostFormData(r)
		// then get the data from db
		PostDataFromDb := dbservice.GetPostById(objId)
		// update the new values recieved from url
		UpdatedPost := mergePostUrlDataDbData(PostDataFromUrl,PostDataFromDb)
		// then call the updatedb function with the updated newUser
		dbservice.UpdatePostById(w, &UpdatedPost,objId)
	}
}


