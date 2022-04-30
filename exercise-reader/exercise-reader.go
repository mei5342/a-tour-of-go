package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	// Validate：https://github.com/golang/tour/blob/master/reader/validate.go#L13
	reader.Validate(MyReader{})
}