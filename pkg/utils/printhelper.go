package utils

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"go.uber.org/zap"
)

type PrinterHelper struct {
	logger *zap.SugaredLogger
}

func NewPrinterHelper(logger *zap.SugaredLogger) PrinterHelper {
	return PrinterHelper{logger: logger}
}

func (h *PrinterHelper) PrintHabits(habits []habit.Habit) {
	for _, v := range habits {
		fmt.Println(v)
	}
}

// TODO: There is a bug here when skipping. It skips every time the date doesn't match. It should only skip once at the end of loop
func (h *PrinterHelper) PrintHabitsWeekProgress(habits []habit.Habit) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, '.', tabwriter.Debug)

	header := []string{" "}
	header = append(header, GetHabitsNames(habits)...)
	fmt.Fprintln(w, strings.Join(header, "\t"))

	dates := GetWeekDates(time.Now())

	for _, day := range dates {
		fmt.Fprintf(w, "%d %s\t", day.Day(), day.Weekday())
		for _, v := range habits {
			for _, planned := range v.Plan {
				h.logger.Debugf("[%s]comparing %s and %s: %d", v.Name, day, planned.Day, planned.Day.Compare(day))
				if planned.Day.Compare(day) == 0 {
					if planned.Executed {
						fmt.Fprintf(w, "%s\t", "x")
						h.logger.Debugf("[%s]adding %s to %s", v.Name, "x", planned.Day)
					} else {
						fmt.Fprintf(w, "%s\t", "o")
						h.logger.Debugf("[%s]adding %s to %s", v.Name, "o", planned.Day)
					}
				} else {
					fmt.Fprintf(w, "%s\t", " ")
					h.logger.Debugf("[%s]skipping %s", v.Name, day)
				}
			}
		}
		fmt.Fprintf(w, "\n")
	}

	w.Flush()
}

func (h *PrinterHelper) PrintHabitsMonthProgress(habits []habit.Habit) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	header := []string{" "}
	header = append(header, GetHabitsNames(habits)...)
	fmt.Fprintln(w, strings.Join(header, "\t"))

	dates := getMonthDates(time.Now().Month(), 2025)

	for _, day := range dates {
		fmt.Fprintf(w, "%d %s\t", day.Day(), day.Weekday())
		for _, v := range habits {
			h.logger.Debugln(v)
			fmt.Fprintf(w, "%s\t", "x")
		}
		if day.Weekday() == 6 {
			w.Flush()
			fmt.Fprintf(os.Stdout, "\n%s\n", strings.Repeat("-", 30))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
			fmt.Fprintln(w, strings.Join(header, "\t"))
		} else {
			fmt.Fprintf(w, "\n")
		}
	}

	w.Flush()
}

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

func getMonthDates(mFilter time.Month, yFilter int) []time.Time {
	dates := []time.Time{}
	loc := time.Now().Location()
	first := time.Date(yFilter, mFilter, 1, 0, 0, 0, 0, loc)

	for month := first.Month(); month == first.Month(); first = first.AddDate(0, 0, 1) {
		dates = append(dates, first)
	}
	return dates
}

func GetHabitsNames(habits []habit.Habit) (names []string) {
	names = []string{}

	for _, value := range habits {
		names = append(names, value.Name)
	}
	return
}
