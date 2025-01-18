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

func (h *PrinterHelper) PrintHabitsProgressInRange(habits []habit.Habit, dates []time.Time) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	header, div := getHeaderAndDiv(GetHabitsNames(habits))
	fmt.Fprintln(w, header)

	for i, day := range dates {
		if day.Weekday() == 0 && i > 0 {
			fmt.Fprint(w, div)
		}
		fmt.Fprintf(w, "%d %s\t", day.Day(), day.Weekday())
		for _, v := range habits {
			h.logger.Debugln(v)
			fmt.Fprint(w, getHabitOut(0, v, day, h.logger))
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}

func GetHabitsNames(habits []habit.Habit) (names []string) {
	names = []string{}

	for _, value := range habits {
		names = append(names, value.Name)
	}
	return
}

func getHabitOut(i int, habit habit.Habit, day time.Time, logger *zap.SugaredLogger) string {
	var out string

	if len(habit.Plan) <= i {
		out = fmt.Sprintf("%s\t", " ")
		logger.Debugf("[%s]skipping %s", habit.Name, day)
		return out
	}

	if habit.Plan[i].Day.Compare(day) == 0 {
		if habit.Plan[0].Executed {
			out = fmt.Sprintf("%s\t", "x")
			logger.Debugf("[%s]adding %s to %s", habit.Name, "x", habit.Plan[i].Day)
			return out
		}
		out = fmt.Sprintf("%s\t", "o")
		logger.Debugf("[%s]adding %s to %s", habit.Name, "o", habit.Plan[i].Day)
		return out
	}

	i++
	return getHabitOut(i, habit, day, logger)
}

func getHeaderAndDiv(headerList []string) (string, string) {
	header, divider := []string{" "}, []string{}
	header = append(header, headerList...)
	headerText := strings.Join(header, "\t")

	for range header {
		divider = append(divider, strings.Repeat("-", 6))
	}

	div := strings.Join(divider, "\t")
	div += "\n"

	return headerText, div
}
