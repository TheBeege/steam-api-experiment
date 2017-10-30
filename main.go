package main

import (
	"log"
	"encoding/json"
	"os"
	"flag"
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"github.com/TheBeege/steam-api-experiment/dto"
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

	playerCountData := GetPlayerCountData(config)
	if playerCountData == nil {
		log.Fatal("Failed to collect player count data.")
	}
	log.Printf("Count players: %d -- Result: %d", playerCountData.Response.PlayerCount, playerCountData.Response.Result)

	assetPriceData := GetPriceData(config)
	if assetPriceData == nil {
		log.Fatal("Failed to collect price data.")
	}
	log.Printf("Prices\n\tKRW: %d\n\tUSD: %d\n\tCNY: %d", assetPriceData.Result.Assets[0].Prices.KRW, assetPriceData.Result.Assets[0].Prices.USD, assetPriceData.Result.Assets[0].Prices.CNY)
}

func GetPlayerCountData(config *Config) (*dto.PlayerCountResponse) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%sISteamUserStats/GetNumberOfCurrentPlayers/v1/", config.ApiBaseUrl), nil)
	if err != nil {
		log.Printf("Error creating request for player count data. Error was: %s", err)
		return nil
	}
	queryParams := request.URL.Query()
	queryParams.Add("appid", strconv.Itoa(config.PubgAppId))
	request.URL.RawQuery = queryParams.Encode()
	log.Print(request.URL.String())

	headers := request.Header
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error in response when collecting player count data. Error was: %s", err)
		return nil
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body when collecting player count data. Error was: %s", err)
		return nil
	}
	var playerCountData = new(dto.PlayerCountResponse)
	err = json.Unmarshal(responseBody, &playerCountData)
	if err != nil {
		log.Printf("Error decoding response body for player count data. Error was: %s", err)
		log.Printf("Response body was: %s", responseBody)
		return nil
	}
	return playerCountData
}

func GetPriceData(config *Config) (*dto.AssetPriceResponse) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%sISteamEconomy/GetAssetPrices/v1/", config.ApiBaseUrl), nil)
	if err != nil {
		log.Printf("Error creating request for price data. Error was: %s", err)
		return nil
	}
	queryParams := request.URL.Query()
	queryParams.Add("appid", strconv.Itoa(config.PubgAppId))
	queryParams.Add("key", config.ApiKey)
	request.URL.RawQuery = queryParams.Encode()
	log.Print(request.URL.String())

	headers := request.Header
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error in response when collecting price data. Error was: %s", err)
		return nil
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body when collecting price data. Error was: %s", err)
		return nil
	}
	var assetPriceData = new(dto.AssetPriceResponse)
	err = json.Unmarshal(responseBody, &assetPriceData)
	if err != nil {
		log.Printf("Error decoding response body for price data. Error was: %s", err)
		log.Printf("Response body was: %s", responseBody)
		return nil
	}
	return assetPriceData
}

// https://stackoverflow.com/a/16466189/795407
func ReadConfigs(fileName string) *Config {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error reading in file '%s'. Error was: %s", fileName, err)
	}

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		log.Fatalf("Error parsing configuration file. Error was: %s", err)
	}
	return config
}
