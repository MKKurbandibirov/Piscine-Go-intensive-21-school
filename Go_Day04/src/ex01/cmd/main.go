package main

import (
	"context"
	"day04/pkg/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/buy_candy", handler.BuyCandyHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error listening for server: %s", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server couldn't close normally: %s", err)
	}
	log.Println("Server shutdown!")
}
