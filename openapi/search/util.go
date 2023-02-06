package search

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/swxctx/ghttp"
)

// http请求
func DoHttpRequest(url string, query interface{}) ([]byte, error) {
	req := ghttp.Request{
		Url:     url,
		Query:   query,
		Method:  "GET",
		Timeout: time.Duration(10) * time.Second,
	}

	req.AddHeader("Content-Type", "application/json")

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("DoHttpRequest 请求失败, status -> %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var (
		reader io.ReadCloser
	)
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	respBs, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return respBs, nil
}
