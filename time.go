package fak

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		timer.Stop()
		Flush(timer.C)
		return ctx.Err()
	}
}

func Timeout[T any](ctx context.Context, timeout time.Duration, fn func(timeoutCtx context.Context) (T, error)) (T, error) {
	timeoutCtx, cancelTimeout := context.WithTimeout(ctx, timeout)
	defer cancelTimeout()
	return fn(timeoutCtx)
}
