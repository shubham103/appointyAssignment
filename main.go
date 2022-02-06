package main

import (
	"Appointy_Assignment/controller"
	"Appointy_Assignment/dbservice"
	"fmt"
	"log"
	"net/http"
)



func main() {

	dbservice.ConnectDb()
	controller.HandleRoutes()
	fmt.Println("server started at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
