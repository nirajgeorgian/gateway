package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigFile :- default config file
var ConfigFile string

// Verbose :- show default output error
var Verbose bool

// UseViper :- use viper for configuration
var UseViper bool

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "c", "", "config file (default is config.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output (default is false)")
	RootCmd.PersistentFlags().BoolVarP(&UseViper, "viper", "r", true, "Use Viper for configuration (default is true)")
}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/job")
		viper.AddConfigPath("$HOME/.job")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found: %v\n", err)
		} else {
			// Config file was found but another error was produced
			fmt.Errorf("%s", err)
			fmt.Println()
		}
		os.Exit(1)
	}
}

// RootCmd :- root command to be exported
var RootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "gateway service application",
	Long:  "gateway service is responsible for serving as a frontend for microservice API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute :- execute startup command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
