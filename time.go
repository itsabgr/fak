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
