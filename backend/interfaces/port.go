package interfaces

import (
	"github.com/arxdsilva/jpar/backend/domain"
	"github.com/arxdsilva/jpar/backend/infrastructure/repository"
)

type portService struct {
	repo repository.Service
}

func NewPortService() domain.PortService {
	return &portService{}
}
