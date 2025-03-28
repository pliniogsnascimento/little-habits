/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/pliniogsnascimento/little-habits/pkg/db"
	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"github.com/pliniogsnascimento/little-habits/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	service    habit.HabitService
	cfgFile    string
	logger     *zap.SugaredLogger
	dbConnOpts *db.DbConnOpts
	pHelper    utils.PrinterHelper
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "little-habits",
	Short: "The little-habits command-line interface (CLI) is a tool to manage and track habits effectively.",
	Long:  `The little-habits command-line interface (CLI) is a tool to manage and track habits effectively. It provides subcommands to add habits, record their progress, and view reports.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func() { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$PWD/.little-habits.yaml", "config file (default is $HOME/.little-habits.yaml)")
	var err error
	var gormDb *gorm.DB

	// TODO: Not loading config from home dir
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("$PWD")
	viper.AddConfigPath("$HOME/.little-habits")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// var logCfg zap.Config

	// TODO: understand this and make it work
	// err = viper.UnmarshalKey("log", &logCfg)
	// if err != nil {
	// 	panic(err)
	// }
	// zapLogger := zap.Must(logCfg.Build())

	switch viper.Get("mode") {
	case "development":
		zapLogger, _ := zap.NewDevelopment()
		defer zapLogger.Sync()
		logger = zapLogger.Sugar()
		dbPath := path.Join(os.Getenv("PWD"), "data.db")
		gormDb, err = db.NewSQLiteGormDb(dbPath, logger)
	case "server":
		zapLogger, _ := zap.NewProduction()
		defer zapLogger.Sync()
		logger = zapLogger.Sugar()
		err = viper.UnmarshalKey("db", &dbConnOpts)
		if err != nil {
			panic(err)
		}
		gormDb, err = db.NewPostgresGormDb(dbConnOpts, logger)
	default:
		zapLogger, _ := zap.NewProduction()
		defer zapLogger.Sync()
		logger = zapLogger.Sugar()
		dbPath := path.Join(os.Getenv("HOME"), ".little-habits", "data.db")
		gormDb, err = db.NewSQLiteGormDb(dbPath, logger)
	}

	if err != nil {
		panic(err)
	}

	service = db.NewHabitRepo(gormDb, logger)
	pHelper = utils.NewPrinterHelper(logger)
}
