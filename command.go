package main

import (
	"fmt"
	"os"
	"round_robin_with_weight/config"

	"github.com/spf13/cobra"
)

var (
	configure string
)

// RootCmd :
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Show all command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(configure) != 0 {
			fmt.Println("Service is running ...")
			readConfig(configure)
		}
	},
}

func readConfig(conf string) {
	err := config.GlobalConfigurations.ReadFrom(configure)
	if err != nil {
		panic(err)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version is", config.GlobalConfigurations.Version)
	},
}

// Execute :
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringVarP(&configure, "configure", "c", "config.yaml", "Current configure")
}
