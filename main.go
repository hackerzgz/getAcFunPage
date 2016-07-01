package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Author: HackerZ
// Time  : 2016-7-1 15:09

/*
 * Get AcFun Page which is "The Hottest Today"
 */

var (
	ptnIndexItem = regexp.MustCompile(`<div class="item "><a href="(/a/ac[0-9]{7,})" target="_blank" data-aid="([0-9]{7,})" title="(.{10,35})" class="title">.{10,35}</a></div>`)
)

// IndexItem Item on "The Hottest Today"
type IndexItem struct {
	url    string
	title  string
	dataid string
}

func main() {
	// Get url Content.
	fmt.Println("=== Get Index... ===")
	raw, statusCode := Get("http://www.acfun.tv/v/list110/index.htm")
	fmt.Println("statusCode --> ", statusCode)
	if statusCode != "200 OK" {
		fmt.Println("err --> False to Get Web Content.Please Check out Your URL!")
		return
	}

	// Find IndexItem.
	index, _ := FindIndexItem(raw)

	fmt.Println("=== IndexItem Match Done. ===")
	for i, item := range index {
		fmt.Printf("\n[%d] ==> Title 【%s】 link ==> [%s]\n", i, item.title, item.url)
	}
}

// Get Web Content.
func Get(url string) (content, statusCode string) {
	resp, err := http.Get(url)
	if err != nil {
		statusCode = "-1"
		fmt.Println("err --> ", err.Error())
		return
	}
	defer resp.Body.Close()

	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = "-2"
		fmt.Println("err --> ", err.Error())
		return
	}
	statusCode = resp.Status
	content = string(bys)
	return
}

// FindIndexItem Find IndexItem Which "The Hottest Today".
func FindIndexItem(content string) (index []IndexItem, err error) {
	matches := ptnIndexItem.FindAllStringSubmatch(content, -1)
	// fmt.Println("matches -->", matches)
	index = make([]IndexItem, len(matches))
	for i, item := range matches {
		index[i] = IndexItem{"http://www.acfun.tv" + item[1], item[3], item[2]}
	}
	return
}
