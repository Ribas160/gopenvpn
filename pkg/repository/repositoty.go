package repository

type Client interface {
	CreateConfig() (string, error)
	BuildNew() error
}

type Files interface {
	Copy(from string, to string) error
	Read(file string) (string, error)
}

type Repository struct {
	Client
	Files
}

func NewRepository(c *Config) *Repository {
	return &Repository{
		Client: newClientRepository(c),
		Files:  newFilesRepository(),
	}
}
