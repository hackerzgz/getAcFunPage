package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	PageInfo "getAcFunPage/PageInfo"
	PageSave "getAcFunPage/PageSave"

	"github.com/bmizerany/pat"
)

// Author: HackerZ
// Time  : 2016-7-1 15:09

/*
 * Get AcFun Page which is "The Hottest Today"
 */

var (
	acfunPageRoot = "http://www.acfun.tv/v/list110/index.htm"
	ptnIndexItem  = regexp.MustCompile(`<div class="item "><a href="(/a/ac[0-9]{7,})" target="_blank" data-aid="([0-9]{7,})" title="(.{10,35})" class="title">.{10,35}</a></div>`)
	DF            = `[0-9]{4,}-[0-9]{2,}-[0-9]{2,}`
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

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(HandleTodayGetResp))

	mux.Get("/data/:data", http.HandlerFunc(HandleDataGetResp))
	mux.Get("/favicon.ico", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))

	http.Handle("/", mux)
	srv := &http.Server{
		Addr:         ":9001",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := srv.ListenAndServe()
	// err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe --> ", err)
	}

}

/* HandleTodayGetResp
 * Handle Today Get Response And Return JSON.
 */
func HandleTodayGetResp(rw http.ResponseWriter, req *http.Request) {
	log.Printf("=== This is <%s> Request ===\n", req.RemoteAddr)
	io.WriteString(rw, GetTodayJSON())
}

/* HandleDataGetResp
 * Handle Today Get Response And Return JSON.
 */
func HandleDataGetResp(rw http.ResponseWriter, req *http.Request) {
	log.Printf("=== This is <%s> Request ===\n", req.RemoteAddr)
	output := "404"
	dataValue := req.URL.Query().Get(":data")
	boolDF, err := regexp.MatchString(DF, dataValue)
	if err != nil {
		log.Fatal("data format error -->", err)
		panic(err)
	}
	if boolDF {
		output = GetDataJSON(dataValue)
	}
	io.WriteString(rw, output)
}

/* GetPageAndSave
 * Get AcFun Page And Save them into Redis.
 */
func GetPageAndSave() {
	// Get url Content.
	log.Println("=== Get Index... ===")
	raw, statusCode := GetPageContent(acfunPageRoot)
	log.Println("statusCode --> ", statusCode)
	if statusCode != "200 OK" {
		log.Println("err --> False to Get Web Content.Please Check out Your URL!")
		return
	}

	// Find IndexItem.
	index, _ := FindIndexItem(raw)

	log.Println("=== IndexItem Match Done ===")

	// Save Page Info to Redis.
	for _, item := range index {
		PageSave.HMset("ac"+item.Pageinfo.PageID, item.Title, item.Url, item.Pageinfo.Onlooker, item.Pageinfo.Comments, item.Pageinfo.Banana)
	}

	log.Println("=== Save Page Info 2 Redis Done ===")

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
	log.Println("=== Get PageList Done ===")

	// Make Pages trans to JSON.
	pageJSON, err = PageSave.Page2JSON(pageList, "./ac_pages")
	if err != nil {
		panic(err)
	}
	log.Println("=== JSON Trans Done ===")
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
	log.Println("=== Save ac_JSON Success ===")
	// Save Page Info Done, Make Response available.
	wg.Done()
}

/* GetTodayJSON Get Today JSON from Redis
 * @return	json
 */
func GetTodayJSON() (json string) {
	dataFormat := time.Now().Format("2006-01-02")
	json, err := PageSave.Get("ac_JSON-" + dataFormat)
	if err != nil {
		panic(err)
	}

	log.Println("=== Return JSON Success ===")
	return
}

/* GetDataJSON Get Data JSON from Redis
 * @param	data
 * @return	JSON
 */
func GetDataJSON(data string) (json string) {
	json, err := PageSave.Get("ac_JSON-" + data)
	if err != nil {
		json = "Have No this Data PageList."
		log.Println("=== Return JSON Failed ===")
	}

	log.Println("=== Return JSON Success ===")
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
		log.Println("Get The \"Hottest\" Page Error --> ", err.Error())
		return
	}
	defer resp.Body.Close()

	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = "-2"
		log.Println("Get The \"Hottest\" Page Error --> ", err.Error())
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
