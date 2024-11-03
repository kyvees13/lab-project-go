package service

import (
	"errors"
	"lab-server/internal/repository"
)

type InfoService struct {
	repository *repository.InfoRepository
}

func NewInfoService(repository *repository.InfoRepository) *InfoService {
	return &InfoService{repository: repository}
}

func (i *InfoService) GetVersion() (string, error) {
	data, err := i.repository.GetVersion()
	if err != nil {
		return "", errors.New("cant get database version")
	}
	return *data, nil
}
