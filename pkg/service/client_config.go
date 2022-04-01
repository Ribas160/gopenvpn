package service

import (
	"github.com/Ribas160/gopenvpn/pkg/repository"
)

type ClientConfigService struct {
	repo repository.ClientConfig
}

func newClientConfigService(repo repository.ClientConfig) *ClientConfigService {
	return &ClientConfigService{repo: repo}
}

func (s *ClientConfigService) Create() (string, error) {
	return s.repo.Create()
}
