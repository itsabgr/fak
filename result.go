package fak

type Result[T any] struct {
	res T
	err error
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Must() T {
	return Must(r.Result())
}

func (r Result[T]) Result() (T, error) {
	return r.res, r.err
}

func OK[T any](res T) Result[T] {
	return Result[T]{
		res,
		nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		Zero[T](),
		err,
	}
}
