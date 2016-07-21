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

	if page.Title != "TESTING HSET TITLE" || page.Url != "TESTING HSET URL" || page.Pageinfo.Onlooker != 35844 || page.Pageinfo.Comments != 1454 || page.Pageinfo.Banana != 2 {
		t.Errorf("%s\n%s\n%d\n%d\n%d\n", page.Title, page.Url, page.Pageinfo.Onlooker, page.Pageinfo.Comments, page.Pageinfo.Banana)
		t.Error("Page information is not correct!")
	}

}

func TestHdel(t *testing.T) {

	n, err := Hdel("ac2905884")

	if err != nil && n != 0 {
		t.Errorf("Testing HDEL error --> %s", err.Error())
	}

}
