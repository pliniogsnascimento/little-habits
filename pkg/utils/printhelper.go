package utils

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
)

func PrintHabitsWeekProgress(habits []habit.Habit) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	weekdays := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	header := []string{" "}
	header = append(header, GetHabitsNames(habits)...)
	fmt.Fprintln(w, strings.Join(header, "\t"))

	for _, day := range weekdays {
		fmt.Fprintf(w, "%s\t", day)
		for range habits {
			fmt.Fprintf(w, "\t")
		}
		fmt.Fprintf(w, "\n")
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
