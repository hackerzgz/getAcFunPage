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
func initRedis(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		IdleTimeout: 60 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}

			// _, err = c.Do("SELECT", RedisDB)

			return c, err
		},
	}
}

// GetRedisClient return an available Redis Client.
/* Usage:
 *      GetRedisClient()
 *      // Get a Conn from Pool
 *      rc := RedisClient.Get()
 *      // Return Conn into Pool When you are Done.
 *      defer rc.Close()
 */
func GetRedisClient() {
	RedisClient = initRedis(":6379")
}
