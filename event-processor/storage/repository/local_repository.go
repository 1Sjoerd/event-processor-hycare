package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/1Sjoerd/event-processor-hycare/processors"
)

type LocalRepository struct {
	processorMap map[string][]*processors.ProcessorConfig
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
		processorMap: make(map[string][]*processors.ProcessorConfig),
	}
}

func (r *LocalRepository) LoadProcessors() error {
	processorDir := "storage/repository/processors"

	files, err := os.ReadDir(processorDir)
	if err != nil {
		return fmt.Errorf("error reading processor directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filePath := filepath.Join(processorDir, file.Name())
			log.Printf("Loading processor from file: %s", filePath)

			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("error reading file %s: %v", file.Name(), err)
			}

			var processor processors.ProcessorConfig
			err = yaml.Unmarshal(data, &processor)
			if err != nil {
				log.Fatalf("error unmarshalling YAML for %s: %v", file.Name(), err)
			}

			if processor.ID == "device_info_processor" {
				r.processorMap["device_info_processor"] = []*processors.ProcessorConfig{&processor}
			}
			for _, id := range processor.HycareItemIds {
				r.processorMap[id] = append(r.processorMap[id], &processor)
			}
		} else {
			log.Printf("Skipping non-YAML file: %s", file.Name())
		}
	}

	if _, found := r.processorMap["device_info_processor"]; !found {
		log.Fatalf("device_info_processor was not found in the loaded files")
	}

	return nil
}

func (r *LocalRepository) GetProcessorMap() map[string][]*processors.ProcessorConfig {
	return r.processorMap
}
