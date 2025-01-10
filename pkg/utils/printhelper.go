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

func (h *PrinterHelper) PrintHabitsWeekProgress(habits []habit.Habit) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	header := []string{" "}
	header = append(header, GetHabitsNames(habits)...)
	fmt.Fprintln(w, strings.Join(header, "\t"))

	dates := getWeekDates(time.Now())

	for _, day := range dates {
		fmt.Fprintf(w, "%d %s\t", day.Day(), day.Weekday())
		for _, v := range habits {
			for _, planned := range v.Plan {
				h.logger.Debugln("comparing", day, "and", planned.Day, ":", planned.Day.Compare(day))
				if planned.Day.Compare(day) == 0 {
					if planned.Executed {
						fmt.Fprintf(w, "%s\t", "x")
					} else {
						fmt.Fprintf(w, "%s\t", "o")
					}
				} else {
					fmt.Fprintf(w, "%s\t", " ")
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

func getWeekDates(filter time.Time) []time.Time {
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
