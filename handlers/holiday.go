package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	QUERY_1 = "http://www.webcal.fi/cal.php?id=125&format=json&start_year="
	QUERY_2 = "&end_year="
	QUERY_3 = "&tz=America%2FToronto"
)

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func Holidays() string {
	cur_year := time.Now().Year()
	query := fmt.Sprintf("%s%v%s%v%s", QUERY_1, cur_year, QUERY_2, cur_year+1, QUERY_3)
	resp, err := http.Get(query)

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
