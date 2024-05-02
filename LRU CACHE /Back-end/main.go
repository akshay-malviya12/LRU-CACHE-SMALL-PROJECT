package main

import (
	"fmt"
	"log"
	"main/router"
	"net/http"
)

func main() {
	log.Println("Hello, World!")
	fmt.Println("Hello, World!")

	port := "8080"
	server := &http.Server{Addr: "localhost:" + port}

	http.HandleFunc("/", handler)

	fmt.Println("Server started at port 8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error occured while starting server", err)
	} else {
		log.Println("Server started successfully")
	}

}
func handler(w http.ResponseWriter, r *http.Request) {
	router.Router(w, r)
}
