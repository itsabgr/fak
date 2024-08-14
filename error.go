package fak

import "errors"

func Throw(err any) {
	if err != nil {
		panic(err)
	}
}

func Assert(cond bool, err error) {
	if !cond {
		if err == nil {
			panic(ErrAssertionFailed)
		}
		panic(err)
	}
}

var ErrAssertionFailed = errors.New("assertion failed")

func Must[T any](res T, err any) T {
	Throw(err)
	return res
}

func Try[T any](panicableFunc func() T) (res T, err any) {
	defer func() {
		err = recover()
	}()
	res = panicableFunc()
	return
}

type wError struct {
	wrapperError  error
	originalError error
}

func (w wError) Err() error {
	return w.wrapperError
}

func (w wError) Error() string {
	return w.wrapperError.Error()
}

func (w wError) Unwrap() error {
	return w.originalError
}

func Wrap(wrapper, original error) error {
	return wError{wrapper, original}
}
