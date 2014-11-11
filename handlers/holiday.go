package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	QUERY = "http://www.webcal.fi/cal.php?id=125&format=json&start_year=2014&end_year=2014&tz=America%2FToronto"
)

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func Holidays() string {
	resp, err := http.Get(QUERY)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var reply []Holiday
	decoder.Decode(&reply)

	holidays := ""
	count := 0
	for _, i := range reply {
		time_holiday, _ := time.Parse("2006-01-2-MST", i.Date+"-EDT")

		if time_holiday.After(time.Now()) {
			holidays += i.Date + " " + i.Name + "\n"
			count++
		}

		// only show the first 3 next holidays
		if count == 3 {
			break
		}
	}
	return holidays
}
