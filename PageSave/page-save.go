package PageSave

import (
	// "errors"

	"github.com/garyburd/redigo/redis"
)

/*
 * HMset use HMSET to save PageInfo.
 * @param pageId
 * @param title
 * @param url
 * @param onLooker
 * @param comments
 * @param banana
 * @return size_length, error
 */
func HMset(pageId, title, url string, onLooker, comments, banana int) (string, error) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	size_length, err := redis.String(rc.Do("HMSET", pageId, "title", title, "url", url, "onLooker", onLooker, "comments", comments, "banana", banana))

	if err != nil {
		return "-1", err
	}
	return size_length, nil
}
