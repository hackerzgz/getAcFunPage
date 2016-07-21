package PageSave

import (
	"testing"

	PageInfo "getAcFunPage/PageInfo"
)

func TestPageToJSON(t *testing.T) {
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

	result, err := Page2JSON(indexItem)
	if err != nil {
		t.Errorf("Page2JSON error --> %s", err.Error())
	}

	if result != "{\"page\":{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON.com/a/ac2907308\",\"title\":\"Testing2JSON\",\"onlooker\":25781,\"comments\":524,\"banana\":2}}" {
		t.Error("Page2JSON not Complete --> ", result)
	}
}

func TestPagesToJSON(t *testing.T) {
	indexItem1 := IndexItem{
		Title:  "Testing2JSON1",
		Url:    "http://www.testing2JSON1.com/a/ac2907308",
		Dataid: "2907308",
		Pageinfo: PageInfo.PageInfo{
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
			Onlooker: 25781,
			Comments: 524,
			Banana:   2,
		},
	}

	indexs := []IndexItem{indexItem1, indexItem2}

	result, err := Pages2JSON(indexs)

	if err != nil {
		t.Errorf("Pages2JSON error --> %s", err.Error())
	}

	if result != "{\"pages\":[{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON1.com/a/ac2907308\",\"title\":\"Testing2JSON1\",\"onlooker\":25781,\"comments\":524,\"banana\":2},{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON2.com/a/ac2907308\",\"title\":\"Testing2JSON2\",\"onlooker\":25781,\"comments\":524,\"banana\":2}]}" {
		t.Error("Pages2JSON not Complete --> ", result)
	}
}
