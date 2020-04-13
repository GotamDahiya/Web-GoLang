package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public", fs))
}

func AddApproutes(route *mux.Router) {

	setStaticFolder(route)
	route.HandleFunc("/",renderHome)
	route.HandleFunc("/users/{name}",getUsers).Methods("GET")
	fmt.Println("Routes are loaded.") 
}