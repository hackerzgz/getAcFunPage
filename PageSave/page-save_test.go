package PageSave

import (
	"fmt"
	"testing"
)

func TestHset(t *testing.T) {
	fmt.Println("=== Testing HSET Start ===")
	n, err := HMset("ac2905884", "TESTING HSET TITLE", "TESTING HSET URL", 35844, 1454, 2)

	if err != nil || n != "5" {
		t.Errorf("error --> %s || size_length --> %s", err.Error(), n)
	}

	fmt.Println("=== Testing HSET Done ===")
}
