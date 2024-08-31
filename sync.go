package fak

import (
	"context"
	"runtime"
)

func LockContext(ctx context.Context, lock interface{ TryLock() bool }) error {
	for lock.TryLock() == false {
		if err := ctx.Err(); err != nil {
			return err
		}
		runtime.Gosched()
	}
	return nil
}

func Async[R any](ctx context.Context, fn func() (R, error)) (r R, err error) {
	done := make(chan struct{}, 1)
	go func() {
		r, err = fn()
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		var zero R
		return zero, ctx.Err()
	case <-done:
		return r, err
	}
}
