package repository

import (
	"github.com/arxdsilva/jpar/backend/domain"
	"github.com/arxdsilva/jpar/backend/infrastructure/config"
	"github.com/go-pg/pg/v10"
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
	dp = []domain.Port{}
	sql := `select id, name, city, province, country, timezone, code FROM ports`
	_, err = config.Get.DB.Query(&dp, sql)
	return
}

func (s *service) UpsertPort(dp domain.Port) (err error) {
	p := &domain.Port{}
	sql := `select * FROM ports WHERE id=?`
	_, err = config.Get.DB.QueryOne(p, sql, dp.ID)
	if err != nil && err != pg.ErrNoRows {
		return
	}
	if p.Name != "" {
		return updatePort(dp)
	}
	return createPort(dp)
}

func createPort(port domain.Port) (err error) {
	sql := `INSERT INTO ports (id, name, city, country, alias, regions, coordinates, province, timezone, unlocs, code) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	empty := struct{}{}
	_, err = config.Get.DB.QueryOne(&empty, sql,
		port.ID, port.Name, port.City, port.Country, pg.Array(port.Alias), pg.Array(port.Regions), pg.Array(port.Coordinates), port.Province, port.Timezone, pg.Array(port.Unlocs), port.Code)
	return
}

func updatePort(port domain.Port) (err error) {
	sql := `UPDATE ports 
	SET name=?, city=?, country=?, alias=?, regions=?, coordinates=?, province=?, timezone=?, unlocs=?, code=? 
	WHERE id=?`
	empty := struct{}{}
	_, err = config.Get.DB.QueryOne(&empty, sql,
		port.Name, port.City, port.Country, pg.Array(port.Alias), pg.Array(port.Regions), pg.Array(port.Coordinates), port.Province, port.Timezone, pg.Array(port.Unlocs), port.Code, port.ID)
	return
}
