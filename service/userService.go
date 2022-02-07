package service

import (
	"fmt"
	"encoding/json"
	"strings"
	"net/http"
	"github.com/shubham103/Appointy_Assignment/model"
	"github.com/shubham103/Appointy_Assignment/dbservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to convert query data into usr struct obj
func getFormData(r *http.Request)(model.NewUser){

	temp := model.NewUser{}
	if name:=r.FormValue("name"); len(name)>0{
		temp.Name = name
	}
	if email:=r.FormValue("email"); len(email)>0{
		temp.Email = email
	}
	if username:=r.FormValue("username"); len(username)>0{
		temp.Username = username
	}
	if password:=r.FormValue("password"); len(password) >0{
		temp.Password = password
	}
	if dob:=r.FormValue("dob"); len(dob)>0{
		temp.Dob = dob
	}
	if phone:=r.FormValue("phone"); len(phone)>0{
		temp.Phone = phone
	}

	return temp


}

// function to create new user , return the objectId from mongodb
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	newUser := getFormData(r)
	dbservice.CreateUser(w,&newUser)

}

// function to break the url to the id from it 
func getURLFields(str string)[]string{
	res := []string{}
	temp := strings.Split(str, "/")
	for _,value := range temp{
		if len(value)>0{
			res = append(res, value)
		}
	}
	return res

}

// function to decide between getAll, getById, updateById
func UserSubRouter(w http.ResponseWriter, r *http.Request)  {

	switch  r.Method {
	case "GET":
			path := getURLFields(r.URL.Path)
			if len(path)==1{
				getAllUser(w,r)
			} else if len(path)==2{
				getUserById(w,r, path[1])
			}
			return
	case "PUT":
		updateUserById(w,r)
		
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}
	
}

// function to merge the data from url and data in db
func mergeUrlDataDbData(UserDataFromUrl model.NewUser,UserDataFromDb primitive.M)model.NewUser{

	if len(UserDataFromUrl.Name) == 0{
		UserDataFromUrl.Name = UserDataFromDb["name"].(string)
	}
	if len(UserDataFromUrl.Email) == 0{
		UserDataFromUrl.Email = UserDataFromDb["email"].(string)
	}
	if len(UserDataFromUrl.Username) == 0{
		UserDataFromUrl.Username = UserDataFromDb["username"].(string)
	}
	if len(UserDataFromUrl.Password) == 0{
		UserDataFromUrl.Password = UserDataFromDb["password"].(string)
	}
	if len(UserDataFromUrl.Dob) == 0{
		UserDataFromUrl.Dob = UserDataFromDb["dob"].(string)
	}
	if len(UserDataFromUrl.Phone )== 0{
		UserDataFromUrl.Phone = UserDataFromDb["phone"].(string)
	}

	return UserDataFromUrl

}

// function to update User details
func updateUserById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	objId := r.FormValue("_id")
	if len(objId) >0{
		// first get the data from url
		UserDataFromUrl := getFormData(r)
		// then get the data from db
		UserDataFromDb := dbservice.GetUserById(objId)
		// update the new values recieved from url
		UpdatedUser := mergeUrlDataDbData(UserDataFromUrl,UserDataFromDb)
		// then call the updatedb function with the updated newUser
		fmt.Println(UpdatedUser)
		dbservice.UpdateUserById(w, &UpdatedUser,objId)
	}
}

// function to get all users
func getAllUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	results:= dbservice.GetAllUsers()
	json.NewEncoder(w).Encode(results)
}

// function to get user by id
func getUserById(w http.ResponseWriter, r *http.Request, id string){

	w.Header().Set("Content-Type", "application/json")
	result := dbservice.GetUserById(id)
	json.NewEncoder(w).Encode(result)
	

}

// function to find usrname and password from query parameters
func getUsernamePasswordFromUrl(r *http.Request)([]string){
	username:=r.FormValue("username")
	password:=r.FormValue("password")
	if len(username)>0{
		if len(password) > 0{
			return []string{username, password}
		}
		return []string{username}
	}
	return []string{}


}

// function to get usrer by username and password
func GetUserByUsernamePassword(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	usernamePassword := getUsernamePasswordFromUrl(r)

	if len(usernamePassword) == 2{
		result := dbservice.GetUserByUsernamePassword(usernamePassword[0],usernamePassword[1])
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
	}

	
}

