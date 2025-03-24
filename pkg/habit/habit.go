package habit

import (
	"time"
)

type Habit struct {
	ID        uint
	Name      string      `gorm:"unique"`
	Plan      []HabitPlan `gorm:"foreignKey:HabitID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HabitPlan struct {
	Day      time.Time `gorm:"type:date;primaryKey"`
	Executed bool
	HabitID  uint `gorm:"primaryKey"`
}

// type HabitTemplate struct {
// 	ID    uint
// 	habit Habit
//
// 	weekRecurrency []time.Weekday
// }

const (
	week  = iota
	month = iota
)

type Recurrency struct {
	weekday time.Weekday
	month   time.Month
}

func NewHabit(name string) Habit {
	return Habit{Name: name}
}

func (h Habit) GetStats() float64 {
	executed := 0.0
	for _, v := range h.Plan {
		if v.Executed {
			executed++
		}
	}

	return (executed / float64(len(h.Plan))) * 100
}

// HabitService it's the service interface to operate habits.
type HabitService interface {
	CreateHabit(habit []*Habit) ([]*Habit, error)

	ListHabits() (*[]Habit, error)

	// GetMonthProgress is a function to get progress of all habits in the month.
	GetHabitsByPlanInTimeRange(init, end time.Time) (*[]Habit, error)
	GetHabitProgress(habitName string, month time.Month) (*Habit, error)
	AddRecord(habitName string, plan HabitPlan) error
	DeleteRecord(habitName string, date time.Time) error
}
