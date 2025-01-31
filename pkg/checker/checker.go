package checker

import (
	"site-checker/pkg/config"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type SiteChecker struct {
	client *http.Client
}

type Result struct {
	Name    string
	URL     string
	Error   error
	Elapsed time.Duration
}

func NewSiteChecker(timeout time.Duration) *SiteChecker {
	return &SiteChecker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *SiteChecker) Check(url string) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при запросе: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("неожиданный статус: %d", resp.StatusCode)
	}

	return nil
}

func (c *SiteChecker) CheckConcurrent(sites []config.URL) []Result {
	results := make([]Result, 0, len(sites))
	resultsChan := make(chan Result, len(sites))
	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)
		go func(site config.URL) {
			defer wg.Done()
			start := time.Now()
			err := c.Check(site.URL)
			resultsChan <- Result{
				Name:    site.Name,
				URL:     site.URL,
				Error:   err,
				Elapsed: time.Since(start),
			}
		}(site)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}
