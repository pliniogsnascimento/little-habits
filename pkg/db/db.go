package db

import (
	"fmt"
	"time"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: Use this parameters on NewGormDb func
type DbConnOpts struct {
	User                     string
	Password                 string
	Host                     string
	Port                     string
	Database                 string
	DefaultMaxConns          int32
	DefaultMinConns          int32
	DefaultMaxConnLifetime   time.Duration
	DefaultMaxConnIdleTime   time.Duration
	DefaultHealthCheckPeriod time.Duration
	DefaultConnectTimeout    time.Duration
}

func NewPostgresGormDb(opts *DbConnOpts, logger *zap.SugaredLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		opts.Host,
		opts.User,
		opts.Password,
		opts.Database,
		opts.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	panicIfErr(err, logger)

	err = db.AutoMigrate(&habit.Habit{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&habit.HabitPlan{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewSQLiteGormDb(filename string, logger *zap.SugaredLogger) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	panicIfErr(err, logger)

	err = db.AutoMigrate(&habit.Habit{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&habit.HabitPlan{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func panicIfErr(err error, logger *zap.SugaredLogger) {
	if err != nil {
		logger.Errorln(err)
		panic(err)
	}
}
