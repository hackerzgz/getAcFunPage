package PageSave

import (
	"github.com/garyburd/redigo/redis"

	PageInfo "getAcFunPage/PageInfo"
	utils "getAcFunPage/utils"
)

// Author: HackerZ
// Time  : 2016/7/19 14:27

/* IndexItem Item on "The Hottest Today".
 * Usage:
 * url:		 			"http://www.acfun.tv/a/ac2907308"
 * title:	 			"JUST A STRING"
 * dataid:				"2907308"
 * Pageinfo.PageID:   	"2907308"
 * Pageinfo.OnLooker: 	25781
 * Pageinfo.Comments: 	524
 * Pageinfo.Banana: 	2
 */
type IndexItem struct {
	Url      string
	Title    string
	Dataid   string
	Pageinfo PageInfo.PageInfo
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
func Keys(keys_arg string) ([]string, error) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	pageId_list, err := redis.Strings(rc.Do("KEYS", keys_arg))
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
			indexItem.Title = v
		case "url":
			indexItem.Url = v
		case "onLooker":
			indexItem.Pageinfo.Onlooker = utils.StrToInt64(v)
		case "comments":
			indexItem.Pageinfo.Comments = utils.StrToInt64(v)
		case "banana":
			indexItem.Pageinfo.Banana = utils.StrToInt64(v)
		default:
		}
	}
	indexItem.Pageinfo.PageID = utils.AcIdToPageId(pageId)
	indexItem.Dataid = utils.AcIdToPageId(pageId)
	return indexItem, nil
}
