package PageSave

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Author: HackerZ
// Time  : 2016/7/18 19:44

var (
	RedisClient *redis.Pool
	RedisDB     = 0
)

// initRedis Connect RedisDB and Make it available.
func initRedis(host, password, dbNum string) *redis.Pool {

	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", host)
		if err != nil {
			return nil, err
		}
		// Handle Auth.
		if password != "" {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}
		// Select dbNum.
		if dbNum != "" {
			_, selecterr := c.Do("SELECT", dbNum)
			if selecterr != nil {
				c.Close()
				return nil, selecterr
			}
		}
		return
	}

	return &redis.Pool{
		MaxIdle:     64,
		IdleTimeout: 3 * time.Second,
		MaxActive:   99999, // max number of connections
		// TestOnBorrow: func(c redis.Conn, t time.Time) error {
		// 	_, err := c.Do("PING")

		// 	return err
		// },
		Dial: dialFunc,
	}
}

// GetRedisClient return an available Redis Client.
/* Usage:
 *      // GetRedisClient()
 *      // Get a Conn from Pool
 *      rc := RedisClient.Get()
 *      // Return Conn into Pool When you are Done.
 *      defer rc.Close()
 */
func init() {
	RedisClient = initRedis(":6379", "", "")
}
