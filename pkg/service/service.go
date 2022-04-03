package service

import (
	"github.com/Ribas160/gopenvpn/pkg/repository"
)

type Client interface {
	CreateConfig() (string, error)
	BuildNew() error
}

type Files interface {
	Copy(from string, to string) error
	Read(file string) (string, error)
}

type Service struct {
	Client
	Files
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Client: newClientConfigService(repos.Client),
		Files:  newFilesService(repos.Files),
	}
}
