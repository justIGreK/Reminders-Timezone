package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	UTC_lat  = 0.0
	UTC_long = 0.0
)

type TimezoneResponse struct {
	UTCtime  string `json:"converted_time"`
	UserTime string `json:"original_time"`
}

type TimeDiff struct{}

func NewTimeDiff() *TimeDiff{
	return &TimeDiff{}
}

const (
	DateTimeFormat string = "2006-01-02 15:04:05"
)

func (t *TimeDiff) GetTimeDiff(lat, lon float64) (int, error) {
	var tzResponse TimezoneResponse
	url := fmt.Sprintf("https://api.ipgeolocation.io/timezone/convert?apiKey=%s&lat_from=%f&long_from=%f&lat_to=%f&long_to=%f",
		os.Getenv("TIMEZONE_API"), lat, lon, UTC_lat, UTC_long)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error making request: %v", err)
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response: %v", err)
		return 0, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("bad response from server: %s", string(body))
		return 0, fmt.Errorf("bad response from server: %s", string(body))
	}

	if err := json.Unmarshal(body, &tzResponse); err != nil {
		log.Printf("error unmarshalling response: %v", err)
		return 0, fmt.Errorf("error unmarshalling response: %v", err)
	}
	parsedUTC, err := time.Parse(DateTimeFormat, tzResponse.UTCtime)
	if err != nil {
		log.Println(err)
	}
	parsedOriginal, err := time.Parse(DateTimeFormat, tzResponse.UserTime)
	if err != nil {
		log.Println(err)
	}
	log.Println(parsedOriginal, parsedUTC)
	diff := parsedOriginal.Sub(parsedUTC)
	diffhour := int(diff.Hours())
	return diffhour, nil
}
