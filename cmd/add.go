/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: executeAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func executeAdd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no name for new habit provided")
	}
	habits := []*habit.Habit{}

	for _, value := range args {
		newHabit := habit.NewHabit(value)
		habits = append(habits, &newHabit)
	}

	_, err := service.CreateHabit(habits)
	if err != nil {
		return err
	}

	fmt.Printf("%s habit(s) created!", strings.Join(args, ", "))

	return nil
}

// TODO: remove this and use helper
func printHabits(habits []habit.Habit) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "Name")
	for _, value := range habits {
		fmt.Fprintln(w, value.Name)
	}
}

// TODO: remove this and use helper
func printHabitsWeekProgress(habits []habit.Habit) {
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
	header = append(header, getHabitsNames(habits)...)
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

func getHabitsNames(habits []habit.Habit) (names []string) {
	names = []string{}

	for _, value := range habits {
		names = append(names, value.Name)
	}
	return
}
