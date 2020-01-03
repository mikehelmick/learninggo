package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Seattle's location.
const lat = 47.6062
const lon = -122.3321

// SunData represents sunrise-sunset API data.
// See https://sunrise-sunset.org/api
type SunData struct {
	Sunrise                   time.Time `json:"sunrise"`
	Sunset                    time.Time `json:"sunset"`
	SolarNoon                 time.Time `json:"solar_noon"`
	DayLength                 int64     `json:"day_length"`
	CivalTwilightBegin        time.Time `json:"civil_twilight_begin"`
	CivalTwilightEnd          time.Time `json:"civil_twilight_end"`
	NauticalTwilightBegin     time.Time `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       time.Time `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin time.Time `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   time.Time `json:"astronomical_twilight_end"`
}

// day length is returned in seconds.
func (s *SunData) calcDayLength() (hours, minutes int64) {
	hours = s.DayLength / 3600
	minutes = (s.DayLength - hours*3600) / 60
	return
}

func (s *SunData) toString() string {
	hours, minutes := s.calcDayLength()
	return fmt.Sprintf("   Sunrise: %s\nSolar Noon: %s\n    Sunset: %s\nDay Length: %02d:%02d\n",
		s.Sunrise.Local().Format(time.UnixDate),
		s.SolarNoon.Local().Format(time.UnixDate),
		s.Sunset.Local().Format(time.UnixDate),
		hours, minutes)

}

// Response represents the outer resonse field for the API
type Response struct {
	Results SunData `json:"results"`
	Status  string  `json:"status"`
}

// Reply is used on the done signaling channel and bundles errors and responses.
type Reply struct {
	Response *Response
	Error    *error
}

func callSunriseAPI(lat, lon float32, ch chan Reply) {
	resp, err := http.Get(fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%f&lng=%f&date=today&formatted=0", lat, lon))

	if err != nil {
		ch <- Reply{nil, &err}
		return
	}
	defer resp.Body.Close()

	apiResp := &Response{}
	json.NewDecoder(resp.Body).Decode(apiResp)
	ch <- Reply{apiResp, nil}
}

func main() {
	done := make(chan Reply)
	defer close(done)
	log.Printf("Calling Sunrise API")
	go callSunriseAPI(lat, lon, done)
	log.Printf("Called API waiting for response")

	// Wait for a response from the API call
	if r := <-done; r.Error != nil {
		log.Fatalf("Error calling the API: %v", r.Error)
	} else {
		log.Printf("Got Response")
		log.Printf("Sunrise-Sunset data for seattle %v", r.Response)
		data, _ := json.Marshal(r.Response)
		log.Printf("JSON: %s", string(data))

		fmt.Printf(r.Response.Results.toString())
	}
}
