package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();
	const port string = ":7000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Fprintln(res, "Up and running...")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
} 