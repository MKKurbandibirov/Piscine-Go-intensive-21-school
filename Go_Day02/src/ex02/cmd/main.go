package main

import (
	"day02/internal/handlers"
	"fmt"
	"log"
	"os"
)

func main() {
	args := handlers.GetArgs()
	result, err := handlers.Handler(os.Args[1], args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(result)
}
