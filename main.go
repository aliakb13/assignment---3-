package main

import (
	"assignment3/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/weather", handler.GetWeather)

	http.ListenAndServe("localhost:8082", nil)

}
