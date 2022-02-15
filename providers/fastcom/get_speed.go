package fastcom

import (
	"fmt"
)

// GetSpeed concrete GetSpeed method of ISpeed contracts for get data from fast.com provider
// this method just provide download and upload always is 0. I can't find in this time a Golang
// package provider download and upload.
func (f *fastcomProvider) GetSpeed() (download float64, upload float64, err error) {
	// get urls
	urls, err := f.GetUrls()
	if err != nil {
		return 0, 0, err
	}

	// measure
	measureChan := make(chan float64)

	go func() {
		// var value, units string
		for measure := range measureChan {
			download = measure / 1000 // kbps / 1000 = mbps
		}

		fmt.Printf("\r%c[2K -> %f\n", 27, download)
	}()

	err = f.Measure(urls, measureChan)
	if err != nil {
		return 0, 0, err
	}

	return download, upload, err
}
