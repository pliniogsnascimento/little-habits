package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := h.gormDb.Create(habits).Error; err != nil {
		return nil, err
	}

	return habits, nil
}

func (h *HabitRepo) ListHabits() (*[]habit.Habit, error) {
	habitList := []habit.Habit{}

	result := h.gormDb.Find(&habitList)
	h.logger.Debugln(result.RowsAffected)
	h.logger.Debugln(habitList)

	return &habitList, nil
}

func (h *HabitRepo) GetHabitsByPlanInTimeRange(init, end time.Time) (*[]habit.Habit, error) {
	habitList := []habit.Habit{}
	habitPlanList := []habit.HabitPlan{}

	h.logger.Debugln(init, end)

	err := h.gormDb.
		Preload("Plan", "day BETWEEN ? AND ?", init, end).
		Find(&habitList).Error
	if err != nil {
		return nil, err
	}

	h.logger.Debugln(habitList, habitPlanList)

	return &habitList, nil
}

func (h *HabitRepo) GetHabitProgress(habitName string, month time.Month) (*habit.Habit, error) {
	var habit habit.Habit

	err := h.gormDb.Preload("Plan").Where("name = ?", habitName).First(&habit).Error
	if err != nil {
		return nil, err
	}

	return &habit, nil
}

// TODO: AddOrUpdateRecord
func (h *HabitRepo) AddRecord(habitName string, plan habit.HabitPlan) error {
	var existingHabit habit.Habit

	err := h.gormDb.Where("name = ?", habitName).First(&existingHabit).Error
	if err != nil {
		return err
	}

	err = h.gormDb.Debug().Clauses(clause.OnConflict{UpdateAll: true}).Model(&existingHabit).Association("Plan").Append(&plan)
	if err != nil {
		return err
	}

	var habitPlanList []habit.HabitPlan
	err = h.gormDb.Model(&existingHabit).Association("Plan").Find(&habitPlanList)
	if err != nil {
		return err
	}

	h.logger.Debugln(habitPlanList)

	return nil
}
