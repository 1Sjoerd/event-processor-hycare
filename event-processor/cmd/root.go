package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "event-processor",
	Short: "Event Processor CLI for HyCare",
	Long:  `This CLI is used to manage the event processor for HyCare.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig initializes configuration using Viper
func initConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	// Lees het configuratiebestand in, anders log een foutmelding
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Can't read config: %v", err)
	}
}

// init is een speciale functie die wordt aangeroepen voor elke run
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is config.yaml)")
}
