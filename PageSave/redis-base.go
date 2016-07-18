package PageSave

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Author: HackerZ
// Time  : 2016/7/18 19:44

type config struct {
	RedisDB: "getAcfunPage"
}


// initRedis Connect RedisDB and Make it available.
func initRedis(host string) *redis.Pool {  
    return &redis.Pool{  
        MaxIdle: 64,      
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
  
            _, err = c.Do("SELECT", config.RedisDb)  
  
            return c, err  
        },  
    }  
} 


func Zadd(pageId string, onLooker, comments, banana int64)(bool, error) {
	return client.Do("ZADD", )
}
