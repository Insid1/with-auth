package utils

func IsInList[T comparable](element T, list []T) bool {
	for _, v := range list {
		if element == v {
			return true
		}
	}

	return false
}
