package main

import (
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	url := "https://www.google.com"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}