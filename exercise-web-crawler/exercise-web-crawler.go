// 参考 https://www.exmedia.jp/blog/a-tour-of-go%e3%81%ae%e7%b7%b4%e7%bf%92%e5%95%8f%e9%a1%8c%e3%82%92%e8%a7%a3%e8%aa%ac%e3%81%99%e3%82%8b%e3%82%b7%e3%83%aa%e3%83%bc%e3%82%ba11-11-exercise-web-crawler/
// https://qiita.com/ruiu/items/dba58f7b03a9a2ffad65

package main

import (
	"fmt"
	"time"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
	IsFetched(url string) bool
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if fetcher.IsFetched(url) || depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	// インクリメント前にgoroutineを起動しなければスケジューリングがうまくいかない。
	wg.Add(len(urls)) // カウンタ：インクリメント
	for _, u := range urls {
		go func(u string) { 
			Crawl(u, depth-1, fetcher)
			defer wg.Done() // カウンタ：デクリメント
		}(u)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait() // カウンタ：0になるまで待つ
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(2 * time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher {
	"https://golang.org/": &fakeResult {
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult {
		"Packages",
		[]string {
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult {
		"Package fmt",
		[]string {
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult {
		"Package os",
		[]string {
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

var (
	fetchedURL = make(map[string]bool)
	mu sync.Mutex
	wg sync.WaitGroup
)

func (f fakeFetcher) IsFetched(url string) bool {
	mu.Lock() // 排他制御開始
	_, ok := fetchedURL[url]
	if !ok {
		fetchedURL[url] = true
	}
	mu.Unlock() // 排他制御終了
	return ok
}