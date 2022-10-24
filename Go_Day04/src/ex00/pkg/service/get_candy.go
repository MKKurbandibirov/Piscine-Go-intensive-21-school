package service

import (
	cnd "day04"
	"encoding/json"
	"fmt"
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

func RequestValid(req cnd.CandyRequest) (int, int) {
	if req.CandyCount <= 0 {
		return http.StatusBadRequest, -1
	}
	var status, change int
	for _, val := range candies {
		if val.Name == req.CandyType {
			status, change = http.StatusCreated, req.Money-val.Price*req.CandyCount
			break
		}
	}
	//fmt.Println(status, change)
	if status == 0 && change == 0 {
		return http.StatusBadRequest, -1
	} else if change < 0 {
		//fmt.Println("aaaaaaa")
		return http.StatusPaymentRequired, change
	}
	return http.StatusCreated, change
}

func MakeResponse(w http.ResponseWriter, req cnd.CandyRequest) error {
	status, change := RequestValid(req)
	if status == http.StatusCreated {
		var respCandy cnd.CandyResponse
		respCandy.Change = change
		respCandy.Thanks = "Thank you!"
		data, err := json.MarshalIndent(&respCandy, "", "    ")
		if err != nil {
			return err
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(data); err != nil {
			return err
		}
	} else if status == http.StatusPaymentRequired {
		var errCandy cnd.ErrorResponse
		errCandy.Error = fmt.Sprintf("You need %d more money!", -change)
		data, err := json.MarshalIndent(&errCandy, "", "    ")
		if err != nil {
			return err
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(data); err != nil {
			return err
		}
	}
	return nil
}
