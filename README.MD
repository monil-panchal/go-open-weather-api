# Go lang application for fetching weather data using OpenWeatherMap API

This is a Go lang web application built `net/http` library for getting live weather data from the OpenWeatherMap API.

## Usage

This application exposes the following API endpoints: 

- Get current weather data for the city
 
API uri: `/api/weather?city=Halifax` method: `GET`

Response: 

```
{
"success": true,
"message": "Weather data for: Halifax retrieved successfully",
"data": {
"city": "Halifax",
"time": "2020-08-04T01:43:43.866217604-03:00",
"Temperature": 18.94,
"feels_like": 18.07,
"min_temp": 18.89,
"max_temp": 19,
"description": "clear sky",
"meta_data": {
"base": "stations",
"clouds": {
"all": 5
},
"cod": 200,
"coord": {
"lat": 44.65,
"lon": -63.57
},
"dt": 1596515739,
"id": 6324729,
"main": {
"feels_like": 18.07,
"humidity": 93,
"pressure": 1014,
"temp": 18.94,
"temp_max": 19,
"temp_min": 18.89
},
"name": "Halifax",
"sys": {
"country": "CA",
"id": 822,
"sunrise": 1596531872,
"sunset": 1596584165,
"type": 1
},
"timezone": -10800,
"visibility": 10000,
"weather": [
{
"description": "clear sky",
"icon": "01n",
"id": 800,
"main": "Clear"
}
],
"wind": {
"deg": 220,
"speed": 5.1
}
}
}
}

```

#### Create account on [OpenWeatherMap](https://openweathermap.org/api)
Create a free account on OpenWeatherMap and choose the current weather api, and generate the api key.
 
## Installation

Import this project in any Go lang supported IDE.

#### Adding environment variables
```
Add the `openWeatherURL` and `apiKey` in the environment variables.
The default `openWeatherURL` should be `http://api.openweathermap.org/data/2.5/` unless changed by the OpenWeatherMap.
```

#### Running the project
```
go build
```

```
./weather-api
```

#### Sanity testing

```
http://localhost:8080/
```

response:

 `Everything is up and running`
