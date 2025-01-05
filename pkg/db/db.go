package db

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
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

func NewGormDb(opts *DbConnOpts, logger *zap.SugaredLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		opts.Host,
		opts.User,
		opts.Password,
		opts.Database,
		opts.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	panicIfErr(err, logger)

	// TODO uncomment code if is not needed
	// dbConn, _ := db.DB()
	// panicIfErr(err, logger)

	// dbConn.Query("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	// panicIfErr(err, logger)

	// TODO add migrations
	err = db.AutoMigrate(&HabitDTO{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&HabitPlanDTO{})
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
