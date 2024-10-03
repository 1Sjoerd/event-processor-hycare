package repository

// Processor defines an interface for processor logic
type Processor interface {
	ProcessEvent(data map[string]interface{}) error
	GetInputTopic() string
	GetOutputTopic() string
}

// ProcessorRepository defines the interface for processor storage
type ProcessorRepository interface {
	GetProcessor(name string) (Processor, error)
}
