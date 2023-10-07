package json

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/format/json"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// TestPrettyJSON
// @Description: TestPrettyJSON
// @author ght
// @date 2023-10-07 17:15:59
// @param t:
func TestPrettyJSON(t *testing.T) {
	animals := []Animal{
		{Name: "兔子", Age: 10},
		{Name: "猫", Age: 11},
	}
	prettyJSON := json.PrettyJSON(animals)
	fmt.Println(prettyJSON)
}
