package processors

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/robertkrimen/otto"
	"gopkg.in/yaml.v2"
)

type ProcessorConfig struct {
	ID          string `yaml:"id"`
	Description string `yaml:"description"`
	Enabled     bool   `yaml:"enabled"`
	Input       string `yaml:"input"`
	Script      string `yaml:"script"`
	RetryPolicy struct {
		Retries int    `yaml:"retries"`
		Backoff string `yaml:"backoff"`
	} `yaml:"retry_policy"`
}

func LoadProcessor(filePath string) (*ProcessorConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var config ProcessorConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	return &config, nil
}

func ProcessEvent(key string, value string) string {
	processor, err := LoadProcessor("storage/repository/processors/device_info_processor.yaml")
	if err != nil {
		log.Fatalf("Failed to load processor: %v", err)
	}

	vm := otto.New()

	_, err = vm.Run(processor.Script)
	if err != nil {
		log.Fatalf("Error executing script: %v", err)
	}

	// Parse het event als een generieke map om JSON te verwerken zonder struct
	var eventMap map[string]interface{}
	err = json.Unmarshal([]byte(value), &eventMap)
	if err != nil {
		log.Fatalf("Error parsing event JSON: %v", err)
	}

	// Marshall het event naar een JSON string om te gebruiken in het JavaScript
	jsEvent, err := json.Marshal(eventMap)
	if err != nil {
		log.Fatalf("Error marshalling event to JSON: %v", err)
	}

	result, err := vm.Call("processEvent", nil, string(jsEvent))
	if err != nil {
		log.Fatalf("Error executing processEvent: %v", err)
	}

	processedEvent, err := result.ToString()
	if err != nil {
		log.Fatalf("Error converting result to string: %v", err)
	}

	return processedEvent
}
