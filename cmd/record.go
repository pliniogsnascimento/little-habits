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
	day      string
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("input is no valid")
		}

		var recordDate time.Time
		var err error

		loc := time.Now().Location()
		recordDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)
		if day != "" {
			providedDate, err := time.Parse("2006-01-02", day)
			if err != nil {
				return err
			}
			recordDate = time.Date(providedDate.Year(), providedDate.Month(), providedDate.Day(), 0, 0, 0, 0, loc)
		}

		plan := habit.HabitPlan{Day: recordDate, Executed: executed}
		err = service.AddRecord(args[0], plan)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	recordCmd.Flags().BoolVarP(&executed, "executed", "e", false, "Help message for toggle")
	recordCmd.Flags().StringVarP(&day, "date", "d", "", "Help message for toggle")
}
