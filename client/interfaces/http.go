package interfaces

import "github.com/arxdsilva/jpar/client/domains"

type HTTP struct {
	service domains.PortService
}

func NewHTTP(s domains.PortService) *HTTP {
	return &HTTP{s}
}
