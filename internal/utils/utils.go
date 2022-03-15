package utils

import "time"

// Date in format YYYY-MM-DD = 2006-01-02
func DatesBetween(startDate, endDate, layout string, includeStart, includeEnd bool) ([]string, error) {

	datesBetween := []string{}

	start, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	if !includeStart {
		start = start.Add(time.Hour * 24)
	}

	for start.Before(end) {

		datesBetween = append(datesBetween, start.Format(layout))

		start = start.Add(time.Hour * 24)
	}

	if includeEnd {
		if start.Equal(end) {
			datesBetween = append(datesBetween, start.Format(layout))
		}
	}

	return datesBetween, nil
}
