package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xanity-07/apiProject/internal/app"
)

func main() {
	//* Instanciating our App
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Our App is running")

	//* Bind the health check
	http.HandleFunc("/health", HealthCheck)

	//* Creating server
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

// * Health check for recieving requests
// * w to communicate back to the client
// * r is for what the caller is sending us
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
