package netspeed

type IProvider interface {
	GetSpeed() (download float64, upload float64, err error)
	Name() string
}
