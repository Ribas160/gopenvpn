package service

import (
	"github.com/Ribas160/gopenvpn/pkg/repository"
)

type ClientService struct {
	repo repository.Client
}

func newClientConfigService(repo repository.Client) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateConfig() (string, error) {
	return s.repo.CreateConfig()
}

func (s *ClientService) BuildNew() error {
	return s.repo.BuildNew()
}
