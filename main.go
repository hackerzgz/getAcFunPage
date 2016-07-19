package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	PageInfo "getAcFunPage/PageInfo"
	PageSave "getAcFunPage/PageSave"
)

// Author: HackerZ
// Time  : 2016-7-1 15:09

/*
 * Get AcFun Page which is "The Hottest Today"
 */

var (
	acfunPageRoot = "http://www.acfun.tv/v/list110/index.htm"
	ptnIndexItem  = regexp.MustCompile(`<div class="item "><a href="(/a/ac[0-9]{7,})" target="_blank" data-aid="([0-9]{7,})" title="(.{10,35})" class="title">.{10,35}</a></div>`)
)

// IndexItem Item on "The Hottest Today"
type IndexItem struct {
	url      string
	title    string
	dataid   string
	pageinfo PageInfo.PageInfo
}

/*
 * MAIN FUNCTION
 * AUTHOR : HACKERZ
 */
func main() {
	// Get url Content.
	fmt.Println("=== Get Index... ===")
	raw, statusCode := Get(acfunPageRoot)
	fmt.Println("statusCode --> ", statusCode)
	if statusCode != "200 OK" {
		fmt.Println("err --> False to Get Web Content.Please Check out Your URL!")
		return
	}

	// Find IndexItem.
	index, _ := FindIndexItem(raw)

	fmt.Println("=== IndexItem Match Done ===")
	// for i, item := range index {
	// 	fmt.Printf("\nPageID ==> [%s] || [%d] ==> Title:【%s】 || Onlooker: [%d] || Comment: [%d] || Banana: [%d] || Link: [%s]\n", item.pageinfo.PageID, i, item.title, item.pageinfo.Onlooker, item.pageinfo.Comments, item.pageinfo.Banana, item.url)
	// }

	fmt.Println("=== Save Page Info 2 Redis ===")
	// Save Page Info to Redis.
	for _, item := range index {
		PageSave.HMset("ac"+item.pageinfo.PageID, item.title, item.url, item.pageinfo.Onlooker, item.pageinfo.Comments, item.pageinfo.Banana)
	}
}

/*
 * Get Web Content.
 * @param url string
 * @return content
 * @return statusCode
 */
func Get(url string) (content, statusCode string) {
	resp, err := http.Get(url)
	if err != nil {
		statusCode = "-1"
		fmt.Println("Get The \"Hottest\" Page Error --> ", err.Error())
		return
	}
	defer resp.Body.Close()

	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = "-2"
		fmt.Println("Get The \"Hottest\" Page Error --> ", err.Error())
		return
	}
	statusCode = resp.Status
	content = string(bys)
	return
}

/*
 * FindIndexItem Find IndexItem Which "The Hottest Today".
 * @param content string
 * @return index
 * @return err
 */
func FindIndexItem(content string) (index []IndexItem, err error) {
	matches := ptnIndexItem.FindAllStringSubmatch(content, -1)

	index = make([]IndexItem, len(matches))
	for i, item := range matches {
		pageInfo := PageInfo.GetPageInfo(item[1])
		index[i] = IndexItem{"http://www.acfun.tv" + item[1], item[3], item[2], pageInfo}
	}
	return
}
