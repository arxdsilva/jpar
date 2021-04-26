package interfaces

import (
	"github.com/arxdsilva/jpar/backend/domain"
	"github.com/arxdsilva/jpar/backend/infrastructure/repository"
)

type portService struct {
	repo repository.Service
}

func NewPortService(r repository.Service) domain.PortService {
	return &portService{r}
}

func (ps *portService) UpsertPort(p domain.Port) (err error) {
	return ps.repo.UpsertPort(p)
}

func (ps *portService) ListPorts() (dp []domain.Port, err error) {
	return ps.repo.GetPorts()
}
