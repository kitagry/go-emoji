package emoji

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
}

var _ transform.Transformer = (*Replacer)(nil)

func NewReplacer() *Replacer {
	return &Replacer{}
}

// Transform implements transform.Transformer.Transform.
// Transform replaces markdown emoji markup to emoji.
func (r *Replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	colon := []byte(":")
	for {
		n := bytes.Index(src[nSrc:], colon)
		// not found
		if n == -1 {
			copy(dst[nDst:], src[nSrc:])
			nDst += len(src[nSrc:])
			return
		}

		m := bytes.Index(src[nSrc+n+1:], colon)
		// not found
		if m == -1 {
			copy(dst[nDst:], src[:n])
			nSrc += n
			nDst += n
			err = transform.ErrShortSrc
			return
		}

		// copy before nSrc+n
		copy(dst[nDst:], src[nSrc:nSrc+n])
		nSrc += n
		nDst += n

		eng := src[nSrc : nSrc+1+m+1]
		target := eng
		for _, e := range emojis {
			if bytes.Compare(eng, e.english) == 0 {
				target = e.emoji
			}
		}
		copy(dst[nDst:], target)
		nDst += len(target)
		nSrc += len(eng)
	}
}

func (r *Replacer) Reset() {
}
