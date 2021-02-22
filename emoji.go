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
			copied := copy(dst[nDst:], src[nSrc:])
			nSrc += copied
			nDst += copied
			if copied != len(src[nSrc:]) {
				err = transform.ErrShortDst
			}
			return
		}

		m := bytes.Index(src[nSrc+n+1:], colon)
		// not found
		if m == -1 {
			copied := copy(dst[nDst:], src[:n])
			nSrc += copied
			nDst += copied
			err = transform.ErrShortSrc
			return
		}

		// copy before nSrc+n
		copied := copy(dst[nDst:], src[nSrc:nSrc+n])
		nSrc += copied
		nDst += copied
		if copied != n {
			err = transform.ErrShortDst
			return
		}

		eng := src[nSrc : nSrc+1+m+1]
		target := eng
		for _, e := range emojis {
			if bytes.Compare(eng, e.english) == 0 {
				target = e.emoji
			}
		}
		copied = copy(dst[nDst:], target)
		nDst += copied
		nSrc += len(eng)
		if copied != len(target) {
			err = transform.ErrShortDst
			return
		}
	}
}

func (r *Replacer) Reset() {
}
