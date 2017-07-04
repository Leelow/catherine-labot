package main

import (
	"strings"
	"log"
)

var config = getConfig()

func cleanAsk(botId string, text string) string {
	cleanedText := "" // clean the text removing bot @
	for _, part := range strings.Fields(text) {
		if (part != "<@" + botId + ">") {
			cleanedText += part + " "
		}
	}
	return cleanedText
}

func main() {

	// slack connection
	ws, id := slackConnect(config.SlackToken)

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		// if the message is for the bot
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@" + id + ">") {
			ask := cleanAsk(id, m.Text)

			// set message handler
			send := func(text string) {
				m.Text = text
				postMessage(ws, m)
			}

			// call ai.say func
			go say(send, ask)
		}
	}
}
