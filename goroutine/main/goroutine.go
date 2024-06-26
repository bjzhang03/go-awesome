package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var tokens = make(chan struct{}, 100)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true;
				n++;
				go func(link string) {
					worklist <- crawl(link)
				}(link)
				//这里的link是给匿名函数传递变量的行为
				//link在作用域内
			}
		}
	}

	//for list := range worklist {
	//	for _, link := range list {
	//		if !seen[link] {
	//			seen[link] = true
	//			go func(link string) {
	//				worklist <- crawl(link)
	//			}(link)
	//		}
	//
	//	}
	//}
}
