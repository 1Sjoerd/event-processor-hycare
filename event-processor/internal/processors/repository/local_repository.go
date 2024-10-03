package repository

type LocalRepository struct {
	basePath string
}

func NewLocalRepository(basePath string) *LocalRepository {
	return &LocalRepository{basePath: basePath}
}

func GetProcessorJSON() {
	// TODO: Implement logic to load processor from JSON
}
