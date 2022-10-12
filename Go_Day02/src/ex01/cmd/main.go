package main

import (
	"day02/internal/domain"
	"day02/internal/handler"
	"log"
)

func main() {
	flags, err := domain.NewFlags()
	if err != nil {
		log.Fatalln(err)
	}
	handler.Handle(*flags)
}
