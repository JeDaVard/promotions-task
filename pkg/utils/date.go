package utils

import (
	"fmt"
	"time"
)

func FormatCsvDate(dateString string) (time.Time, error) {
	// Defining the layout of our csv date string, a map can also be used if this can vary from csv to csv
	layout := "2006-01-02 15:04:05 -0700 MST"

	// Parse the date string using the specified layout
	t, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}, err
	}
	return t, nil
}
