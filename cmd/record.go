/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"github.com/spf13/cobra"
)

var (
	executed bool
	dates    []string
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Records a habit as planned or executed for specific dates.",
	Long: `The record command is used to log a habit, either as planned or executed.

Examples:
    # Creates a planned record for daily-walk habit in dates 2020-01-01 and 2020-01-02
    little-habits record daily-walk --executed --dates 2020-01-01 --dates 2020-01-02

    # Creates a planned record for gym-session habit today
    little-habits record gym-session

    # Creates an executed record for meditation habit
    little-habits record meditation -e

Options:
    -e, --executed=false:
      Specify whether the habit was executed. by default, this flag is false, indicating the habit was planned but not executed. If set to true, the habit is marked as executed.
    -d, --dates:
      Accepts an array of strings representing dates in the format YYYY-MM-DD. These dates are converted into time.Time objects. If no date is provided, the current day is used.

Usage:
    little-habits record <habit-name> [options]

Important:
  Dates must be formatted correctly (YYYY-MM-DD). Providing an invalid format will result in a parse error.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("input is no valid")
		}

		var recordDate time.Time
		loc := time.Now().Location()

		if len(dates) == 0 {
			recordDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)
			dates = append(dates, recordDate.Format("2006-01-02"))
		}

		for _, d := range dates {
			providedDate, err := time.Parse("2006-01-02", d)
			if err != nil {
				fmt.Println(err)
				continue
			}
			recordDate = time.Date(providedDate.Year(), providedDate.Month(), providedDate.Day(), 0, 0, 0, 0, loc)

			plan := habit.HabitPlan{Day: recordDate, Executed: executed}
			err = service.AddRecord(args[0], plan)
			if err != nil {
				fmt.Println(err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	recordCmd.Flags().BoolVarP(&executed, "executed", "e", false, "Specify whether the habit was executed.")
	recordCmd.Flags().StringArrayVarP(&dates, "dates", "d", []string{}, "Accepts an array of strings representing dates in the format YYYY-MM-DD. These dates are converted into time.Time objects. If no date is provided, the current day is used.")
}
