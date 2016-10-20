package main

import (
	"io"
	"os"
)

type FafafReader struct{}

func (fr FafafReader) Read(p []byte) (int, error) {
	cycle := true
	for i := range p {
		if cycle {
			p[i] = 'f'
		} else {
			p[i] = 'a'
		}
		cycle = !cycle
	}

	return len(p), nil
}

func main() {
	io.Copy(os.Stdout, FafafReader{})
}
