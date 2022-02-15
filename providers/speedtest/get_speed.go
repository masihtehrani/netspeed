package speedtest

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func (s *speedtestProvider) GetSpeed() (download float64, upload float64, err error) {
	serverList, err := speedtest.FetchServerList(s.user)
	if err != nil {
		return 0, 0, err
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return 0, 0, err
	}

	for _, s := range targets {
		_ = s.PingTest()
		_ = s.DownloadTest(false)
		_ = s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		download = s.DLSpeed
		upload = s.ULSpeed
	}

	return download, upload, nil
}
