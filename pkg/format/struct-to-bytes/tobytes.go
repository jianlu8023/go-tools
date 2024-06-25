package struct_to_bytes

import (
	"bytes"
	"encoding/gob"
)

// ToBytes 将结构体转换为字节数组
// 使用 gob 包进行编码，确保结构体中的字段可以被编码
// @param v 结构体对象
// @return []byte 字节数组
// @return error 编码过程中遇到的错误
func ToBytes(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
