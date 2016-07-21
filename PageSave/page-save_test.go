package PageSave

import (
	"testing"
)

func TestHset(t *testing.T) {

	n, err := HMset("ac2905884", "TESTING HSET TITLE", "TESTING HSET URL", 35844, 1454, 2)

	if err != nil && n == "-1" {
		t.Errorf("Testing HSET error --> %s", err.Error())
	}

}

func TestKeys(t *testing.T) {

	pageId_index, err := Keys("ac2905884")
	if err != nil || len(pageId_index) != 1 {
		t.Errorf("Testing KEYS error --> %s", err.Error())
	}

}

func TestHgetall(t *testing.T) {

	page, err := Hgetall("ac2905884")
	if err != nil {
		t.Errorf("Testing HGETALL error --> %s", err.Error())
	}

	if page.title != "TESTING HSET TITLE" || page.url != "TESTING HSET URL" || page.pageinfo.Onlooker != 35844 || page.pageinfo.Comments != 1454 || page.pageinfo.Banana != 2 {
		t.Errorf("%s\n%s\n%d\n%d\n%d\n", page.title, page.url, page.pageinfo.Onlooker, page.pageinfo.Comments, page.pageinfo.Banana)
		t.Error("Page information is not correct!")
	}

}

func TestHdel(t *testing.T) {

	n, err := Hdel("ac2905884")

	if err != nil && n != 0 {
		t.Errorf("Testing HDEL error --> %s", err.Error())
	}

}
