package db

import (
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
)

type HabitDTO struct {
	ID        uint
	Name      string          `gorm:"unique"`
	Plan      *[]HabitPlanDTO `gorm:"foreignKey:HabitID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HabitPlanDTO struct {
	ID       uint
	Day      time.Time
	Executed bool
	HabitID  uint
}

type HabitDTOList []HabitDTO

func NewHabitDTO(habit *habit.Habit) *HabitDTO {
	return &HabitDTO{
		Name: habit.Name,
	}
}

func NewHabitDTOList(habits *[]habit.Habit) *[]HabitDTO {
	habitsDto := []HabitDTO{}

	for _, value := range *habits {
		habitsDto = append(habitsDto, *NewHabitDTO(&value))
	}

	return &habitsDto
}

func (h HabitDTOList) toEntity() *[]habit.Habit {
	habits := []habit.Habit{}

	for _, value := range h {
		habits = append(habits, value.toEntity())
	}

	return &habits
}

func (h HabitDTO) toEntity() habit.Habit {
	plan := []habit.HabitPlan{}

	if h.Plan != nil {
		for _, value := range *h.Plan {
			plan = append(plan, habit.HabitPlan{Day: value.Day, Executed: value.Executed})
		}
	}

	return habit.Habit{
		Name: h.Name,
		Plan: &plan,
	}
}
