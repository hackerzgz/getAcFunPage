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
		t.Error("Page2JSON error -->", err.Error())
	}

	if result != "{\"page\":{\"page-id\":\"2907308\",\"url\":\"http://www.testing2JSON.com/a/ac2907308\",\"title\":\"Testing2JSON\",\"onlooker\":25781,\"comments\":524,\"banana\":2}}" {
		t.Error("Page2JSON not Complete --> ", result)
	}
}
