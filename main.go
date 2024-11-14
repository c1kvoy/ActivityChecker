package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ResponseData struct {
	SiteName string
	RTime    time.Duration
	RStatus  int
}

var urls = []string{
	"https://www.google.com",
	"https://www.youtube.com",
	"https://www.facebook.com",
	"https://www.amazon.com",
	"https://www.twitter.com",
	"https://www.github.com",
}

func CheckingRequest(url string, wg *sync.WaitGroup, dataChan chan []ResponseData) {
	defer wg.Done()
	responseArray := []ResponseData{}
	for i := 0; i < 10; i++ {
		StartTime := time.Now()
		r, _ := http.Get(url)
		rd := ResponseData{
			SiteName: url,
			RTime:    time.Since(StartTime),
			RStatus:  r.StatusCode,
		}
		responseArray = append(responseArray, rd)
	}
	dataChan <- responseArray
}

func checkerGet() [][]ResponseData {
	var c = make(chan []ResponseData)
	var wg sync.WaitGroup
	var ResponsesData [][]ResponseData
	for _, url := range urls {
		wg.Add(1)
		go CheckingRequest(url, &wg, c)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for rd := range c {
		ResponsesData = append(ResponsesData, rd)
	}
	return ResponsesData
}

func main() {
	data := checkerGet()
	for _, siteData := range data {
		for _, result := range siteData {
			fmt.Printf("Site: %s, Status: %d, Response Time: %v\n", result.SiteName, result.RStatus, result.RTime)
		}
	}
}
