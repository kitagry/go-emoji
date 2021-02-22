package emoji_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kitagry/go-emoji"
	"golang.org/x/text/transform"
)

func TestReplacer(t *testing.T) {
	tests := map[string]struct {
		// input
		src, dst []byte

		// output
		nsrc, ndst int
		atEOF      bool
		expected   []byte
	}{
		"normal test": {
			src:      []byte(":+1:"),
			dst:      make([]byte, 100),
			nsrc:     4,
			ndst:     4,
			atEOF:    true,
			expected: []byte("👍"),
		},
		"another emoji": {
			src:      []byte(":man_with_turban:"),
			dst:      make([]byte, 100),
			nsrc:     17,
			ndst:     13,
			atEOF:    true,
			expected: []byte("👳‍♂️"),
		},
		"not found emoji": {
			src:      []byte(":not found:"),
			dst:      make([]byte, 100),
			nsrc:     len(":not found:"),
			ndst:     len(":not found:"),
			atEOF:    true,
			expected: []byte(":not found:"),
		},
		"text with emoji": {
			src:      []byte("LGTM:+1:"),
			dst:      make([]byte, 100),
			nsrc:     8,
			ndst:     8,
			atEOF:    true,
			expected: []byte("LGTM👍"),
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			r := emoji.NewReplacer()
			ndst, nsrc, err := r.Transform(tt.dst, tt.src, tt.atEOF)
			if err != nil {
				t.Errorf("failed to Transform: %+v", err)
			}

			if ndst != tt.ndst {
				t.Errorf(`ndst expected %d, got %d`, tt.ndst, ndst)
			} else if bytes.Compare(tt.dst[:ndst], tt.expected) != 0 {
				t.Errorf(`expected %v, got %v`, tt.dst[:ndst], tt.expected)
			}

			if tt.nsrc != nsrc {
				t.Errorf("nsrc expected %d, got %d", tt.nsrc, nsrc)
			}
		})
	}
}

func TestReplacerWithReader(t *testing.T) {
	tests := map[string]struct {
		r io.Reader

		expected []byte
	}{
		"Normal replace": {
			r:        strings.NewReader(":+1:"),
			expected: []byte("👍"),
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			r := emoji.NewReplacer()
			tf := transform.NewReader(tt.r, r)
			result, err := ioutil.ReadAll(tf)
			if err != nil {
				t.Errorf("failed to ReadAll: %+v", err)
			}

			if bytes.Compare(result, tt.expected) != 0 {
				t.Errorf("readAll expected %v, got %v", tt.expected, result)
			}
		})
	}
}
