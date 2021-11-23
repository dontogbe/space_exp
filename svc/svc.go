package svc

import (
	"fmt"
)

type svc struct {
}

func NewSvc() Svc {
	return &svc{}
}
func (s *svc) GetLocation(x, y, z, vel float64) (location float64, err error) {
	if vel < 0 {
		return 0, fmt.Errorf("velocity cannot be negative")
	}
	location = x*SectorID + y*SectorID + z*SectorID + vel
	return location, nil
}
