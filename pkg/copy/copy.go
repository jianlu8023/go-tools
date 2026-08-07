package copy

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
)

// DeepCopy 使用JSON序列化和反序列化实现深拷贝
//
// 参数:
//   - src: 要拷贝的源对象
//   - dst: 指向目标对象的指针（不可为 nil）
//
// 返回值:
//   - 如果拷贝过程中出现错误则返回错误，否则返回nil
//
// 示例:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	src := Person{Name: "Alice", Age: 30}
//	var dst Person
//	err := DeepCopy(src, &dst)
//	// dst 现在是 src 的深拷贝
//
// 限制（基于 JSON 序列化实现，调用方需知悉）:
//   - 未导出字段（小写字母开头）不会被拷贝
//   - time.Time 的 monotonic 时钟会被丢弃
//   - 接口类型字段会被反序列化为底层具体类型，丢失原接口语义
//   - 不支持循环引用，遇到循环引用会因 JSON 递归序列化导致栈溢出
//   - map 的 key 必须为字符串，其他类型（如 int key）会拷贝失败
func DeepCopy[T any](src T, dst *T) error {
	if dst == nil {
		return fmt.Errorf("copy: dst must not be nil")
	}
	srcBytes, err := jsoniter.Marshal(src)
	if err != nil {
		return err
	}
	if err := jsoniter.Unmarshal(srcBytes, dst); err != nil {
		return err
	}
	return nil
}
