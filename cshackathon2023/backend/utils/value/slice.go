package value

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func RemoveIndex[T any](slice []T, index int) []T {
	if index < 0 || index > len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func Index[T comparable](slice []T, element T) int {
	for k, v := range slice {
		if element == v {
			return k
		}
	}
	return -1 // not found.
}

func IndexVal[T comparable](slice []*T, element T) int {
	for k, v := range slice {
		if element == *v {
			return k
		}
	}
	return -1 // not found.
}
