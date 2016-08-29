package PageSave

import (
	"fmt"
	"testing"

	PageInfo "getAcFunPage/PageInfo"
)

func TestPage2JSON(t *testing.T) {
	indexItem := IndexItem{
		Title:  "Testing2JSON",
		Url:    "http://www.testing2JSON.com/a/ac2907308",
		Dataid: "2907308",
		Pageinfo: PageInfo.PageInfo{
			Onlooker: 25781,
			Comments: 524,
			Banana:   2,
		},
	}

	result, err := Page2JSON(indexItem, "../ac_page")
	if err != nil {
		t.Errorf("Page2JSON error --> %s", err.Error())
	}

	if result != "{\"page\":{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON.com/a/ac2907308\",\"title\":\"Testing2JSON\",\"ac-id\":\"2907308\",\"onlooker\":25781,\"comments\":524,\"banana\":2}}" {
		t.Error("Page2JSON not Complete --> ", result)
	}
}

func TestPages2JSON(t *testing.T) {
	indexItem1 := IndexItem{
		Title:  "Testing2JSON1",
		Url:    "http://www.testing2JSON1.com/a/ac2907308",
		Dataid: "2907308",
		Pageinfo: PageInfo.PageInfo{
			PageID:   "ac2907308",
			Onlooker: 25781,
			Comments: 524,
			Banana:   2,
		},
	}

	indexItem2 := IndexItem{
		Title:  "Testing2JSON2",
		Url:    "http://www.testing2JSON2.com/a/ac2907308",
		Dataid: "2907308",
		Pageinfo: PageInfo.PageInfo{
			PageID:   "ac2907308",
			Onlooker: 25781,
			Comments: 524,
			Banana:   2,
		},
	}

	indexs := []IndexItem{indexItem1, indexItem2}

	fmt.Println("index --> ", indexs)

	result, err := Page2JSON(indexs, "../ac_pages")

	if err != nil {
		t.Errorf("Pages2JSON error --> %s", err.Error())
	}

	if result != "{\"pages\":[{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON1.com/a/ac2907308\",\"title\":\"Testing2JSON1\",\"ac-id\":\"ac2907308\",\"onlooker\":25781,\"comments\":524,\"banana\":2},{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON2.com/a/ac2907308\",\"title\":\"Testing2JSON2\",\"ac-id\":\"ac2907308\",\"onlooker\":25781,\"comments\":524,\"banana\":2}]}" {
		t.Error("Pages2JSON not Complete --> ", result)
	}
}
