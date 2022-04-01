package service

import "github.com/Ribas160/gopenvpn/pkg/repository"

type FilesService struct {
	repo repository.Files
}

func newFilesService(repo repository.Files) *FilesService {
	return &FilesService{
		repo: repo,
	}
}

func (s *FilesService) Copy(from string, to string) error {
	return s.repo.Copy(from, to)
}
