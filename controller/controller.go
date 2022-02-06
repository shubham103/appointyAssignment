package controller

import (
	"net/http"

	"Appointy_Assignment/service"
)

func HandleRoutes() {
	// POST request to create new user
	http.HandleFunc("/user/newuser/", service.CreateUser) 
	
	
	// includes GET request for getAlluser, GET request for getUserById,PUT req. for updateUserById 
	http.HandleFunc("/user/", service.UserSubRouter)
	
	
	// POST requeset for authenticating username and password
	http.HandleFunc("/user/userlogin/", service.GetUserByUsernamePassword)



	// POST request to create new post
	http.HandleFunc("/post/create", service.CreatePost)
	
	// GET request for getAllPosts
	// PUT request for updatePostById
	http.HandleFunc("/post/", service.PostSubRouter)

	// DELETE request to deleteById
	http.HandleFunc("/post/delete/", service.PostDelete)

}
