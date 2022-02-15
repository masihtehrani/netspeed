package fastcom

import (
	"github.com/ddo/go-fast"
	"github.com/masihtehrani/netspeed"
)

// fastcomProvider implement Netflix fast.com provider.
type fastcomProvider struct {
	fastCom *fast.Fast
	name    string
}

const NAME = "FASTCOM"

// New make an instance of fast.com.
func New() (netspeed.IProvider, error) {
	fastProvider := &fastcomProvider{
		fastCom: fast.New(),
		name:    NAME,
	}
	// init
	err := fastProvider.Init()
	if err != nil {
		return nil, err
	}

	return fastProvider, err
}

func (f fastcomProvider) Name() string {
	return f.name
}
