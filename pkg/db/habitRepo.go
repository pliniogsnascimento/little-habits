package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HabitRepo struct {
	gormDb *gorm.DB
	logger *zap.SugaredLogger
}

func NewHabitRepo(gormDb *gorm.DB, logger *zap.SugaredLogger) *HabitRepo {
	return &HabitRepo{gormDb: gormDb, logger: logger}
}

func (db *HabitRepo) closeConnFromPool(conn *pgxpool.Conn) {
	conn.Release()
}

func (h *HabitRepo) CreateHabit(habits []*habit.Habit) ([]*habit.Habit, error) {
	hDtos := []*HabitDTO{}

	for _, habit := range habits {
		hDtos = append(hDtos, NewHabitDTO(habit))
	}

	if err := h.gormDb.Create(hDtos).Error; err != nil {
		return nil, err
	}

	return habits, nil
}

func (h *HabitRepo) ListHabits() (*[]habit.Habit, error) {
	dtoList := HabitDTOList{}

	result := h.gormDb.Find(&dtoList)
	h.logger.Debugln(result.RowsAffected)
	h.logger.Debugln(dtoList)

	return dtoList.toEntity(), nil
}

// GetMonthProgress is a function to get progess of all habits in the month.
func (h *HabitRepo) GetMonthProgress(month time.Month) (*[]habit.Habit, error) {
	return nil, nil
}

func (h *HabitRepo) GetHabitProgress(habitName string, month time.Month) (*habit.Habit, error) {
	return nil, nil
}

func (h *HabitRepo) AddRecord(habit habit.Habit, day time.Time) error {
	return nil
}
