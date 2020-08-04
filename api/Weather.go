package api

import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Res struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    *Weather `json:"data,omitempty"`
	ErrData string  `json:"err_data,omitempty"`
}

type Weather struct {
	City        string                 `json:"city"`
	Time        time.Time              `json:"time"`
	Temperature float64                `json:"Temperature"`
	FeelsLike   float64                `json:"feels_like"`
	MinTemp     float64                `json:"min_temp"`
	MaxTemp     float64                `json:"max_temp"`
	Description string                 `json:"description"`
	MetaData    map[string]interface{} `json:"meta_data"`
}

func WeatherAPIHandler(rs http.ResponseWriter, re *http.Request) {
	cityName := re.FormValue("city")

	switch method := re.Method; method {
	case http.MethodGet:
		res, success := GetWeather(cityName)
		if !success {
			rs.WriteHeader(http.StatusBadRequest)
		} else {
			rs.WriteHeader(http.StatusOK)
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Fprintf(rs, string(resJson))

	}

}

func GetWeather(cityName string) (Res, bool) {

	apiKey := os.Getenv("apiKey")
	weatherUrl := os.Getenv("openWeatherURL")
	weatherUrl = weatherUrl + "weather?q=" + cityName + "&units=metric&appid=" + apiKey

	response, err := http.Get(weatherUrl)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	} else {
		data, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Cannot read response from the api %s\n", err)
			panic(err)
		}

		if response.StatusCode != 200 {
			b, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(b))
			var res Res
			res.Message = "Failed to retrieve weather data"
			res.ErrData = string(data)
			res.Success = false
			return res, false
		}

		var jsonResponse map[string]interface{}
		jsonMarshalErr := json.Unmarshal([]byte(data), &jsonResponse)
		if jsonMarshalErr != nil {
			panic(jsonMarshalErr)
		}

		fmt.Printf("response from weather api is: %s", data)

		var weather Weather

		city, success := jsonResponse["name"].(string)
		if !success {
			panic(success)
		}

		mainTemp := jsonResponse["main"].(map[string]interface{})
		temp := mainTemp["temp"].(float64)
		feelsLike := mainTemp["feels_like"].(float64)
		minTemp := mainTemp["temp_min"].(float64)
		maxTemp := mainTemp["temp_max"].(float64)

		weatherData := jsonResponse["weather"].([]interface{})
		weatherDesc := weatherData[0].(map[string]interface{})["description"].(string)

		weather.City = city
		weather.Time = time.Now()
		weather.Temperature = temp
		weather.FeelsLike = feelsLike
		weather.MinTemp = minTemp
		weather.MaxTemp = maxTemp
		weather.MetaData = jsonResponse
		weather.Description = weatherDesc

		var res Res
		res.Success = true
		res.Message = "Weather data for: " + city + " retrieved successfully"
		res.Data = &weather

		return res, true

	}
}
