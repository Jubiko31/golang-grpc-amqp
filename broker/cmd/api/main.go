package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8000"

type Config struct {
}

func main() {
	app := Config{}

	log.Println("🐋 Starting Broker Service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
