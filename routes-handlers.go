package main

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string //Name of user in  database
	Email string //Email of user
	Hobbies string // Hobbies of user
	Achivements string // Achivements of user
}

func renderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./views/index.html")
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	var (
		user User
		users []User
	)
	// var (
	// 	name string
	// 	email string
	// 	hobbies string
	// 	achivements string
	// )

	username := mux.Vars(request)["name"]

	rows, err := db.Query("select Name,Email,Hobbies,Achivements from User_Profile where UserName like '%"+username+"%'")
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	for rows.Next() {
		rows.Scan(&user.Name, &user.Email, &user.Hobbies, &user.Achivements)
		// fmt.Println(&user.Name, &user.Email, &user.Hobbies, &user.Achivements)
		users = append(users, user)
		// fmt.Println(users)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}
	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-type","application/json")
		response.Write(jsonResponse)
	}
}
	func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
		jsonResponse, err := json.Marshal("Please check your input")
		if err != nil {
			panic(err)
		}
	
		response.Header().Set("Content-type","application-json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonResponse)
}