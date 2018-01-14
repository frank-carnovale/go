package main

import (
	"fmt"
	"github.com/frank-carnovale/go/stringutil"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var threads int
var channel = make(chan string)
var urlmap = make(map[string]int)
var mux sync.Mutex

// use fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	defer func() { channel <- url }()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	mux.Lock()
	urlmap[url]++
	var newUrl = urlmap[url] == 1
	mux.Unlock()
	if newUrl {
		for _, u := range urls {
			threads++
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {

	threads = 1

	go Crawl("http://golang.org/", 4, fetcher)
	for threads > 0 {
		url := <-channel
		threads--
		fmt.Printf("%s fetched.  threads now %d\n", url, threads)
	}
	fmt.Printf("done\n")
	for k, v := range urlmap {
		fmt.Printf("url %s found %d times\n", k, v)
	}

	str1 := stringutil.Reverse("\xe2\x98\x8e they think it's all over ☃")
	str2 := stringutil.Reverse("☎ they think it's all over ☃")
	fmt.Printf("reversed: [%s]\n", str1)
	fmt.Printf("reversed: [%s]\n", str2)

}

// fakeFetcher is a Fetcher that returns canned results.
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},

	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
			"http://dbnet.com.au/pkg/",
		},
	},

	"http://dbnet.com.au/pkg/": &fakeResult{
		"dbNet Australia",
		[]string{
			"http://dbnet.com.au/pkg/",
		},
	},
}
