//My solution to Go Tour's web crawler exercise using a Mutex and Waitgroup

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type fetchState struct {
	mu      sync.Mutex
	fetched map[string]bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url
func Crawl(url string, fetcher Fetcher, f *fetchState) {
	f.mu.Lock()
	alreadyFetched := f.fetched[url]
	f.fetched[url] = true
	f.mu.Unlock()

	if alreadyFetched {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var done sync.WaitGroup
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		done.Add(1)
		u2 := u
		go func() {
			defer done.Done()
			Crawl(u2, fetcher, f)
		}()
	}

	done.Wait()
	return
}

func makeState() *fetchState {
	f := &fetchState{}
	f.fetched = make(map[string]bool)
	return f
}

func main() {
	Crawl("https://golang.org/", fetcher, makeState())
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
