package service

import (
	cnd "day04"
	"net/http"
)

var candies = []cnd.Candy{
	{
		Name:  "CE",
		Price: 10,
	},
	{
		Name:  "AA",
		Price: 15,
	},
	{
		Name:  "NT",
		Price: 17,
	},
	{
		Name:  "DE",
		Price: 21,
	},
	{
		Name:  "YR",
		Price: 23,
	},
}

func MakeResponse(req cnd.CandyRequest) *http.Response {
	var status int
	for _, val := range candies {
		if val.Name == req.CandyType && req.Money >= val.Price {
			status = http.StatusCreated
			break
		}
	}
	if status == 0 {
		status = http.StatusBadRequest
	}
	if req.CandyCount <= 0 {
		status = http.StatusBadRequest
	}

	//response := &http.Response{
	//
	//}
}
