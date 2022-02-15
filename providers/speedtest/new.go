package speedtest

import (
	"github.com/masihtehrani/netspeed"
	"github.com/showwin/speedtest-go/speedtest"
)

type speedtestProvider struct {
	user *speedtest.User
	name string
}

const NAME = "SPEEDTEST"

func New() (netspeed.IProvider, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	return &speedtestProvider{
		user: user,
		name: NAME,
	}, nil
}

func (s speedtestProvider) Name() string {
	return s.name
}
