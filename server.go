// but first run the program in the terminal/command prompt "go run server.go"
// Run http:/localhost:8080 in the web browser
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server will start at http://localhost:8000")
	connectDB()

	route := mux.NewRouter()

	AddApproutes(route)

	log.Fatal(http.ListenAndServe(":8000", route))
}
