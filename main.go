package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote      string `json:"quote"`
			Length     int    `json:"length"`
			Author     string `json:"author"`
			Category   string `json:"category"`
			Language   string `json:"language"`
			Date       string `json:"date"`
			Permalink  string `json:"permalink"`
			ID         string `json:"id"`
			Background string `json:"background"`
			Title      string `json:"title"`
		} `json:"quotes"`
	} `json:"contents"`
	Baseurl   string `json:"baseurl"`
	Copyright struct {
		Year string `json:"year"`
		URL  string `json:"url"`
	} `json:"copyright"`
}

func main() {
	request, err := http.NewRequest("GET", "https://quotes.rest/qod?language=en&quot", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-TheySaidSo-Api-Secret", "PjMpSvDAM2rwhtiD6q279G2GW7wQGez6mJsBmxQi")
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	fmt.Println("Response body:", res.Body)

	response := &Response{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		panic(err)
	}
	quote := response.Contents.Quotes[0]
	fmt.Println("Quote", quote.Quote)
	fmt.Println("Author", quote.Author)

	f, err := os.Create("README.md")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(">%s\n\n**%s**", quote.Quote, quote.Author))

	if err != nil {
		panic(err)
	}
}
