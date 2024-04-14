package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	BaseURL      string
	ApiPre       string
	ApiVer       string
	ApiEndpoint  string
	ApiKey       string
	City         string
	Units		 string
}

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getEnvVars() EnvVars {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return EnvVars{
		BaseURL:      os.Getenv("BASE_URL"),
		ApiPre:       os.Getenv("OWM_API_PRE"),
		ApiVer:       os.Getenv("OWM_API_VER"),
		ApiEndpoint:  os.Getenv("OWM_API_ENDPOINT"),
		ApiKey:       os.Getenv("OWM_API_KEY"),
		City:         os.Getenv("OWM_CITY"),
		Units:        os.Getenv("OWM_UNITS"),
	}
}

func (e EnvVars) buildFQDN() string {
	fqdn := fmt.Sprintf("%s/%s/%s/%s?q=%s&units=%s&appid=%s",
		e.BaseURL,
		e.ApiPre,
		e.ApiVer,
		e.ApiEndpoint,
		e.City,
		e.Units,
		e.ApiKey)
	return fqdn
}

// TODO: add more weather data and location
// TODO: read from template html file
// TODO: accept input from form
// TODO: tailwindcss
// TODO: get unsplash/cc photo of location
// TODO: use geolocation from browser
// TODO: cache weather data (boltdb?)
// TODO: migrate from make to task
func main() {
	envVars := getEnvVars()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Weather</title>
				<script src="https://unpkg.com/htmx.org"></script>
			</head>
			<body>
				<h1>Today's Weather</h1>
				<div hx-get="/weather" hx-trigger="load" hx-swap="outerHTML">
					<p>Loading...</p>
				</div>
			</body>
			</html>
		`)
	})

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(envVars.buildFQDN())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data WeatherData
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var units string
		switch envVars.Units {
		case "metric":
			units = "°C"
		case "imperial":
			units = "°F"
		default:
			units = "K"
		}

		if len(data.Weather) > 0 {
			fmt.Fprintf(w, "<p>Description: %s</p><p>Temperature: %.2f %s</p>", data.Weather[0].Description, data.Main.Temp, units)
		} else {
			http.Error(w, "Weather data is not available", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
