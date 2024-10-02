package fak

import "errors"
import "strings"

func ConcatErrors(errs ...error) string {
	builder := strings.Builder{}
	for i, err := range errs {
		if i > 0 {
			builder.WriteString(": ")
		}
		builder.WriteString(err.Error())
	}
	return builder.String()
}

func UnwrapAll(err error) []error {
	if err == nil {
		return nil
	}
	errs := []error{err}
	for {
		uErr := errors.Unwrap(err)
		if uErr == nil {
			return errs
		}
		errs = append(errs, uErr)
		err = uErr
	}
}

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
	if original == nil {
		return nil
	}
	return wError{wrapper, original}
}
