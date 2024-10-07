package storage

import "github.com/1Sjoerd/event-processor-hycare/internal/processors"

// ProcessorRepository defines the interface for processor storage
type ProcessorRepository interface {
	GetAll() (processors.Processor, error)
	GetProcessor(name string) (processors.Processor, error)
}
