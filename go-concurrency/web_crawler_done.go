// 修改后的
package main

import (
	"fmt"
	"sync"
)

type safeUrlCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

// checkAndAdd 检查缓存，有就返回 true，没有的话先加入缓存（设置为 true 做标记），再返回 false
func (suc *safeUrlCache) checkAndAdd(url string) bool {
	suc.mu.Lock()
	defer suc.mu.Unlock()
	if _, ok := suc.cache[url]; ok {
		return true
	}
	suc.cache[url] = true
	return false
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	var suc = &safeUrlCache{cache: make(map[string]bool)}

	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		// This implementation doesn't do either:
		// 如果没有这句代码的话，会无限打印
		if depth < 0 {
			return
		}

		// Don't fetch the same URL twice.
		if suc.checkAndAdd(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			// Fetch URLs in parallel.
			go crawl(u, depth-1)
		}
	}

	wg.Add(1)
	crawl(url, depth)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func ShowWebCrawlerDone() {
	Crawl("https://golang.org/", 4, fetcher)
}
