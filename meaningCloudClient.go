package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
	"time"
	"errors"
)

type Status struct {
	Code             string `json:"code"`
	Msg              string `json:"msg"`
	Credits          string `json:"credits"`
	RemainingCredits string `json:"remaining_credits"`
}

type Entity struct {
	Form         string `json:"form"`
	ID           string `json:"id"`
	Sementity    []interface{} `json:"sementity"`
	SemgeoList   []interface{} `json:"semgeo_list"`
	SemldList    []interface{} `json:"semld_list"`
	VariantList  []interface{} `json:"variant_list"`
	Relevance    string `json:"relevance"`
	SemthemeList []interface{} `json:"semtheme_list,omitempty"`
}

type Time struct {
	Form           string `json:"form"`
	NormalizedForm string `json:"normalized_form"`
	ActualTime     string `json:"actual_time"`
	Precision      string `json:"precision"`
	Inip           string `json:"inip"`
	Endp           string `json:"endp"`
}

type mc_Response struct {
	Status     Status`json:"status"`
	EntityList []Entity `json:"entity_list"`
	TimeList   []Time`json:"time_expression_list"`
}

func extractLocationAndDate(api_key string, txt string) (location string, parsedDate time.Time, fullDate bool, err error) {
	response, err := mc_request(api_key, txt)

	if (err != nil) {
		return
	}

	if (len(response.EntityList) == 0) {
		return location, parsedDate, fullDate, errors.New("Cannot extract a location from this sentence.")
	}

	if (len(response.TimeList) == 0) {
		return location, parsedDate, fullDate, errors.New("Cannot extract a date from this sentence.")
	}

	location = response.EntityList[0].Form

	date := response.TimeList[0].ActualTime

	if (len(response.TimeList) > 1 && (response.TimeList[1].Precision == "minutes" || response.TimeList[1].Precision == "hourAMPM")) {
		fullDate = true
		parsedDate, err = time.Parse("2006-01-02 15:04:05 GMT+00:00", date + " " + response.TimeList[1].ActualTime)
	} else {
		fullDate = false
		parsedDate, err = time.Parse("2006-01-02 15:04:05", date + " 12:00:00") // force midday to have a representative temperature
	}

	return
}

func mc_request(api_key string, txt string) (parsedResp mc_Response, err error) {

	form := url.Values{}
	form.Add("key", api_key)
	form.Add("lang", "en")
	form.Add("tt", "et")
	form.Add("txt", txt)

	hc := http.Client{}
	req, err := http.NewRequest("POST", "http://api.meaningcloud.com/topics-2.0", strings.NewReader(form.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	if (err != nil) {
		return
	}

	resp, err := hc.Do(req)

	parsedResp = mc_Response{}
	json.NewDecoder(resp.Body).Decode(&parsedResp)

	return
}
