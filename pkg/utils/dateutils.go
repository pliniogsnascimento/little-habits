package utils

import "time"

func GetWeekDates(filter time.Time) []time.Time {
	dates := []time.Time{}
	loc := time.Now().Location()
	first := time.Date(filter.Year(), filter.Month(), (filter.Day() - int(filter.Weekday())), 0, 0, 0, 0, loc)

	for i := 0; i < 7; i++ {
		dates = append(dates, first)
		first = first.Add(time.Duration(24) * time.Hour)
	}
	return dates
}

func GetMonthDates(mFilter time.Month, yFilter int) []time.Time {
	dates := []time.Time{}
	loc := time.Now().Location()
	first := time.Date(yFilter, mFilter, 1, 0, 0, 0, 0, loc)

	for month := first.Month(); month == first.Month(); first = first.AddDate(0, 0, 1) {
		dates = append(dates, first)
	}
	return dates
}
