package main

import (
	"Unit-Converter/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Unit Converter at http://localhost:8080")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/length", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeightHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
