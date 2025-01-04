package habit

import "time"

type Habit struct {
	Name string
	Plan *[]HabitPlan
}

type HabitPlan struct {
	Day      time.Time
	Executed bool
}

func NewHabit(name string) Habit {
	return Habit{Name: name}
}

// HabitService it's the service interface to operate habits.
type HabitService interface {
	CreateHabit(habit []*Habit) ([]*Habit, error)

	ListHabits() (*[]Habit, error)

	// GetMonthProgress is a function to get progess of all habits in the month.
	GetMonthProgress(month time.Month) (*[]Habit, error)
	GetHabitProgress(habitName string, month time.Month) (*Habit, error)
	AddRecord(habit Habit, day time.Time) error
}
