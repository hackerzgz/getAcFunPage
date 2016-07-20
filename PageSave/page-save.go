package PageSave

import (
	"github.com/garyburd/redigo/redis"
)

/*
 * HMset use HMSET to save PageInfo.
 * @param pageId string
 * @param title
 * @param url
 * @param onLooker int
 * @param comments
 * @param banana
 * @return size_length, error
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

/*
 * Hdel use HDEL to delete a PageInfo.
 * @param pageId string
 * @return hdel_size, err
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
