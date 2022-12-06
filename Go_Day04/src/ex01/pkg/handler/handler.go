package handler

import (
	candies "day04"
	"day04/pkg/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func BuyCandyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("request does not have body")
		return
	}
	var candyReq candies.CandyRequest
	err = json.Unmarshal(body, &candyReq)
	if err != nil {
		log.Println("couldn't unmarshal body")
		return
	}
	if err := service.MakeResponse(w, candyReq); err != nil {
		log.Println("couldn't create response")
	}
}
