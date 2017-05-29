package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Ticker struct {
	Base string `json:"base"`
	Target string `json:"target"`
	Price string `json:"price"`
	Volume string `json:"volume"`
	Change string `json:"change"`
}

type TickerResult struct {
	Ticker Ticker `json:"ticker"`
	Timestamp int `json:"timestamp"`
	Success bool `json:"success"`
	Error string `json:"error"`
}

func GetTickerValue(base string, target string) (*TickerResult) {
	safeBase := url.QueryEscape(base)
	safeTarget := url.QueryEscape(target)

	urlStr := fmt.Sprintf("https://api.cryptonator.com/api/ticker/%s-%s", safeBase, safeTarget)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}
	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record TickerResult

	log.Print(resp.Body)

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return &record
}