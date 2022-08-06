package generator_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/tinyfont/cmd/tinyfontgen/generator"
)

func TestTinyFontFile(t *testing.T) {

	font := tinyfont.Font{
		BBox: [4]int8{28, 26, -2, -25},
		Glyphs: []tinyfont.Glyph{
			/*   */ tinyfont.Glyph{Rune: 32, Width: 0x1, Height: 0x1, XAdvance: 0xb, XOffset: 0, YOffset: 0, Bitmaps: []byte{0x0}},
			/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x8, Height: 0x15, XAdvance: 0xd, XOffset: 1, YOffset: -21, Bitmaps: []byte{0xff, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0xff}},
		},
		YAdvance: 0x1c,
	}

	tt := map[string]struct {
		file      string
		fontName  string
		font      tinyfont.Font
		expected  string
		wantedErr error
	}{
		"gopher font 32pt": {
			file:      "gophers32pt.go",
			fontName:  "Regular32pt",
			font:      font,
			expected:  "./text-expected/gophers32pt_expected.go",
			wantedErr: nil,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := generator.TinyFontFile(tc.file, tc.fontName, tc.font)
			if tc.wantedErr != nil {
				assert.ErrorIs(t, err, tc.wantedErr)
			} else {
				assert.NoError(t, err)

				// open the expected file for comparison
				expectedFile, err := os.Open(tc.expected)
				assert.NoErrorf(t, err, "failed to open expected file")

				assert.Equal(t, expectedFile, got)
			}
		})
	}
}
