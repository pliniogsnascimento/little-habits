/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

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

func getHabitsNames(habits []habit.Habit) (names []string) {
	names = []string{}

	for _, value := range habits {
		names = append(names, value.Name)
	}
	return
}
