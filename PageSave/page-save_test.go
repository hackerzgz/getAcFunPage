package PageSave

import (
	"fmt"
	"testing"
)

func TestHset(t *testing.T) {
	fmt.Println("=== Testing HSET Start ===")
	n, err := HMset("ac2905884", "TESTING HSET TITLE", "TESTING HSET URL", 35844, 1454, 2)

	if err != nil || n == "-1" {
		t.Errorf("error --> %s --> %s", err.Error())
	}

	fmt.Println("=== Testing HSET Done ===")
}
