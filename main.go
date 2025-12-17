package main

import (
	"fmt"
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
	code, bs, err := httplib.Get("http://localhost:3456/print").SetTimeout(time.Second).BytesV2()
	fmt.Printf("code=%d, response body: %s, error: %v\n", code, string(bs), err)
}
