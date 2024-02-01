package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to a class of Modules...")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	//	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func greeter() {
	fmt.Println("Hey there mod users!")
}
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
