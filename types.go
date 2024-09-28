package fak

func Ptr[T any](t T) *T {
  return &t
}

func Zero[T any]() (z T) {
	return z
}

func CastInterfaces[T any, T2 any](slice ...T) []T2 {
	res := make([]T2, len(slice))
	for i, s := range slice {
		res[i] = any(s).(T2)
	}
	return res
}
