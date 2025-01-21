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
	Short: "Adds one or more habits to the database.",
	Long: `The add command accepts a list of strings as arguments, representing the names of the habits to be stored in the database. Each habit name provided will be added as a new entry.

Examples:
  # Add multiple habits: daily-walk and work-out
  little-habits add run daily-walk work-out

  # Add a single habit: meditation
  little-habits add meditation

Options:
  (No options available for this command.)

Usage:
  little-habits add <habit-name> [<habit-name> ...]

Important Notes:
  Ensure the habit names are unique to avoid duplicate entries in the database.
  If a habit with the same name already exists, it will not be added again.`,
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
