package PageSave

import (
	"github.com/garyburd/redigo/redis"

	PageInfo "getAcFunPage/PageInfo"
	utils "getAcFunPage/utils"
)

/* IndexItem Item on "The Hottest Today".
 * Usage:
 * url:		 			"http://www.acfun.tv/a/ac2907308"
 * title:	 			"JUST A STRING"
 * dataid:				"2907308"
 * pageinfo.PageID:   	"2907308"
 * pageinfo.OnLooker: 	25781
 * pageinfo.Comments: 	524
 * pageinfo.Banana: 	2
 */
type IndexItem struct {
	url      string
	title    string
	dataid   string
	pageinfo PageInfo.PageInfo
}

/* HMset use HMSET to save PageInfo.
 * @param 	pageId string
 * @param 	title
 * @param 	url
 * @param 	onLooker int
 * @param 	comments
 * @param 	banana
 * @return 	size_length, error
 */
func HMset(pageId, title, url string, onLooker, comments, banana int64) (string, error) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	// HMSET Value of the Key.
	hms_size, err := redis.String(rc.Do("HMSET", pageId, "title", title, "url", url, "onLooker", onLooker, "comments", comments, "banana", banana))

	// EXPIRE set TTL of the Key.(30 min)
	rc.Do("EXPIRE", pageId, 30*60)

	if err != nil {
		return "-1", err
	}
	return hms_size, nil
}

/* Hdel use HDEL to delete a PageInfo.
 * @param 	pageId string
 * @return 	hdel_size, err
 */
func Hdel(pageId string) (int64, error) {
	sub_key := []string{"title", "url", "onLooker", "comments", "banana"}

	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	var hdel_size int64 = -1
	var err error
	// HDEL Value of the Key.
	for _, skey := range sub_key {
		hdel_size, err = redis.Int64(rc.Do("HDEL", pageId, skey))
		if err != nil {
			return -1, err
		}
	}
	return hdel_size, nil
}

/* Keys use KEYS to find acfun pageId
 * @return 	pageId_index, err
 */
func Keys() ([]string, error) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	pageId_list, err := redis.Strings(rc.Do("KEYS", "ac29*"))
	if err != nil {
		return []string{""}, err
	}
	return pageId_list, nil
}

/* Hgetall use HGETALL to find the Hash Key.
 * @param 	pageId
 * @return 	IndexItem, err
 */
func Hgetall(pageId string) (IndexItem, error) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	pageId_index, err := redis.StringMap(rc.Do("HGETALL", pageId))
	if err != nil {
		return IndexItem{}, err
	}

	var indexItem IndexItem
	for k, v := range pageId_index {
		switch k {
		case "title":
			indexItem.title = v
		case "url":
			indexItem.url = v
		case "onLooker":
			indexItem.pageinfo.Onlooker = utils.StrToInt64(v)
		case "comments":
			indexItem.pageinfo.Comments = utils.StrToInt64(v)
		case "banana":
			indexItem.pageinfo.Banana = utils.StrToInt64(v)
		default:
		}
	}
	indexItem.pageinfo.PageID = utils.AcIdToPageId(pageId)
	indexItem.dataid = utils.AcIdToPageId(pageId)
	return indexItem, nil
}
