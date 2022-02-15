package fastcom

type Fastcom interface {
	Init() error
	GetUrls() (urls []string, err error)
	Measure(urls []string, KbpsChan chan<- float64) (err error)
}

func (f fastcomProvider) Init() error {
	return f.fastCom.Init()
}

func (f fastcomProvider) GetUrls() (urls []string, err error) {
	return f.fastCom.GetUrls()
}

func (f fastcomProvider) Measure(urls []string, KbpsChan chan<- float64) (err error) {
	return f.fastCom.Measure(urls, KbpsChan)
}
