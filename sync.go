package fak

import "context"

func LockContext(ctx context.Context, lock interface{ TryLock() bool }) error {
	for lock.TryLock() == false {
		if err := ctx.Err(); err != nil {
			return err
		}
		runtime.Gosched()
	}
	return nil
}
