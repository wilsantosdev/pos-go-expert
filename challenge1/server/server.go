package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	EXCHANGE_URL        = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	TIMEOUT_MS          = 200
	DATABASE_TIMEOUT_MS = 10
	SQLITE_DB           = "./exchange.db"
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

	createTableExchange()

	http.HandleFunc("/cotacao", exchangeHandler)
	http.ListenAndServe(":8080", nil)
}

func exchangeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exchange := getExchange(ctx)
	saveExchange(ctx, exchange)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(exchange)
	if err != nil {
		panic(err)
	}
	w.Write(response)

}

func getExchange(ctx context.Context) Exchange {
	ctx, cancel := context.WithTimeout(ctx, TIMEOUT_MS*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, EXCHANGE_URL, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var exchange Exchange

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &exchange)

	return exchange
}

func saveExchange(ctx context.Context, exchange Exchange) error {
	db, err := sql.Open("sqlite3", SQLITE_DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(ctx, DATABASE_TIMEOUT_MS*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-time.After(DATABASE_TIMEOUT_MS * time.Millisecond):
		stmt, err := db.Prepare("INSERT INTO exchange(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) values(?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(exchange.Usdbrl.Code, exchange.Usdbrl.Codein, exchange.Usdbrl.Name, exchange.Usdbrl.High, exchange.Usdbrl.Low, exchange.Usdbrl.VarBid, exchange.Usdbrl.PctChange, exchange.Usdbrl.Bid, exchange.Usdbrl.Ask, exchange.Usdbrl.Timestamp, exchange.Usdbrl.CreateDate)
		return err

	}
	return nil

}

func createTableExchange() {
	db, err := sql.Open("sqlite3", SQLITE_DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS exchange (code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, createDate TEXT)")
	if err != nil {
		panic(err)
	}
}
