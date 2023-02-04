package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string
	Author string
}

func main() {
	var quotes = []Quote{}
	c := colly.NewCollector(colly.AllowedDomains("quotes.toscrape.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Println("Printing: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q := Quote{
			Quote:  quote,
			Author: author,
		}

		quotes = append(quotes, q)
	})

	/* c.OnHTML(".text", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnHTML(".author", func(h *colly.HTMLElement) {
		fmt.Println("--", h.Text)
	}) */

	// To get one random quotes we must use http://quotes.toscrape.com/random
	c.Visit("http://quotes.toscrape.com/")
	// Transforming our data into a JSON format
	json, _ := json.MarshalIndent(quotes, "", " ")
	fmt.Println(string(json))
}
