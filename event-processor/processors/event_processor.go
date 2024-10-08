package processors

import (
	"encoding/json"
	"log"

	"github.com/robertkrimen/otto"
)

type ProcessorConfig struct {
	ID            string   `yaml:"id"`
	Description   string   `yaml:"description"`
	Enabled       bool     `yaml:"enabled"`
	Input         string   `yaml:"input"`
	HycareItemIds []string `yaml:"hycareItemIds"`
	Script        string   `yaml:"script"`
}

var processorMap map[string][]*ProcessorConfig

func InitProcessorMap(procMap map[string][]*ProcessorConfig) {
	processorMap = procMap
}

func ProcessEvent(key string, value string, currentData map[string]interface{}) string {
	var eventMap map[string]interface{}
	err := json.Unmarshal([]byte(value), &eventMap)
	if err != nil {
		log.Fatalf("Error parsing event JSON: %v", err)
	}

	deviceInfo, ok := eventMap["device_info"].(map[string]interface{})
	if !ok {
		log.Fatalf("Error extracting device_info from event: expected map but got %T", eventMap["device_info"])
	}

	deviceInfoProcessor, found := processorMap["device_info_processor"]
	if !found || len(deviceInfoProcessor) == 0 {
		log.Fatalf("device_info_processor is not loaded")
	}

	// Pas de aanroep van runProcessor aan om currentData mee te geven
	processedEvent := runProcessor(deviceInfoProcessor[0], eventMap, currentData)

	tags, ok := deviceInfo["tags"].(map[string]interface{})
	if !ok {
		return processedEvent
	}

	hycareItemID, ok := tags["hycareItemId"].(string)
	if !ok {
		return processedEvent
	}

	if processorsToRun, found := processorMap[hycareItemID]; found {
		for _, processor := range processorsToRun {
			runProcessor(processor, eventMap, currentData)
		}
	}

	return processedEvent
}

func runProcessor(processor *ProcessorConfig, eventMap map[string]interface{}, currentData map[string]interface{}) string {
	if processor == nil {
		log.Fatalf("Processor is nil")
	}

	vm := otto.New()

	_, err := vm.Run(processor.Script)
	if err != nil {
		log.Fatalf("Error executing script for processor %s: %v", processor.ID, err)
	}

	jsEvent, err := json.Marshal(eventMap)
	if err != nil {
		log.Fatalf("Error marshalling event to JSON: %v", err)
	}

	jsCurrentData, err := json.Marshal(currentData)
	if err != nil {
		log.Fatalf("Error marshalling currentData to JSON: %v", err)
	}

	result, err := vm.Call("processEvent", nil, string(jsEvent), string(jsCurrentData))
	if err != nil {
		log.Fatalf("Error executing processEvent for processor %s: %v", processor.ID, err)
	}

	processedEvent, err := result.ToString()
	if err != nil {
		log.Fatalf("Error converting result to string: %v", err)
	}

	log.Printf("Processor %s executed", processor.ID)

	return processedEvent
}
