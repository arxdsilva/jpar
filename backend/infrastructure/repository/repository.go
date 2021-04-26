package repository

import (
	"github.com/arxdsilva/jpar/backend/domain"
)

type Service interface {
	GetPorts() ([]domain.Port, error)
	UpsertPort(domain.Port) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GetPorts() (dp []domain.Port, err error) {
	return
}

func (s *service) UpsertPort(domain.Port) (err error) {

	return
}
