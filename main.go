package main

import (
	"../weather-api/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/weather", api.WeatherAPIHandler)
	err := http.ListenAndServe(port(), nil)

	if err != nil {
		panic(err)

	}

}

func index(rs http.ResponseWriter, re *http.Request) {
	fmt.Println("Everything looks okay. The server is running on the port: " + port())
	rs.WriteHeader(http.StatusOK)
	fmt.Fprintf(rs, "Everything is up and running")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
