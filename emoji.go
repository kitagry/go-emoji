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
		start, end := findEmojiMarkupPosition(src[nSrc:])
		if start == -1 {
			copied := copy(dst[nDst:], src[nSrc:])
			nSrc += len(src[nSrc:])
			nDst += copied
			if copied != len(src[nSrc:]) {
				err = transform.ErrShortDst
			}
			return
		}

		if end == -1 {
			if !atEOF {
				r.preSrc = src[nSrc:]
				nSrc = len(src[nSrc:])
				err = transform.ErrShortSrc
			} else {
				copied := copy(dst[nDst:], src[:start])
				nSrc += start
				nDst += copied
			}
			return
		}

		// copy before nSrc+start
		copied := copy(dst[nDst:], src[nSrc:nSrc+start])
		nSrc += start
		nDst += copied
		if copied != start {
			err = transform.ErrShortDst
			return
		}

		eng := src[nSrc : end+1]
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
// When not found, start or end is -1
func findEmojiMarkupPosition(src []byte) (start, end int) {
	start = bytes.Index(src, []byte(":"))
	if start == -1 {
		return
	}

	m := bytes.Index(src[start+1:], []byte(":"))
	if m == -1 {
		end = -1
		return
	}

	end = start + 1 + m
	return
}

func (r *Replacer) Reset() {
	r.preSrc = nil
}
