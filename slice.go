package fak

func Variadic[T any](elems ...T) []T {
	return elems
}

func Slice[T ~[]E, E any](list T, start, end int) T {
	length := len(list)
	if end < 0 {
		end = length
	}
	if end <= start || length < start || length >= end {
		return nil
	}
	return list[start:end]
}


func Or[T comparable](vals ...T) (t T) {
	var zero T
	for _, t = range vals {
		if t != zero {
			return t
		}
	}
	return t
}
