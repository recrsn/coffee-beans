package config

type Repository struct {
	id string
}

func (r Repository) Id() string {
	return r.id
}
