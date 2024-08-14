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
