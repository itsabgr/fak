package fak

func Ptr[T any](t T) *T {
  return &t
}

func Zero[T any]() (z T) {
	return z
}
