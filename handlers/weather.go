package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	APP_ID                 = "0wteVDjV34EbGx294H6lTFdJL4tobqIMSUu9LX9fNNpDPt6E.GngOgbsqbwy5fR8StxPdA--"
	YAHOO_WEATHER          = "https://query.yahooapis.com/v1/public/yql?&format=json"
	YAHOO_WEATHER_ENDPOINT = YAHOO_WEATHER + "&appid=" + APP_ID + "&q="
)

type YahooWeather struct {
	Forecast string `json:"query>results>channel>item>condition>text"`
	Query    struct {
		Results struct {
			Channel struct {
				Item struct {
					Title     string
					Condition struct {
						Text string
						Temp string
					}
				}
			}
		}
	}
}

func WeatherHandler(msg string) string {
	var location string
	text := strings.Replace(msg, "!weather", "", 1)
	trimmed_text := strings.Trim(text, " ")

	if trimmed_text == "" {
		location = "toronto, canada"
	} else {
		location = trimmed_text
	}
	return GetWeatherForecast(location)
}

func GetWeatherForecast(location string) string {
	query := "select * from weather.forecast where woeid in " +
		"(select woeid from geo.places(1) where text=\"" +
		location + "\") and u='c'"

	query_escaped := url.QueryEscape(query)
	resp, err := http.Get(YAHOO_WEATHER_ENDPOINT + query_escaped)

	if err != nil {
		log.Fatal(err)
		return "Booboo :("
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	reply := new(YahooWeather)
	decoder.Decode(reply)

	title := reply.Query.Results.Channel.Item.Title
	condition := reply.Query.Results.Channel.Item.Condition.Text
	temperature := reply.Query.Results.Channel.Item.Condition.Temp
	return title + ": " + condition + " " + temperature + "°C"
}
