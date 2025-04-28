package main

import (
	"api/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	log.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	handlers.StartActor()

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: ContextMiddleware(LogTraceId(routes(mux))),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
