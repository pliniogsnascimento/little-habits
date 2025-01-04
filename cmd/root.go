/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/pliniogsnascimento/little-habits/pkg/db"
	"github.com/pliniogsnascimento/little-habits/pkg/habit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	service    habit.HabitService
	cfgFile    string
	logger     *zap.SugaredLogger
	dbConnOpts *db.DbConnOpts
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "little-habits",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$PWD/.little-habits.yaml", "config file (default is $HOME/.little-habits.yaml)")

	viper.SetConfigFile(".little-habits.yaml")
	viper.AddConfigPath("$PWD")
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("db", &dbConnOpts)
	if err != nil {
		panic(err)
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()

	logger.Debugln("db configs:", dbConnOpts)

	gormDb, err := db.NewGormDb(dbConnOpts, logger)
	if err != nil {
		panic(err)
	}

	service = db.NewHabitRepo(gormDb, logger)
}
