package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"

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
	wg            sync.WaitGroup
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

	// Wait a Second to prevent Jump over a Shorter time.
	savePageTime := time.NewTimer(time.Second)

	go func() {
		for {
			select {
			case <-savePageTime.C:
				wg.Add(1)
				// wg.Done() when GetPageAndSave() Done
				GetPageAndSave()
				json := GetPageAndJSON()
				SaveTodayJSON(json)
				savePageTime.Reset(30 * time.Minute)
			}
		}
	}()

	// Wait for GetPageAndSave() Done.
	wg.Wait()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleTodayGetResp)
	mux.HandleFunc("/data", HandleDataGetResp)
	// solve brower requset favicon problem.
	mux.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, req *http.Request) {})
	http.ListenAndServe(":9001", mux)

}

/* HandleTodayGetResp
 * Handle Today Get Response And Return JSON.
 */
func HandleTodayGetResp(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		
	} else {
		fmt.Printf("=== This is <%s> Request ===", req.RemoteAddr)
		io.WriteString(rw, GetTodayJSON())
	}
}

/* HandleDataGetResp
 * Handle Today Get Response And Return JSON.
 */
func HandleDataGetResp(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		
	} else {
		/*
		 * req.ParseForm() is import!
		 * http://it.taocms.org/01/6527.htm
		 */
		req.ParseForm()
		fmt.Printf("=== This is <%s> Request ===", req.RemoteAddr)
		dataValue := req.FormValue("data")
		io.WriteString(rw, dataValue)
		// io.WriteString(rw, GetTodayJSON())
	}
}


/* GetPageAndSave
 * Get AcFun Page And Save them into Redis.
 */
func GetPageAndSave() {
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

}

/* GetPageAndJSON
 * Get Acfun Page List from Redis
 * and trans them to JSON
 */
func GetPageAndJSON() (pageJSON string) {
	// Get Page Info from Redis.
	keys, err := PageSave.Keys("ac29*")
	if err != nil {
		panic(err)
	}

	// Get PageList base on keys.
	var pageList []PageSave.IndexItem = make([]PageSave.IndexItem, len(keys))
	for k, v := range keys {
		pifr, err := PageSave.Hgetall(v)
		if err != nil {
			panic(err)
		}
		pageList[k] = pifr
	}
	fmt.Println("=== Get PageList Done ===")

	// Make Pages trans to JSON.
	pageJSON, err = PageSave.Page2JSON(pageList, "./ac_pages")
	if err != nil {
		panic(err)
	}
	fmt.Println("=== JSON Trans Done ===")
	return
}

/* SaveTodayJSON Save JSON to Redis
 * @param	json
 */
func SaveTodayJSON(json string) {
	dataFormat := time.Now().Format("2006-01-02")
	err := PageSave.Set("ac_JSON-"+dataFormat, json)
	if err != nil {
		panic(err)
	}
	fmt.Println("=== Save ac_JSON Success ===")
	// Save Page Info Done, Make Response available.
	wg.Done()
}

/* SaveTodayJSON Save JSON to Redis
 * @return	json
 */
func GetTodayJSON() (json string) {
	dataFormat := time.Now().Format("2006-01-02")
	json, err := PageSave.Get("ac_JSON-"+dataFormat)
	if err != nil {
		panic(err)
	}

	fmt.Println("=== Return JSON Success ===")
	return
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
