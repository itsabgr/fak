package fak 

import (
  "io"
  "context"
)

func ReadAll(ctx context.Context, r io.Reader, b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, 0, 256)
	}
	for {
		if err := ctx.Err(); err != nil {
			return b, err
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}

		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
	}
}
