package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	//2 ways to initialize empty map:
	//var results = map[string]string{}
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.google.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		fmt.Println("Checking ", url)
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
