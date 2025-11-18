package copy

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json"
)

// DeepCopy 使用JSON序列化和反序列化实现深拷贝
//
// 参数:
//   - src: 要拷贝的源对象
//   - dst: 指向目标对象的指针
//
// 返回值:
//   - 如果拷贝过程中出现错误则返回错误，否则返回nil
//
// 示例:
//   type Person struct {
//       Name string
//       Age  int
//   }
//   src := Person{Name: "Alice", Age: 30}
//   var dst Person
//   err := DeepCopy(src, &dst)
//   // dst 现在是 src 的深拷贝
func DeepCopy[T any](src T, dst *T) error {
	srcBytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(srcBytes, dst); err != nil {
		return err
	}
	return nil
}
