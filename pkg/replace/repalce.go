package replace

import (
	"regexp"
)

// Replace
// @Description: Replace
// @author ght
// @date 2023-10-07 16:59:00
// @param template: 模板
// @param replacements: 替换的信息
// @return string: 替换后
func Replace(template string, replacements map[string]string) string {
	for pattern, replacement := range replacements {
		re := regexp.MustCompile(pattern)
		template = re.ReplaceAllString(template, replacement)

	}
	return template
}
