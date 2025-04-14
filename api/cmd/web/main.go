package main

import (
	"api/internal/config"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the main function
func main() {

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
