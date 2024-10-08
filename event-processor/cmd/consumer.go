package cmd

import (
	"log"

	"github.com/1Sjoerd/event-processor-hycare/processors"
	"github.com/1Sjoerd/event-processor-hycare/storage/repository"
	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the event processor",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting event processor...")

		repo := repository.NewLocalRepository()

		err := repo.LoadProcessors()
		if err != nil {
			log.Fatalf("Error loading processors: %v", err)
		}

		processors.InitProcessorMap(repo.GetProcessorMap())

		processors.StartEventProcessor()
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
