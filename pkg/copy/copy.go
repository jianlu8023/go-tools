package copy

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json"
)

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
