package main

import (
	"encoding/json"
	"os"
	"runtime"
	"errors"
	"path"
)

type Config struct {
	SlackToken           string `json:"slack_token"`
	OpenweathermapApiKey string `json:"openweathermap_api_key"`
	MeaningcloudApiKey   string `json:"meaningcloud_api_key"`
}

func getConfigFilePath() (configPath string, err error) {
	_, scriptPath, _, ok := runtime.Caller(1)
	if (ok == false) {
		err = errors.New("An error occured getting config file.")
		return
	}

	configPath = path.Join(path.Dir(scriptPath), "./config/config.json")
	return
}

func getConfig() (config Config) {
	configPath, err := getConfigFilePath()
	if (err != nil) {
		panic(err)
	}

	file, err := os.Open(configPath)
	if (err != nil) {
		panic(err)
	}

	config = Config{}
	json.NewDecoder(file).Decode(&config)
	file.Close()

	return
}