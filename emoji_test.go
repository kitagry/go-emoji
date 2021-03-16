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
	customEmoji, err := emoji.NewEmoji([]byte("emo"), []byte(":custom:"))
	if err != nil {
		t.Fatalf("failed to create Emoji: %+v", err)
	}

	tests := map[string]struct {
		// input
		src, dst []byte
		custom   []emoji.Emoji
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
			src:      []byte("hello:+1"),
			dst:      make([]byte, 100),
			atEOF:    false,
			nsrc:     8,
			ndst:     5,
			expected: []byte("hello"),
			hasErr:   true,
		},
		"not closed": {
			src:      []byte(":+1"),
			dst:      make([]byte, 100),
			atEOF:    true,
			nsrc:     3,
			ndst:     3,
			expected: []byte(":+1"),
			hasErr:   false,
		},
		"custom emoji": {
			src: []byte(":custom:"),
			dst: make([]byte, 100),
			custom: []emoji.Emoji{
				*customEmoji,
			},
			atEOF:    true,
			nsrc:     8,
			ndst:     3,
			expected: []byte("emo"),
			hasErr:   false,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			r := emoji.NewReplacer(tt.custom...)
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

func TestReplacer_TransformFlow(t *testing.T) {
	type flow []struct {
		dst, src []byte
		atEOF    bool

		nSrc int
		out  []byte
		err  error
	}
	tests := map[string]struct {
		flow flow
	}{
		"two flow": {
			flow: flow{
				{
					dst:   make([]byte, 0),
					src:   []byte(":+1"),
					atEOF: false,
					nSrc:  3,
					out:   []byte(""),
					err:   transform.ErrShortSrc,
				},
				{
					dst:   make([]byte, 4),
					src:   []byte(":"),
					atEOF: true,
					nSrc:  1,
					out:   []byte("👍"),
					err:   nil,
				},
			},
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			r := emoji.NewReplacer()
			for _, f := range tt.flow {
				nDst, nSrc, err := r.Transform(f.dst, f.src, f.atEOF)
				if nDst != len(f.out) {
					t.Errorf("nDst expected %d, got %d", len(f.out), nDst)
				}
				if !bytes.Equal(f.dst[:nDst], f.out) {
					t.Errorf("dst expected %v, got %v", f.out, f.dst[:nDst])
				}
				if nSrc != f.nSrc {
					t.Errorf("nSrc expected %d, got %d", f.nSrc, nSrc)
				}
				if err != f.err {
					t.Errorf("error expected %v, got %v", f.err, err)
				}
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
		"Normal replace with text": {
			r:        strings.NewReader("Hello World:blush:"),
			expected: []byte("Hello World😊"),
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
