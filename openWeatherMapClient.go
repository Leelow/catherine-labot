package main

import (
	"net/http"
	"encoding/json"
	"time"
	"math"
	"errors"
)

type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type Weather struct {
	ID          int `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Rain struct {
	ThreeH float64 `json:"3h"`
}

type Sys struct {
	Pod string `json:"pod"`
}

type Coord struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type City struct {
	ID      int `json:"id"`
	Name    string `json:"name"`
	Coord   Coord `json:"coord"`
	Country string `json:"country"`
}

type WeatherInformation struct {
	Dt      int `json:"dt"`
	Main    Main `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  Clouds `json:"clouds"`
	Wind    Wind`json:"wind"`
	Rain    Rain `json:"rain"`
	Sys     Sys `json:"sys"`
	DtTxt   string `json:"dt_txt"`
}

type own_Response struct {
	Cod     string `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int `json:"cnt"`
	List    []WeatherInformation `json:"list"`
	City    City `json:"city"`
}

func getForecast(api_key string, location string, date time.Time) (nearestWeatherInformation WeatherInformation, err error) {
	response, err := owm_request(api_key, location)

	if(err != nil) {
		return
	}

	if (response.Cod != "200") {
		return nearestWeatherInformation, errors.New("An error occured during request.")
	}

	nearestWeatherInformation = response.List[0]
	minimumGap := math.Inf(+1)
	for _, list := range response.List {
		gap := float64(date.Unix()) - float64(list.Dt)
		if (gap <= minimumGap && gap >= 0) {
			minimumGap = gap
			nearestWeatherInformation = list
		}
	}

	return
}

func owm_request(api_key string, location string) (parsedResp own_Response, err error) {
	hc := http.Client{}
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/forecast?units=metric&q=" + location + "&appid=" + api_key, nil)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	if (err != nil) {
		return
	}

	resp, err := hc.Do(req)

	parsedResp = own_Response{}
	json.NewDecoder(resp.Body).Decode(&parsedResp)

	return
}

