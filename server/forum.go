package main

import (
	"encoding/json"
	"github.com/fberrez/forum/datastore"
	"github.com/fberrez/forum/jsonconfig"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database datastore.Info `json:"Database"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func main() {
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	datastore.Connect(config.Database)

	/* Graceful stop */
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChannel
		stop()
		os.Exit(0)
	}()

	router := getRouter()
	router.Run(":8080")
}

func stop() {
	log.Println("Gracefully stopping...")
}
