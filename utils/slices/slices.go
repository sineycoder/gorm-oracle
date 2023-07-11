package slices

func Map[T, V any](arr []T, fn func(T) V) []V {
	ret := make([]V, len(arr))
	for idx, item := range arr {
		ret[idx] = fn(item)
	}
	return ret
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	newArr := make([]T, 0)
	for _, item := range arr {
		if !fn(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func ToBoolMap[T comparable](arr []T) map[T]bool {
	ret := make(map[T]bool, 0)
	for _, item := range arr {
		if _, exist := ret[item]; !exist {
			ret[item] = true
		}
	}
	return ret
}
