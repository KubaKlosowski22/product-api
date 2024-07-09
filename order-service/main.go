package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8082"

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "order endpoint")
}

func main() {
	log.Println("starting order server ...")

	http.HandleFunc("/order", getOrder)
	http.ListenAndServe(port, nil)

	// Spinning up the server.
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
