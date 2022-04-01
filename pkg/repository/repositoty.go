package repository

type ClientConfig interface {
	Create() (string, error)
}

type Files interface {
	Copy(from string, to string) error
	Read(file string) (string, error)
}

type Repository struct {
	ClientConfig
	Files
}

func NewRepository(c *Config) *Repository {
	return &Repository{
		ClientConfig: newClientConfigRepository(c),
		Files:        newFilesRepository(),
	}
}
