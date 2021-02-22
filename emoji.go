package emoji

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
	preSrc []byte
}

var _ transform.Transformer = (*Replacer)(nil)

func NewReplacer() *Replacer {
	return &Replacer{}
}

// Transform implements transform.Transformer.Transform.
// Transform replaces markdown emoji markup to emoji.
func (r *Replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for {
		offset, length := findEmojiMarkupPosition(src[nSrc:])
		if offset == -1 {
			copied := copy(dst[nDst:], src[nSrc:])
			nSrc += len(src[nSrc:])
			nDst += copied
			if copied != len(src[nSrc:]) {
				err = transform.ErrShortDst
			}
			return
		}

		// copy before nSrc+sOffset
		// And then, src[nSrc] should be ':'
		copied := copy(dst[nDst:], src[nSrc:nSrc+offset])
		nSrc += offset
		nDst += copied
		if copied != offset {
			err = transform.ErrShortDst
			return
		}

		if length == -1 {
			if !atEOF {
				r.preSrc = src[nSrc:]
				nSrc += len(src[nSrc:])
				err = transform.ErrShortSrc
			} else {
				copied := copy(dst[nDst:], src[nSrc:])
				nSrc += len(src[nSrc:])
				nDst += copied
			}
			return
		}

		eng := src[nSrc : nSrc+length]
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

// find emoji markup.
// When not found, offset or length is -1
func findEmojiMarkupPosition(src []byte) (offset, length int) {
	offset = bytes.Index(src, []byte(":"))
	if offset == -1 {
		return
	}

	m := bytes.Index(src[offset+1:], []byte(":"))
	if m == -1 {
		length = -1
		return
	}

	length = m + 2
	return
}

func (r *Replacer) Reset() {
	r.preSrc = nil
}
