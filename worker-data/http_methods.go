package main

import (
	"io"
	"log"
	"net/http"
)

func _httpGET(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
