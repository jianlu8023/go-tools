package uuid

import (
	"github.com/google/uuid"
)

// GenUUID
// @Description: GenUUID
// @return string: uuid字符串
func GenUUID() string {
	return uuid.NewString()
}
