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

/* IndexItem Item on "The Hottest Today".
 * Usage:
 * Url:		 			"http://www.acfun.tv/a/ac2907308"
 * Title:	 			"JUST A STRING"
 * Dataid:				"2907308"
 * pageinfo.PageID:   	"2907308"
 * pageinfo.OnLooker: 	25781
 * pageinfo.Comments: 	524
 * pageinfo.Banana: 	2
 */
type IndexItem struct {
	Url      string
	Title    string
	Dataid   string
	Pageinfo PageInfo.PageInfo
}

/*
 * MAIN FUNCTION
 * AUTHOR : HACKERZ
 */
func main() {
	// Get url Content.
	fmt.Println("=== Get Index... ===")
	raw, statusCode := GetPageContent(acfunPageRoot)
	fmt.Println("statusCode --> ", statusCode)
	if statusCode != "200 OK" {
		fmt.Println("err --> False to Get Web Content.Please Check out Your URL!")
		return
	}

	// Find IndexItem.
	index, _ := FindIndexItem(raw)

	fmt.Println("=== IndexItem Match Done ===")

	// Save Page Info to Redis.
	for _, item := range index {
		PageSave.HMset("ac"+item.Pageinfo.PageID, item.Title, item.Url, item.Pageinfo.Onlooker, item.Pageinfo.Comments, item.Pageinfo.Banana)
	}

	fmt.Println("=== Save Page Info 2 Redis Done ===")

	// Get Page Info from Redis.
	keys, err := PageSave.Keys("ac*")
	if err != nil {
		panic(err)
	}
	fmt.Println("keys --> ", keys)

	// Get PageList base on keys.
	var pageList []PageSave.IndexItem = make([]PageSave.IndexItem, 9)
	for k, v := range keys {
		pifr, err := PageSave.Hgetall(v)
		if err != nil {
			panic(err)
		}
		pageList[k] = pifr
	}
	fmt.Println("pageList --> ", pageList)

	// Make Pages trans to JSON.
	pageJSON, err := PageSave.Page2JSON(pageList, "./ac_pages")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON --> ", pageJSON)
}

/* GetPageContent Get Acfun Page Content.
 * @param url string
 * @return content
 * @return statusCode
 */
func GetPageContent(url string) (content, statusCode string) {
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

/* FindIndexItem Find IndexItem Which "The Hottest Today".
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
