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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// progressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// progressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	progressCmd.Flags().UintVar(&month, "month", 0, "month")
}
