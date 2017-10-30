package main

import (
	"log"
	"encoding/json"
	"os"
	"flag"
)

type Config struct {
	ApiKey string
	ApiBaseUrl string
	PubgAppId int
}

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config_file_path", "config.json", "Path to the JSON configuration file. See default.json.template for an example")
	flag.Parse()

	config := ReadConfigs(configFilePath)
	log.Print(config)


}

// https://stackoverflow.com/a/16466189/795407
func ReadConfigs(fileName string) *Config {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error reading in file '%s'. Error was: %s", fileName, err)
		os.Exit(-1)
	}

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		log.Fatalf("Error parsing configuration file. Error was: %s", err)
	}
	return config
}
