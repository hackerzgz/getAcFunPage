package PageSave

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestPageSave(t *testing.T) {
	fmt.Println("=== Testing Redis Start ===")

	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	if _, err := rc.Do("SET", "ACP_TEST1", "test"); err != nil {
		t.Error(err.Error())
	}

	if test_value, _ := redis.String(rc.Do("GET", "ACP_TEST1")); test_value != "test" {
		t.Error("GET KEY ACP_TEST1 VALUE WRONG!")
	}

	if _, err := rc.Do("DEL", "ACP_TEST1"); err != nil {
		t.Error(err.Error())
	}

	fmt.Println("=== Testing Redis Done ===")
}
