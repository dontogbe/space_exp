package svc

const (
	SectorID = 1
)

type Svc interface {
	GetLocation(x, y, z, vel float64) (location float64, err error)
}
