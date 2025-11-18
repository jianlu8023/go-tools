package check

import (
	"cmp"
)

func OrUnwish[T comparable](unwish T, arr ...T) (res T) {
	var zero = res
	for _, v := range arr {
		if v != unwish && v != zero {
			return v
		}
	}
	return zero
}

func Or[T comparable](arr ...T) T {
	return cmp.Or(arr...)
}
