package cmd

import (
	"fmt"
	"log"

	"github.com/1Sjoerd/event-processor-hycare/internal/processors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startProcessorCmd = &cobra.Command{
	Use:   "start-processor",
	Short: "Start the Kafka processor",
	Run: func(cmd *cobra.Command, args []string) {
		brokers := viper.GetString("kafka_brokers")
		if brokers == "" {
			log.Fatal("No Kafka brokers provided")
		}
		fmt.Printf("Starting processor with brokers: %s\n", brokers)

		// Hier kun je bijvoorbeeld processors.StartProcessor() aanroepen
		err := processors.StartProcessor()
		if err != nil {
			log.Fatalf("Failed to start processor: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startProcessorCmd)
}
