package repositories

import "github.com/recrsn/coffee-beans/config"

type RepositoryManager struct {
	repositories map[string]*Repository
}

func NewRepositoryManager(base string, repositories []config.Repository) *RepositoryManager {
	manager := &RepositoryManager{
		repositories: make(map[string]*Repository),
	}

	for _, r := range repositories {
		manager.repositories[r.Id] = NewRepositoryFromConfig(base, r)
	}

	return manager
}

func (m *RepositoryManager) GetRepository(id string) (*Repository, bool) {
	repository, found := m.repositories[id]
	return repository, found
}

func (m *RepositoryManager) GetRepositories() []*Repository {
	repositories := make([]*Repository, len(m.repositories))

	i := 0
	for _, r := range m.repositories {
		repositories[i] = r
		i++
	}

	return repositories
}
