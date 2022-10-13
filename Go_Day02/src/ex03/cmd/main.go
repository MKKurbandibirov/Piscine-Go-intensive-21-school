package main

import (
	"day02/internal/domain"
	"day02/internal/handler"
	"fmt"
	"log"
)

func main() {
	flags := domain.NewFlags()
	if !flags.A {
		target, err := handler.OnceHandler(flags.Files[0])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(target)
	}
}
