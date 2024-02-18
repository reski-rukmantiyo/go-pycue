/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/reski-rukmantiyo/go-pycue/pkg/config"
	"github.com/reski-rukmantiyo/go-pycue/pkg/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "web",
	Short: "This is template for your web application",
	Long:  `This is template for your web application. Application is based on gin framework and uses gorm as ORM.`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := getParameters(cmd)

		if len(flags) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		if flags["config"] == "" {
			fmt.Println("No config file specified")
			os.Exit(1)
		}

		// Get config file from flag
		configFile := flags["config"]

		// Load environment first
		webConfig, err := config.LoadEnv(configFile)
		if err != nil {
			log.Errorf("Error loading config file: %v", err)
		}

		// function to load router
		router.LoadRouter(webConfig)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("config", "c", "config.env", "Config file for application (required)")
}
