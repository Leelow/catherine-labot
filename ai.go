package main

import (
	"time"
	"fmt"
	"strings"
	"math"
	"strconv"
)

var weatherDict = []string{"weather", "sun", "rain", "cloud", "sky"}
var frenchDict = []string{"bonjour", "salut", "coucou", "francais", "français", "france"}
var helloDict = []string{"good morning", "good afternoon", "hello", "what's up"}
var ptitLoupDict = []string{"pti lou", "ptit lou", "ptit loup", "petit loup", "p'tit loup"}

func contains(text string, dict []string) bool {
	for _, word := range dict {
		if (strings.Contains(text, word)) {
			return true
		}
	}
	return false
}

func digitEmoji(n int) string {
	switch(n) {
	case 0: return ":zero:"
	case 1: return ":one:"
	case 2: return ":two:"
	case 3: return ":three:"
	case 4: return ":four:"
	case 5: return ":five:"
	case 6: return ":six:"
	case 7: return ":seven:"
	case 8: return ":eight:"
	case 9: return ":nine:"
	default: return ":keycap_ten:"
	}
}

func numberToEmoji(n int) string {
	if (n < 10) {
		return digitEmoji(n)
	} else {
		return numberToEmoji(n / 10) + digitEmoji(n % 10)
	}
}

func say(handler func(text string), rawText string) {
	if (contains(rawText, ptitLoupDict)) {
		handler(":wolf::wolf: Coucou p'tit loup ! :wolf::wolf:")
		handler("http://static.fnac-static.com/multimedia/Images/FR/NR/3c/23/59/5841724/1540-1/tsp20140506093236/P-tit-Loup-part-en-vacances.jpg")
		return
	}

	if (contains(rawText, helloDict)) {
		handler("Hello, you can ask me the weather for the next five days everywhere all around the world. :smiley:")
		return
	}

	if (contains(rawText, frenchDict)) {
		handler("Sorry, I only speak english, not french ! :flag-mf: :slightly_smiling_face:")
		return
	}

	if (!contains(rawText, weatherDict)) {
		handler("Sorry, I don't understand, I am just a weather bot ! :wink:")
		return
	}

	location, date, fullDate, err := extractLocationAndDate(config.MeaningcloudApiKey, rawText) // extract data from text

	if (err != nil || location == "") {
		handler("Sorry, I don't understand. :cold_sweat:")
		return
	}

	until := time.Until(date).Hours() / 24 // check date is not to far in the future
	if (until > 5) {
		handler("I am sorry but I can only forecast five days in the future. :sunglasses:")
		return
	}

	if (until < 0) {
		date = time.Now()
	}

	fmt.Println("Extracted data: {location : \"" + location + "\", \"" + date.Format("01/02 at 03 pm") + "\"}")

	handler("Hum, let me look... :thinking_face:")

	fmt.Println(location)
	fmt.Println(date)

	wi, err := getForecast(config.OpenweathermapApiKey, location, date) // get forecast

	if (err != nil) {
		handler("Sorry, I don't understand. :cold_sweat:")
		return
	}

	handler(prettyAnswer(location, fullDate, wi))
}

func prettyAnswer(location string, fullDate bool, wi WeatherInformation) (answer string) {
	shortDate := time.Unix(int64(wi.Dt), 0)
	dateFormat := "01/02 at 03 pm"
	if (fullDate == false) {
		dateFormat = "01/02"
	}

	shortDateString := string(shortDate.Format(dateFormat))
	temperature := int(math.Ceil(wi.Main.Temp))
	fmt.Println(shortDateString)

	answer = "The estimated temperature at " + location + " the " + shortDateString + " is " + strconv.Itoa(temperature) + "°C."

	if (len(wi.Weather) > 0) {
		// plural adaptation
		d := wi.Weather[0].Description
		if (d[len(d) - 1:] == "s") {
			answer += " A " + d + " are also planned."
		} else {
			answer += " A " + d + " is also planned."
		}
	}

	// add emoji number for the temperature
	answer += " " + numberToEmoji(temperature)

	// add emoji depending of the description
	switch(wi.Weather[0].Main) {
	case "Clouds":
		answer += ":cloud:"
		break
	case "Rain":
		answer += ":rain_cloud:"
		break
	case "Clear":
		answer += ":sunny:"
		break
	}

	fmt.Println(answer)

	return
}
