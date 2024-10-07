package repository

import (
	"embed"
	"fmt"

	"github.com/1Sjoerd/event-processor-hycare/internal/processors"
	"github.com/lovoo/goka"
	"gopkg.in/yaml.v3"
)

//go:embed processors/*.yaml
var folder embed.FS

type LocalRepository struct {
	basePath string
	tm       goka.TopicManager
}

func NewLocalRepository(basePath string, tm goka.TopicManager) *LocalRepository {
	return &LocalRepository{basePath: basePath, tm: tm}
}

func (r *LocalRepository) GetProcessor(name string) (*processors.Processor, error) {

	folder.ReadFile(name + ".yaml")
	// TODO: Implement logic to load processor from JSON
	data, err := folder.ReadFile(name + ".yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading processor config file: %w", err)
	}

	var config processors.ProcessorConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	processor := processors.NewProcessor(r.tm, config)

	return &processor, nil
}
