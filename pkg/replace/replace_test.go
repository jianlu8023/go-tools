package replace

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	template := "https://api.github.com/repo/{{owner}}/{{repo}}"

	str := Replace(template, map[string]string{
		"{{owner}}": "ceshi",
		"{{repo}}":  "ceshi",
	})
	fmt.Println(template)
	fmt.Println(str)
}
