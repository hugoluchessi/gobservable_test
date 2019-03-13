package main

import (
	"net/http"

	"github.com/hugoluchessi/gobservable_test/config"
)

func main() {
	ms, err := config.NewMonitorServices()

	if err != nil {
		panic(err)
	}

	mux := ConfigureRoutes(ms)

	http.ListenAndServe(":8080", mux)
}
