package main

import (
	"flag"
	"fmt"
	"github.com/c3b2a7/ylgy/constant"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var config struct {
	token     string
	count     int
	threshold int
	version   bool
}

func main() {
	flag.StringVar(&config.token, "token", "", "request token")
	flag.IntVar(&config.count, "count", 1000, "number of requests")
	flag.IntVar(&config.threshold, "threshold", 1, "how many requests is a coroutine responsible for， 1 means one coroutine per request")
	flag.BoolVar(&config.version, "v", false, "show version information")
	flag.Parse()

	if config.version {
		fmt.Printf("%s, %s\n", constant.Version, constant.BuildTime)
		os.Exit(0)
	}

	if config.token == "" {
		flag.Usage()
		os.Exit(1)
	}

	start := time.Now().UnixMilli()
	cnt := do()
	end := time.Now().UnixMilli()
	diff := end - start
	s := diff / 1000
	ms := diff - s*1000
	log.Printf("total: %d, success: %d, took: %ds%dms", config.count, cnt, s, ms)
}

func do() (cnt int) {
	client := http.DefaultClient
	client.Timeout = 3 * time.Second
	request, _ := http.NewRequest("GET", "http://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=1&rank_time=547&rank_role=1&skin=1", nil)
	request.Header.Set("Host", "cat-match.easygame2021.com")
	request.Header.Set("Accept-Encoding", "gzip,compress,br,deflate")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/17/page-frame.html")
	request.Header.Set("t", config.token)
	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c26) NetType/4G Language/zh_CN")
	wg := sync.WaitGroup{}
	wg.Add(config.count)
	for i := 0; i < config.count; i += config.threshold {
		left, right := i, i+config.threshold
		go func() {
			for j := left; j < right && j < config.count; j++ {
				response, err := client.Do(request)
				if err == nil && response.StatusCode == 200 {
					cnt++
				}
				wg.Done()
			}
		}()
	}
	wg.Wait()
	return
}
