package replace

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/replace"
)

// TestReplace
// @Description: TestReplace
// @author ght
// @date 2023-10-07 17:22:06
// @param t:
func TestReplace(t *testing.T) {
	template := "https://api.github.com/repo/{{owner}}/{{repo}}"

	str := replace.Replace(template, map[string]string{
		"{{owner}}": "ceshi",
		"{{repo}}":  "ceshi",
	})
	fmt.Println(template)
	fmt.Println(str)
}
