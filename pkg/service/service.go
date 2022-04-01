package service

import (
	"github.com/Ribas160/gopenvpn/pkg/repository"
)

type ClientConfig interface {
	Create() (string, error)
}

type Files interface {
	Copy(from string, to string) error
}

type Service struct {
	ClientConfig
	Files
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		ClientConfig: newClientConfigService(repos.ClientConfig),
		Files:        newFilesService(repos.Files),
	}
}
