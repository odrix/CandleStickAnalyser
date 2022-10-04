package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/binance-exchange/go-binance"
	"klintt.io/detect/detectdaily"
)

type Model struct {
	Pairs   []string `json:"pairs"`
	OnlyFor string   `json:"onlyFor"`
}

func Handle(w http.ResponseWriter, r *http.Request) {

	model := Model{}

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// detectdaily.DetectAndEmail(model.Pairs, model.OnlyFor, binance.Day)
	detectdaily.DetectAndTweet(model.Pairs, binance.Hour)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(""))
}
