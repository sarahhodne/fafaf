package main

import (
	"io"
	"os"
)

type FafafReader struct {
	state bool
}

func (fr *FafafReader) Read(p []byte) (int, error) {
	for i := range p {
		if fr.state {
			p[i] = 'a'
		} else {
			p[i] = 'f'
		}
		fr.state = !fr.state
	}

	return len(p), nil
}

func main() {
	io.Copy(os.Stdout, &FafafReader{})
}
