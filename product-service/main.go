package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8081"

func getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "product endpoint")
}

func main() {
	log.Println("starting product server ...")

	http.HandleFunc("/product", getProduct)
	http.ListenAndServe(port, nil)

	// Spinning up the server.
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
