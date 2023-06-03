package value

func Contain[T comparable](arr []T, elem T) bool {
	for _, a := range arr {
		if a == elem {
			return true
		}
	}
	return false
}

func ContainVal[T comparable](arr []*T, elem *T) bool {
	for _, a := range arr {
		if *a == *elem {
			return true
		}
	}
	return false
}
