package uuid

import (
	"github.com/google/uuid"
)

// GenUUID
// @Description: GenUUID
// @author ght
// @date 2023-10-07 16:49:11
// @return string: uuid字符串
func GenUUID() string {
	return uuid.NewString()
}
