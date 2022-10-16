package main

import (
	"day02/internal/domain"
	"day02/internal/handler"
	"log"
	"sync"
)

func main() {
	flags := domain.NewFlags()
	if !flags.A {
		_, err := handler.OnceHandler(flags.Files[0], "")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		archName, err := handler.CreateArchivesDir(flags.Archive)
		if err != nil {
			log.Fatalln(err)
		}
		wg := new(sync.WaitGroup)
		for _, file := range flags.Files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				_, err := handler.OnceHandler(file, archName)
				if err != nil {
					log.Fatalln(err)
				}
			}(file)
		}
		wg.Wait()
	}
}
