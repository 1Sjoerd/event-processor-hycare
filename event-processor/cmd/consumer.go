package cmd

import (
	"log"

	"github.com/1Sjoerd/event-processor-hycare/processors"

	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the event processor",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting event processor to consume iot-stream...")
		processors.StartEventProcessor()
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
