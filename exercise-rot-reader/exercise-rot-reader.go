// 参考
// https://www.exmedia.jp/blog/a-tour-of-go%E3%81%AE%E7%B7%B4%E7%BF%92%E5%95%8F%E9%A1%8C%E3%82%92%E8%A7%A3%E8%AA%AC%E3%81%99%E3%82%8B%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA8-11-exercise-rot13reader/

package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	encryptionKey = 13
	table = alphabet + alphabet
)

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := range p {
		p[i] = r.rotate(p[i], encryptionKey)
	}
	return n, nil
}

func (r *rot13Reader) rotate(b byte, key int) byte {
	ch := rune(b)
	if !unicode.IsLetter(ch) {
		return b
	}

	key = key % len(alphabet)

	isUpper := false
	if unicode.IsUpper(ch) {
		isUpper = true
		ch = unicode.ToLower(ch)
	}

	idx := byte(ch) - 'a' + byte(key)
	rotatedCh := rune(table[idx])
	if isUpper {
		rotatedCh = unicode.ToUpper(rotatedCh)
	}
	return byte(rotatedCh)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbgr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}