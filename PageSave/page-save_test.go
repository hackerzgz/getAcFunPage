package PageSave

import (
	"fmt"
	"testing"
)

func TestHset(t *testing.T) {
	fmt.Println("=== Testing HSET Start ===")
	n, err := HMset("ac2905884", "TESTING HSET TITLE", "TESTING HSET URL", 35844, 1454, 2)

	if err != nil && n == "-1" {
		t.Errorf("error --> %s", err.Error())
	}

	fmt.Println("=== Testing HSET Done ===")
}

func TestHdel(t *testing.T) {
	fmt.Println("=== Testing HDEL Start ===")
	n, err := Hdel("ac2905884")

	if err != nil && n != 0 {
		t.Errorf("error --> %s", err.Error())
	}

	fmt.Println("=== Testing HDEL Done ===")
}
