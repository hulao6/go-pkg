package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/net/httplib"
	"github.com/toolkits/pkg/runner"
	"github.com/toolkits/pkg/slice"
)

func main1() {
	a := []int64{1, 3, 5, 56, 67, 7}
	b := []int64{3, 5, 6, 7}
	fmt.Println(slice.SubInt64(a, b))

	runner.Init()
	fmt.Println(runner.Hostname)
	fmt.Println(runner.Cwd)

	fmt.Println(file.FilesUnder(runner.Cwd))
}

var transport http.RoundTripper = &http.Transport{
	MaxIdleConnsPerHost:   100,
	ResponseHeaderTimeout: time.Second * 10,
	IdleConnTimeout:       time.Second * 10,
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			callHTTP()
		}()
	}
	wg.Wait()
}

func callHTTP() {
	req := httplib.Get("http://localhost:3456/print").SetTimeout(time.Second * 5)
	str, err := req.String()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if len(str) == 0 {
		log.Println("httplib Get failed, response is empty")
		return
	}

	fmt.Println("response body:", str)
}
