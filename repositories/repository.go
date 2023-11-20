package repositories

import (
	"fmt"
	"github.com/recrsn/coffee-beans/config"
)

type RepositoryType string

const (
	RepositoryTypeMaven   RepositoryType = "maven"
	RepositoryTypePypi    RepositoryType = "pypi"
	RepositoryTypeGeneric RepositoryType = "generic"
)

type Repository struct {
	base        string
	Id          string
	Name        string
	Description string
	Type        RepositoryType
	Root        string
}

func NewRepositoryFromConfig(base string, repository config.Repository) *Repository {
	return &Repository{
		base:        base,
		Id:          repository.Id,
		Name:        repository.Name,
		Description: repository.Description,
		Type:        RepositoryType(repository.Type),
		Root:        repository.Root,
	}
}

func (r *Repository) Path() string {
	return fmt.Sprintf("%s/%s/*artifact", r.base, r.Id)
}

func (r *Repository) StaticPath() string {
	return fmt.Sprintf("%s/%s", r.base, r.Id)
}
