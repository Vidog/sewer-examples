package main

import (
	sw "github.com/Vidog/sewer"

	"net/http"
	"io/ioutil"
	"time"
	"fmt"
)

func check_status(dr DownloadResult) bool {
	return dr.status == 200
}

func check_body_len(dr DownloadResult) bool {
	return len(dr.body) > 0
}

type DownloadResult struct {
	status int
	body []byte
	timeSpent float64
}

func download(url string) DownloadResult {
	var resp *http.Response
	var err error
	var body []byte

	timeStart := time.Now()

	resp, err = http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return DownloadResult{
		status: resp.StatusCode,
		body: body,
		timeSpent: time.Since(timeStart).Seconds(),
	}
}

func main() {
	urls := []string{
		"http://goo.gl/",
		"http://google.com/wtf", // will be filtered because of 404 response status
		"http://google.com",
	}

	data := sw.MakeTupleChanFromSlice( urls ).Apply(download).FilterAll(check_status, check_body_len)

	for r := range data {
		dr := r.Value.(DownloadResult)
		fmt.Printf("-- Read status %#v in %#v ms\n", dr.status, dr.timeSpent)
	}
}