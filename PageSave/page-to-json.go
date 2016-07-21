package PageSave

import (
	// "fmt"

	"github.com/emilsjolander/goson"
)

/* Page2JSON use goson to reflex IndexItem to JSON
 * @param 	indexItem
 * @return 	result, err
 */
func Page2JSON(indexItem IndexItem) (string, error) {
	itemJSON, err := goson.Render("../ac_page", goson.Args{"IndexItem": indexItem})

	if err != nil {
		return "", err
	}

	return string(itemJSON), nil
}

/* Pages2JSON use goson to reflex []IndexItem to JSON
 * @param 	itemList
 * @return 	result, err
 */
func Pages2JSON(indexList []IndexItem) (string, error) {
	itemJSON, err := goson.Render("../ac_pages", goson.Args{"IndexItem": indexList})

	if err != nil {
		return "", err
	}

	return string(itemJSON), nil
}
