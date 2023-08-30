package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	SERVER_URL = "http://localhost:8080/cotacao"
	TIMEOUT_MS = 300
)

type Exchange struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {

	price := getPrice()
	saveToFile(price)

}

func getPrice() string {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MS*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SERVER_URL, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var exchange Exchange

	if err := json.NewDecoder(res.Body).Decode(&exchange); err != nil {
		panic(err)
	}
	return exchange.Usdbrl.Bid
}

func saveToFile(price string) {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", price))
	if err != nil {
		panic(err)
	}
}
