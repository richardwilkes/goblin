package util

import (
	"io"
	"unicode/utf8"
)

// Prefixer provides an io.Writer that will add a prefix to each line output.
type Prefixer struct {
	Prefix string
	Writer io.Writer
}

func (prefixer *Prefixer) Write(p []byte) (n int, err error) {
	if prefixer.Prefix == "" {
		return prefixer.Writer.Write(p)
	}
	prefix := []byte(prefixer.Prefix)
	var n1 int
	i := 0
	j := 0
	for j < len(p) {
		r, size := utf8.DecodeRune(p[j:])
		if r == '\n' {
			j++
			n1, err = prefixer.Writer.Write(p[i:j])
			n += n1
			if err != nil {
				return
			}
			n1, err = prefixer.Writer.Write(prefix)
			n += n1
			if err != nil {
				return
			}
			i = j
		} else {
			j += size
		}
	}
	if i != j {
		n1, err = prefixer.Writer.Write(p[i:j])
		n += n1
	}
	return
}
