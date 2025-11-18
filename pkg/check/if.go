package check

func IF[T any](check bool, rtTrue T, rtFalse T) T {
	if check {
		return rtTrue
	}
	return rtFalse
}
