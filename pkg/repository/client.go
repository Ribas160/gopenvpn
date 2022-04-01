package repository

type Config struct {
	Name string
}

func NewClient(cfg Config) *Config {
	return &Config{
		Name: cfg.Name,
	}
}
