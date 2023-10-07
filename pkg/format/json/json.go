package json

import (
	"encoding/json"
)

// PrettyJSON
// @Description: PrettyJSON
// @author ght
// @date 2023-10-07 16:55:51
// @param any:
// @return string:
func PrettyJSON(any interface{}) string {
	bytes, _ := json.MarshalIndent(any, "", "    ")
	return string(bytes)
}
