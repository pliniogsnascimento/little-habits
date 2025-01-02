package habit

import "time"

type Habit struct {
	Name string
	Plan []HabitPlan
}

type HabitPlan struct {
	Day      int
	Month    int
	Year     int
	Executed bool
}

// HabitService it's the service interface to operate habits.
type HabitService interface {
	CreateHabit(habit Habit) (Habit, error)

	// GetMonthProgress is a function to get progess of all habits in the month.
	GetMonthProgress(month time.Month) ([]Habit, error)
	GetHabitProgress(habitName string, month time.Month) (Habit, error)
}
