package json

import (
	"encoding/json"

	"github.com/bytedance/sonic"
)

// PrettyJSON
// @param any:
// @return string:
func PrettyJSON(any interface{}) (string, error) {
	bytes, err := json.MarshalIndent(any, "", "    ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToJSON 转成 json 字符串
func ToJSON(obj interface{}) (string, error) {
	bytes, err := sonic.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
