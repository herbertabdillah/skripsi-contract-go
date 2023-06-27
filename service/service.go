package service

import (
	"errors"

	"github.com/herbertabdillah/skripsi-contract-new/repository"
	"github.com/hyperledger-labs/cckit/router"
)

type Service struct {
	context    router.Context
	repository repository.Repository
}

func NewService(c router.Context, r repository.Repository) Service {
	return Service{context: c, repository: r}
}

func (s Service) Updatable(year int, semester string) error {
	appConfig, err := s.repository.GetApplicationConfig()
	if err != nil {
		return err
	}

	if appConfig.Year == year && appConfig.Semester == semester {
		return nil
	} else {
		return errors.New("Data can't updated")
	}
}
