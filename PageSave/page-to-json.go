package PageSave

import (
	// "fmt"

	"github.com/emilsjolander/goson"
)

// Author: HackerZ
// Time  : 2016/7/20 11:43

/* Page2JSON use goson to reflex IndexItem or []IndexItem to JSON
 * @param 	indexItem
 * @return 	result, err
 */
func Page2JSON(indexItem interface{}, args_file string) (string, error) {
	itemJSON, err := goson.Render(args_file, goson.Args{"IndexItem": indexItem})

	if err != nil {
		return "", err
	}

	return string(itemJSON), nil
}
