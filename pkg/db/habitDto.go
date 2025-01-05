package db

import (
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
)

// TODO change this to single entity
type HabitDTO struct {
	ID        uint
	Name      string         `gorm:"unique"`
	Plan      []HabitPlanDTO `gorm:"foreignKey:HabitID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HabitPlanDTO struct {
	Day      time.Time `gorm:"type:date;primaryKey"`
	Executed bool
	HabitID  uint `gorm:"primaryKey"`
}

type HabitDTOList []HabitDTO

func NewHabitDTO(habit *habit.Habit) *HabitDTO {
	return &HabitDTO{
		Name: habit.Name,
	}
}

func NewHabitPlanDTO(plan *habit.HabitPlan) HabitPlanDTO {
	return HabitPlanDTO{
		Executed: plan.Executed,
		Day:      plan.Day,
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
		for _, value := range h.Plan {
			plan = append(plan, habit.HabitPlan{Day: value.Day, Executed: value.Executed})
		}
	}

	return habit.Habit{
		Name: h.Name,
		Plan: &plan,
	}
}
