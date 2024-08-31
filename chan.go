package fak

func Flush[T any](c <-chan T) (i int) {
	for {
		select {
		case <-c:
			i++
		default:
			return
		}
	}
}

func Chan[R any](fn func() R) <-chan R {
	c := make(chan R, 1)
	go func() {
		c <- fn()
	}()
	return c
}
