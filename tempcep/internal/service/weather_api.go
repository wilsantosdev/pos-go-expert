package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	WEATHER_API_URL = "https://api.weatherapi.com/v1/current.json"
)

type weatherApi struct {
	apiKey string
}

type weatherAPIResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}
type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
type Current struct {
	LastUpdatedEpoch int       `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

func NewWeatherApi(apiKey string) weatherApi {
	return weatherApi{apiKey: apiKey}
}

func (w weatherApi) GetTemperatureByCity(city string) (celsius, farenheit, kelvin float64, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?key=%s&q=%s", WEATHER_API_URL, w.apiKey, url.QueryEscape(city)), nil)
	if err != nil {
		return 0, 0, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return 0, 0, 0, err
	}

	defer resp.Body.Close()

	var data weatherAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, 0, err
	}

	return data.Current.TempC, data.Current.TempF, data.Current.TempC, nil
}
