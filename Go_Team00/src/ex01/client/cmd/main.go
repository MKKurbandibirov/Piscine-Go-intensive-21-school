package main

import (
	"client/internal/adapters"
	"context"
	"log"
)

func main() {
	client := &adapters.Client{}
	err := client.Connect("localhost:4000")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Disconnect()

	if err := client.GetStatistics(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
