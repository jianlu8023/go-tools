package check

func In[T comparable](target T, arr ...T) bool {
	for _, obj := range arr {
		if target == obj {
			return true
		}
	}
	return false
}
