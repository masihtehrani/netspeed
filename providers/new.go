package providers

import (
	"strings"

	"github.com/masihtehrani/netspeed"
	"github.com/masihtehrani/netspeed/providers/fastcom"
	"github.com/masihtehrani/netspeed/providers/speedtest"
)

type Provider struct {
	fastcom   netspeed.IProvider
	speedtest netspeed.IProvider
}

func New() (*Provider, error) {
	fastcom, err := fastcom.New()
	if err != nil {
		return nil, err
	}
	speedtest, err := speedtest.New()
	if err != nil {
		return nil, err
	}

	return &Provider{
		fastcom:   fastcom,
		speedtest: speedtest,
	}, nil
}

func (p Provider) Builder(name string) netspeed.IProvider {
	switch strings.ToUpper(name) {
	case p.fastcom.Name():
		return p.fastcom
	case p.speedtest.Name():
		return p.speedtest
	default:
		return p.speedtest
	}
}
