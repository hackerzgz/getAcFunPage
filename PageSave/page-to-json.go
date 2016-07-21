package PageSave

import (
	// "fmt"

	"github.com/emilsjolander/goson"
)

func Page2JSON(indexItem IndexItem) (string, error) {
	itemJSON, err := goson.Render("../ac_page", goson.Args{"IndexItem": indexItem})

	if err != nil {
		return "", err
	}

	return string(itemJSON), nil
}
