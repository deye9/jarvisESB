package main

import (
	app "jarvisESB/application"
	"log"
	"net/http"
)

func main() {
	// Load the env file by default
	app.GetEnvironment(".env")

	apiURL := app.GetEnv("API_URL")
	apiPort := app.GetEnv("API_PORT")

	log.Println("Application is listening on " + apiURL + "index")
	log.Fatal(http.ListenAndServe(":"+apiPort, nil))
}
