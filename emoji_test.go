package emoji_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/kitagry/go-emoji"
	"golang.org/x/text/transform"
)

func ExampleReplacer() {
	r := emoji.NewReplacer()
	tf := transform.NewReader(strings.NewReader("Hello World:blush:"), r)
	io.Copy(os.Stdout, tf)

	// Output:
	// Hello World😊
}

func TestReplacer(t *testing.T) {
	tests := map[string]struct {
		// input
		src, dst []byte
		atEOF    bool

		// output
		nsrc, ndst int
		expected   []byte
		hasErr     bool
	}{
		"normal test": {
			src:      []byte(":+1:"),
			dst:      make([]byte, 100),
			atEOF:    true,
			nsrc:     4,
			ndst:     4,
			expected: []byte("👍"),
		},
		"not found emoji": {
			src:      []byte(":not found:"),
			dst:      make([]byte, 100),
			atEOF:    true,
			nsrc:     len(":not found:"),
			ndst:     len(":not found:"),
			expected: []byte(":not found:"),
		},
		"text with emoji": {
			src:      []byte("LGTM:+1:"),
			dst:      make([]byte, 100),
			atEOF:    true,
			nsrc:     8,
			ndst:     8,
			expected: []byte("LGTM👍"),
		},
		"short dist": {
			src:      []byte("LGTM:+1:"),
			dst:      make([]byte, 4),
			atEOF:    true,
			nsrc:     8,
			ndst:     4,
			expected: []byte("LGTM"),
			hasErr:   true,
		},
		"not completed emoji": {
			src:      []byte(":+1"),
			dst:      make([]byte, 100),
			atEOF:    false,
			nsrc:     3,
			ndst:     0,
			expected: []byte(""),
			hasErr:   true,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			r := emoji.NewReplacer()
			ndst, nsrc, err := r.Transform(tt.dst, tt.src, tt.atEOF)
			if !tt.hasErr && err != nil {
				t.Errorf("failed to Transform: %+v", err)
			} else if tt.hasErr && err == nil {
				t.Errorf("Transform error expected but got nil")
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
