package main

import (
	"fmt"
)

//并行
//缓冲量为1的信道
var lockx = make(chan int, 1)

//确定此时只有goroutine对map进行修改
func safeRun(f func()) {
	<-lockx
	f()
	lockx <- 1
}

//不重复
var checked map[string]bool = make(map[string]bool)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, quit chan int) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	defer func() {
		quit <- 1
	}()

	//设置depth为了避免查找死循环
	//深度为0或者该url已经被查找过
	if depth <= 0 || checked[url] {
		return
	}
	body, urls, err := fetcher.Fetch(url)

	//将这次查找的url设置已经查找过的状态
	safeRun(func() {
		checked[url] = true
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	sub_quit := make(chan int, len(urls))
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, sub_quit)
	}
	for i := 0; i < len(urls); i++ {
		<-sub_quit
	}
	return
}

func main() {
	lockx <- 1
	//channel
	quit := make(chan int, 1)
	Crawl("https://golang.org/", 4, fetcher, quit)
	fmt.Println(<-quit)
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
