/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/utils"
	"github.com/spf13/cobra"
)

var month uint

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Displays the progress of all habits for the current week or a specified month of the current year.",
	Long: `The progress command displays the progress of all habits. By default, it shows the progress for the current week. If the --month flag is provided, it displays the progress for the specified month of the current year.

Examples:
  # Show progress for the current week
  little-habits progress

  # Show progress for December
  little-habits progress --month=12

Options:
  --month=<int>:
    An optional flag to filter progress by a specific month. Accepts an integer from 1 (January) to 12 (December). Defaults to the current week if not provided.

Usage:
  little-habits progress [options]

Important Notes:
  The --month flag filters data for the specified month in the current year only.
  Ensure the value for the --month flag is within the range of 1 to 12; otherwise, an error will occur.`,
	RunE: printProgress,
}

func printProgress(cmd *cobra.Command, args []string) error {
	if month > 12 {
		return fmt.Errorf("%d month is invalid", month)
	}
	var dateList []time.Time

	if month == 0 {
		dateList = utils.GetWeekDates(time.Now())
	} else {
		dateList = utils.GetMonthDates(time.Month(month), time.Now().Year())
	}

	habits, err := service.GetHabitsByPlanInTimeRange(dateList[0], dateList[len(dateList)-1])
	if err != nil {
		return err
	}
	logger.Debugln(habits)

	pHelper.PrintHabitsProgressInRange(*habits, dateList)
	return nil
}

func init() {
	rootCmd.AddCommand(progressCmd)
	progressCmd.Flags().UintVarP(&month, "month", "m", 0, "An optional flag to filter progress by a specific month.")
}
