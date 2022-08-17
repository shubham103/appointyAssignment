package main

import (
	"github.com/shubham103/Appointy_Assignment/controller"
	"github.com/shubham103/Appointy_Assignment/dbservice"
	"fmt"
	"log"
	"net/http"
)



func main() {
	// to make connection with mongodb
	//dbservice.ConnectDb()
	controller.HandleRoutes()
	fmt.Println("server started at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
