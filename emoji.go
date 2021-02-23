package emoji

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
	mapMarkupToEmoji map[string]Emoji
	preSrc           []byte
}

var _ transform.Transformer = (*Replacer)(nil)

func NewReplacer(customEmojis ...Emoji) *Replacer {
	mapMarkupToEmoji := make(map[string]Emoji, len(emojis))
	for _, e := range emojis {
		mapMarkupToEmoji[string(e.markup)] = e
	}
	for _, e := range customEmojis {
		mapMarkupToEmoji[string(e.markup)] = e
	}
	return &Replacer{mapMarkupToEmoji: mapMarkupToEmoji}
}

// Transform implements transform.Transformer.Transform.
// Transform replaces markdown emoji markup to emoji.
func (r *Replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	_src := src
	if len(r.preSrc) != 0 {
		_src = make([]byte, len(r.preSrc)+len(src))
		copy(_src, r.preSrc)
		copy(_src[len(r.preSrc):], src)
	}

	var preSrc []byte
	nDst, nSrc, preSrc, err = r.transform(dst, _src, atEOF)

	if nSrc < len(r.preSrc) {
		nSrc = 0
		r.preSrc = r.preSrc[nSrc:]
	} else {
		nSrc -= len(r.preSrc)
		r.preSrc = preSrc
	}
	return
}

func (r *Replacer) transform(dst, src []byte, atEOF bool) (nDst, nSrc int, preSrc []byte, err error) {
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
				preSrc = src[nSrc:]
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
		emoji, ok := r.mapMarkupToEmoji[string(eng)]
		if ok {
			target = emoji.emoji
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
