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

func Async[R any](ctx context.Context, fn func() R) (R, error) {
	done := make(chan R, 1)
	go func() {
		done <- fn()
	}()
	select {
	case <-ctx.Done():
		var z R
		return z, ctx.Err()
	case r := <-done:
		return r, nil
	}
}
