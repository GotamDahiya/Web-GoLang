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


	username := mux.Vars(request)["name"]

	rows, err := db.Query("select * from User_Profile where like '%" + username + "%'")
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	for rows.Next() {
		rows.Scan(&user.Name, &user.Email, &user.Hobbies, &user.Achivements)
		users = append(users, user)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}
	if jsonError == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-type","application/json")
		response.Write(jsonResponse)
	}
}
	func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
		jsonResponse, err := json.Marshal("It is simple check what you have written")
		if err != nil {
			panic(err)
		}
	
		response.Header().Set("Content-type","application-json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonResponse)
}