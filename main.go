package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Quote is a representation of the json quote
type Quote struct {
	Category string `json:"category"`
	IconURL  string `json:"icon_url"`
	ID       string `json:"id"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

func main() {

	res, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	q, err := getQuote([]byte(body))
	fmt.Println(q.Value)
}

func getQuote(body []byte) (*Quote, error) {
	var q = new(Quote)
	err := json.Unmarshal(body, &q)
	if err != nil {
		fmt.Println("Error trying to unmarshal:", err)
	}
	return q, err
}
