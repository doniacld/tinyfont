package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var BoldItalic12pt7b = tinyfont.Font{
	BBox: [4]int8{27, 23, -3, -17},
	Glyphs: []tinyfont.Glypher{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x6, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x8, Height: 0x11, XAdvance: 0x9, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0x7, 0x7, 0x7, 0xf, 0xe, 0xe, 0xc, 0xc, 0x8, 0x18, 0x10, 0x0, 0x0, 0x60, 0xf0, 0xf0, 0x60}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x9, Height: 0x7, XAdvance: 0xd, XOffset: 4, YOffset: -15, Bitmaps: []uint8{0x61, 0xf1, 0xf8, 0xf8, 0x6c, 0x34, 0x12, 0x8}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xe, Height: 0x10, XAdvance: 0xc, XOffset: -1, YOffset: -15, Bitmaps: []uint8{0x1, 0x8c, 0x6, 0x60, 0x31, 0x80, 0xcc, 0x6, 0x30, 0xff, 0xf0, 0xc6, 0x3, 0x18, 0xc, 0xc0, 0x63, 0xf, 0xff, 0xc, 0x60, 0x33, 0x1, 0x8c, 0x6, 0x30, 0x19, 0x80}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xc, Height: 0x14, XAdvance: 0xc, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x0, 0x80, 0x8, 0x7, 0xc1, 0x96, 0x31, 0x33, 0x13, 0x3a, 0x23, 0xe0, 0x1e, 0x1, 0xf0, 0x7, 0x80, 0x7c, 0x5, 0xc4, 0xcc, 0x48, 0xcc, 0x8c, 0xf8, 0x83, 0x30, 0x1e, 0x1, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x12, Height: 0x12, XAdvance: 0x14, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x0, 0x2, 0x7, 0x83, 0x3, 0x9f, 0x81, 0xc4, 0x20, 0x71, 0x10, 0x3c, 0x44, 0xe, 0x22, 0x3, 0x88, 0x80, 0xe4, 0x40, 0x1e, 0x31, 0xe0, 0x8, 0xe4, 0x6, 0x71, 0x1, 0x3c, 0x40, 0x8e, 0x10, 0x23, 0x88, 0x10, 0xe2, 0x4, 0x39, 0x2, 0x7, 0x80}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0x10, Height: 0x11, XAdvance: 0x13, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x0, 0xf0, 0x1, 0x98, 0x3, 0x98, 0x3, 0x98, 0x3, 0xb0, 0x3, 0xe0, 0x3, 0x80, 0xf, 0x9f, 0x19, 0xce, 0x31, 0xcc, 0x61, 0xc8, 0xe1, 0xc8, 0xe0, 0xf0, 0xe0, 0xe0, 0xf0, 0x70, 0x78, 0x79, 0x3f, 0xbe}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x3, Height: 0x7, XAdvance: 0x7, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0x7f, 0xed, 0x20}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x7, Height: 0x15, XAdvance: 0x8, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x2, 0x8, 0x20, 0xc3, 0xe, 0x18, 0x30, 0xe1, 0x83, 0x6, 0xc, 0x18, 0x30, 0x20, 0x40, 0x80, 0x81, 0x1, 0x0}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x7, Height: 0x15, XAdvance: 0x8, XOffset: -1, YOffset: -15, Bitmaps: []uint8{0x10, 0x10, 0x20, 0x20, 0x40, 0xc1, 0x83, 0x6, 0xc, 0x18, 0x70, 0xe1, 0x83, 0xc, 0x18, 0x61, 0x86, 0x0, 0x0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0xa, Height: 0xa, XAdvance: 0xc, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0xc, 0x33, 0x6c, 0x9b, 0xae, 0x1c, 0x3f, 0xec, 0x9b, 0x36, 0xc, 0x2, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0xf, 0xff, 0xff, 0xf0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x5, Height: 0x8, XAdvance: 0x6, XOffset: -2, YOffset: -3, Bitmaps: []uint8{0x31, 0xce, 0x31, 0x8, 0x98}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x6, Height: 0x3, XAdvance: 0x8, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xff, 0xff, 0xc0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x4, Height: 0x4, XAdvance: 0x6, XOffset: 0, YOffset: -2, Bitmaps: []uint8{0x6f, 0xf6}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0xa, Height: 0x10, XAdvance: 0x8, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1, 0x80, 0x60, 0x30, 0xc, 0x7, 0x1, 0x80, 0xe0, 0x30, 0x1c, 0x6, 0x1, 0x80, 0xc0, 0x30, 0x18, 0x6, 0x3, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x3, 0x81, 0xc8, 0x71, 0x1c, 0x33, 0x86, 0xe1, 0xdc, 0x3b, 0x87, 0xe0, 0xfc, 0x3b, 0x87, 0x70, 0xec, 0x39, 0x87, 0x31, 0xc2, 0x30, 0x3c, 0x0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0xa, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1, 0xc3, 0xf0, 0x38, 0xe, 0x3, 0x81, 0xe0, 0x70, 0x1c, 0xf, 0x3, 0x80, 0xe0, 0x38, 0x1e, 0x7, 0x1, 0xc0, 0xf0, 0xff, 0x80}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x7, 0x81, 0xf8, 0x47, 0x90, 0x70, 0xe, 0x1, 0xc0, 0x30, 0xe, 0x1, 0x80, 0x60, 0x18, 0x6, 0x1, 0x80, 0x40, 0x8f, 0xf3, 0xfc, 0xff, 0x80}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x7, 0xc3, 0x3c, 0x3, 0x80, 0x70, 0xc, 0x3, 0x81, 0xc0, 0xfc, 0x7, 0xc0, 0x78, 0x7, 0x0, 0xe0, 0x1c, 0x3, 0x30, 0xe7, 0x10, 0x7c, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0xd, Height: 0x10, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x0, 0x10, 0x1, 0x80, 0x3c, 0x3, 0xe0, 0x2e, 0x2, 0x70, 0x23, 0x82, 0x38, 0x21, 0xc2, 0xe, 0x1f, 0xf9, 0xff, 0xc0, 0x38, 0x1, 0xc0, 0x1c, 0x0, 0xe0}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0xc, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x7, 0xf0, 0x7e, 0xf, 0xe0, 0x80, 0x8, 0x1, 0xe0, 0x1f, 0x83, 0xf8, 0x3, 0xc0, 0x1c, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0x8, 0x61, 0x8f, 0x30, 0x7c, 0x0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x0, 0x60, 0x78, 0x1c, 0xf, 0x1, 0xc0, 0x70, 0x1f, 0xc3, 0x8c, 0xe1, 0xdc, 0x3b, 0x87, 0x61, 0xec, 0x3d, 0x87, 0x31, 0xe2, 0x38, 0x3c, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0xb, Height: 0x10, XAdvance: 0xc, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0x3f, 0xef, 0xf9, 0xff, 0x60, 0xc8, 0x18, 0x6, 0x0, 0x80, 0x30, 0xc, 0x1, 0x80, 0x60, 0x1c, 0x3, 0x0, 0xc0, 0x18, 0x6, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x3, 0x81, 0x88, 0x61, 0x8c, 0x31, 0x86, 0x38, 0xc7, 0xb0, 0x78, 0xf, 0x86, 0x71, 0x87, 0x60, 0x6c, 0xd, 0x81, 0xb0, 0x63, 0x18, 0x3e, 0x0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0xb, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x7, 0x81, 0xc8, 0x71, 0x8e, 0x33, 0xc6, 0x70, 0xce, 0x39, 0xc7, 0x38, 0xe3, 0x38, 0x3f, 0x1, 0xc0, 0x38, 0xe, 0x3, 0x81, 0xc0, 0xe0, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x7, Height: 0xc, XAdvance: 0x6, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0xc, 0x3c, 0x78, 0x60, 0x0, 0x0, 0x0, 0x61, 0xe3, 0xc3, 0x0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x8, Height: 0xf, XAdvance: 0x6, XOffset: -1, YOffset: -10, Bitmaps: []uint8{0xe, 0xf, 0xf, 0xe, 0x0, 0x0, 0x0, 0x0, 0x38, 0x38, 0x38, 0x18, 0x10, 0x20, 0x40}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0xc, Height: 0xd, XAdvance: 0xe, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x0, 0x10, 0x7, 0x1, 0xf0, 0x7c, 0x3f, 0xf, 0x80, 0xe0, 0xf, 0x80, 0x3e, 0x0, 0xf8, 0x3, 0xe0, 0x7, 0x0, 0x10}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xc, Height: 0x6, XAdvance: 0xe, XOffset: 2, YOffset: -8, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xd, Height: 0xd, XAdvance: 0xe, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x80, 0x7, 0x0, 0x3f, 0x0, 0x3e, 0x0, 0x7c, 0x0, 0xf8, 0x1, 0xe0, 0x1f, 0x7, 0xe0, 0xf8, 0x1f, 0x1, 0xe0, 0xc, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x9, Height: 0x11, XAdvance: 0xc, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0x1e, 0x19, 0x8c, 0xe6, 0x70, 0x38, 0x38, 0x1c, 0x18, 0x18, 0x8, 0x8, 0x0, 0x0, 0x3, 0x3, 0xc1, 0xe0, 0x60, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x11, Height: 0x10, XAdvance: 0x14, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x3, 0xf0, 0x7, 0x6, 0x6, 0x0, 0x86, 0xe, 0x66, 0xd, 0xdb, 0xc, 0xe7, 0x6, 0x33, 0x83, 0x31, 0xc3, 0x18, 0xe1, 0x8c, 0x70, 0xcc, 0x4c, 0x66, 0x46, 0x1f, 0xc1, 0x80, 0x0, 0x30, 0x10, 0x7, 0xf0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x10, Height: 0x11, XAdvance: 0x11, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x0, 0x10, 0x0, 0x30, 0x0, 0x70, 0x0, 0x70, 0x0, 0xf0, 0x1, 0xf0, 0x1, 0x78, 0x3, 0x78, 0x2, 0x38, 0x4, 0x38, 0xc, 0x38, 0xf, 0xf8, 0x18, 0x3c, 0x30, 0x3c, 0x20, 0x3c, 0x60, 0x3c, 0xf8, 0x7f}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x10, Height: 0x11, XAdvance: 0xf, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xfc, 0x7, 0x9e, 0x7, 0xf, 0x7, 0xf, 0xf, 0xf, 0xf, 0x1e, 0xe, 0x3c, 0xf, 0xe0, 0x1e, 0x3c, 0x1e, 0x1e, 0x1c, 0x1e, 0x3c, 0x1e, 0x3c, 0x1e, 0x3c, 0x3e, 0x38, 0x3c, 0x7c, 0x78, 0xff, 0xe0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xf, Height: 0x11, XAdvance: 0xf, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1, 0xf2, 0xe, 0x1c, 0x38, 0x18, 0xe0, 0x33, 0xc0, 0x4f, 0x0, 0x9e, 0x0, 0x7c, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0xc0, 0x7, 0x80, 0xf, 0x0, 0x1e, 0x0, 0x1e, 0x4, 0x1e, 0x30, 0xf, 0x80}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x12, Height: 0x11, XAdvance: 0x11, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xfc, 0x1, 0xe3, 0xc0, 0x70, 0x78, 0x1c, 0xe, 0xf, 0x3, 0xc3, 0xc0, 0xf0, 0xe0, 0x3c, 0x38, 0xf, 0x1e, 0x3, 0xc7, 0x81, 0xf1, 0xc0, 0x78, 0xf0, 0x1e, 0x3c, 0xf, 0xf, 0x3, 0xc3, 0x81, 0xc1, 0xe1, 0xe0, 0xff, 0xe0, 0x0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x11, Height: 0x11, XAdvance: 0xf, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xff, 0x83, 0xc1, 0xc1, 0xc0, 0x40, 0xe0, 0x20, 0xf0, 0x0, 0x78, 0xc0, 0x38, 0x40, 0x1f, 0xe0, 0x1e, 0x70, 0xf, 0x18, 0x7, 0x8, 0x3, 0x84, 0x3, 0xc0, 0x61, 0xe0, 0x20, 0xe0, 0x30, 0xf8, 0x78, 0xff, 0xfc, 0x0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x10, Height: 0x11, XAdvance: 0xf, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xff, 0x7, 0x87, 0x7, 0x2, 0x7, 0x2, 0xf, 0x0, 0xf, 0x18, 0xe, 0x10, 0xf, 0xf0, 0x1e, 0x70, 0x1e, 0x30, 0x1c, 0x20, 0x1c, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x38, 0x0, 0x7c, 0x0, 0xfe, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x11, Height: 0x11, XAdvance: 0x11, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1, 0xf9, 0x3, 0xc3, 0x83, 0x81, 0xc3, 0x80, 0x43, 0xc0, 0x23, 0xc0, 0x1, 0xe0, 0x1, 0xf0, 0x0, 0xf0, 0x3f, 0xf8, 0xf, 0x3c, 0x7, 0x9e, 0x3, 0xcf, 0x1, 0xc3, 0x80, 0xe1, 0xe0, 0xf0, 0x78, 0x70, 0xf, 0xe0, 0x0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x14, Height: 0x11, XAdvance: 0x12, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xe7, 0xf0, 0x78, 0x3c, 0x7, 0x83, 0xc0, 0x70, 0x3c, 0xf, 0x3, 0x80, 0xf0, 0x78, 0xe, 0x7, 0x80, 0xe0, 0x70, 0x1f, 0xff, 0x1, 0xe0, 0xf0, 0x1c, 0xf, 0x3, 0xc0, 0xe0, 0x3c, 0x1e, 0x3, 0xc1, 0xe0, 0x38, 0x1e, 0x7, 0xc3, 0xe0, 0xfe, 0x7f, 0x0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0xa, Height: 0x11, XAdvance: 0x9, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xc1, 0xe0, 0x70, 0x1c, 0xf, 0x3, 0xc0, 0xe0, 0x38, 0x1e, 0x7, 0x81, 0xc0, 0x70, 0x3c, 0xf, 0x3, 0x81, 0xf0, 0xfe, 0x0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0xe, Height: 0x12, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1, 0xfc, 0x3, 0xc0, 0xf, 0x0, 0x38, 0x0, 0xe0, 0x7, 0x80, 0x1e, 0x0, 0x70, 0x1, 0xc0, 0xf, 0x0, 0x3c, 0x0, 0xe0, 0x7, 0x80, 0x1e, 0xe, 0x70, 0x3b, 0xc0, 0xce, 0x1, 0xf0, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x11, Height: 0x11, XAdvance: 0x10, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xef, 0x83, 0xc1, 0x81, 0xc1, 0x80, 0xe1, 0x80, 0xf1, 0x80, 0x79, 0x0, 0x39, 0x0, 0x1f, 0x80, 0x1f, 0xe0, 0xf, 0x70, 0x7, 0x3c, 0x7, 0x8e, 0x3, 0xc7, 0x1, 0xe3, 0xc0, 0xe0, 0xe0, 0xf8, 0x78, 0xfe, 0xfe, 0x0}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xf, Height: 0x11, XAdvance: 0xf, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xe0, 0xf, 0x0, 0x1c, 0x0, 0x38, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0x80, 0x7, 0x0, 0x1e, 0x0, 0x3c, 0x0, 0x70, 0x0, 0xe0, 0x3, 0xc0, 0x27, 0x0, 0xce, 0x3, 0x3c, 0x1e, 0xff, 0xfc}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x17, Height: 0x11, XAdvance: 0x15, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0xf, 0x80, 0x7e, 0xf, 0x0, 0xf0, 0x1e, 0x3, 0xe0, 0x3c, 0xf, 0x80, 0xb8, 0x17, 0x1, 0x70, 0x5e, 0x2, 0xf1, 0xbc, 0x5, 0xe2, 0x70, 0x11, 0xc8, 0xe0, 0x23, 0xb3, 0xc0, 0x47, 0x47, 0x81, 0xf, 0x8e, 0x2, 0x1e, 0x1c, 0x4, 0x38, 0x78, 0x8, 0x70, 0xf0, 0x30, 0xc3, 0xe0, 0xf9, 0x8f, 0xe0}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x13, Height: 0x11, XAdvance: 0x11, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0x3, 0xe0, 0xf0, 0x38, 0x1e, 0x2, 0x3, 0xe0, 0xc0, 0xbc, 0x10, 0x13, 0xc2, 0x2, 0x78, 0x40, 0x47, 0x90, 0x10, 0xf2, 0x2, 0xf, 0x40, 0x41, 0xe8, 0x18, 0x1e, 0x2, 0x3, 0xc0, 0x40, 0x38, 0x8, 0x6, 0x3, 0x0, 0x40, 0x10, 0x8, 0x0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x10, Height: 0x11, XAdvance: 0x10, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1, 0xf8, 0x7, 0x1c, 0xe, 0xe, 0x1e, 0xf, 0x3c, 0xf, 0x3c, 0xf, 0x78, 0xf, 0x78, 0xf, 0xf8, 0x1f, 0xf0, 0x1e, 0xf0, 0x1e, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x78, 0x70, 0x70, 0x38, 0xe0, 0x1f, 0x80}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x10, Height: 0x11, XAdvance: 0xe, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xfc, 0x7, 0x9e, 0x7, 0xf, 0x7, 0xf, 0xf, 0xf, 0xf, 0xf, 0xe, 0x1e, 0xe, 0x3c, 0x1f, 0xf0, 0x1e, 0x0, 0x1c, 0x0, 0x1c, 0x0, 0x3c, 0x0, 0x38, 0x0, 0x38, 0x0, 0x7c, 0x0, 0xfe, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x10, Height: 0x15, XAdvance: 0x10, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1, 0xf8, 0x7, 0x1c, 0xe, 0xe, 0x1e, 0xf, 0x3c, 0xf, 0x3c, 0xf, 0x78, 0xf, 0x78, 0x1f, 0xf8, 0x1f, 0xf0, 0x1e, 0xf0, 0x1e, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x78, 0x70, 0x70, 0x39, 0xc0, 0xe, 0x0, 0x8, 0x2, 0x3f, 0x4, 0x7f, 0xf8, 0x83, 0xf0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x10, Height: 0x11, XAdvance: 0x10, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xf8, 0x7, 0x9e, 0x7, 0x8f, 0x7, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0x1e, 0xe, 0x3c, 0x1f, 0xf0, 0x1e, 0xf0, 0x1c, 0xf0, 0x3c, 0xf0, 0x3c, 0x78, 0x3c, 0x78, 0x3c, 0x78, 0x7c, 0x3c, 0xfe, 0x3e}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0xc, Height: 0x11, XAdvance: 0xc, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x7, 0x91, 0xc7, 0x18, 0x73, 0x82, 0x38, 0x23, 0xc0, 0x3e, 0x1, 0xf0, 0xf, 0x80, 0x7c, 0x1, 0xe0, 0x1e, 0x40, 0xe4, 0xe, 0x60, 0xce, 0x1c, 0x9f, 0x0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xf, Height: 0x11, XAdvance: 0xe, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0x7f, 0xfe, 0xe7, 0x9d, 0xe, 0x16, 0x3c, 0x20, 0x78, 0x40, 0xe0, 0x1, 0xc0, 0x7, 0x80, 0xf, 0x0, 0x1c, 0x0, 0x38, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0x80, 0xf, 0x0, 0x1e, 0x0, 0xff, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x10, Height: 0x11, XAdvance: 0x11, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0x7f, 0x1f, 0x3c, 0xe, 0x38, 0x4, 0x38, 0xc, 0x78, 0x8, 0x78, 0x8, 0x70, 0x8, 0x70, 0x10, 0xf0, 0x10, 0xf0, 0x10, 0xf0, 0x10, 0xf0, 0x20, 0xf0, 0x20, 0xf0, 0x20, 0xf0, 0x40, 0x78, 0xc0, 0x3f, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0x10, Height: 0x10, XAdvance: 0x11, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0xff, 0x1f, 0x3c, 0x6, 0x3c, 0x4, 0x3c, 0x8, 0x3c, 0x8, 0x3c, 0x10, 0x1c, 0x20, 0x1c, 0x20, 0x1e, 0x40, 0x1e, 0x80, 0x1e, 0x80, 0x1f, 0x0, 0xe, 0x0, 0xe, 0x0, 0xc, 0x0, 0x8, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x15, Height: 0x10, XAdvance: 0x16, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0xfe, 0x7c, 0x79, 0xe1, 0xc1, 0x8f, 0xe, 0x8, 0x78, 0x70, 0x43, 0xc7, 0x84, 0x1e, 0x3e, 0x20, 0x72, 0xf2, 0x3, 0x97, 0x90, 0x1d, 0x1d, 0x0, 0xe8, 0xe8, 0x7, 0x87, 0x80, 0x3c, 0x3c, 0x1, 0xc1, 0xc0, 0xe, 0xe, 0x0, 0x20, 0x60, 0x1, 0x2, 0x0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x11, Height: 0x11, XAdvance: 0x11, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xcf, 0x83, 0xc1, 0x81, 0xe1, 0x80, 0x71, 0x80, 0x39, 0x80, 0x1f, 0x80, 0x7, 0x80, 0x3, 0x80, 0x1, 0xe0, 0x1, 0xf0, 0x0, 0xb8, 0x0, 0x9c, 0x0, 0x8f, 0x0, 0x83, 0x80, 0xc1, 0xc0, 0xe0, 0xf0, 0xf9, 0xfe, 0x0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xe, Height: 0x11, XAdvance: 0xf, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0xfe, 0x7c, 0xe0, 0x63, 0x81, 0xf, 0x8, 0x1c, 0x40, 0x71, 0x1, 0xe8, 0x3, 0xc0, 0xe, 0x0, 0x38, 0x1, 0xe0, 0x7, 0x80, 0x1c, 0x0, 0x70, 0x3, 0xc0, 0xf, 0x0, 0xff, 0x0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xf, Height: 0x11, XAdvance: 0xd, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x1f, 0xfe, 0x38, 0x78, 0x60, 0xf1, 0x83, 0xc2, 0xf, 0x0, 0x1e, 0x0, 0x78, 0x1, 0xe0, 0x7, 0xc0, 0xf, 0x0, 0x3c, 0x0, 0xf8, 0x1, 0xe0, 0x47, 0x81, 0x1f, 0x6, 0x3c, 0x3c, 0xff, 0xf0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0xa, Height: 0x14, XAdvance: 0x8, XOffset: -1, YOffset: -15, Bitmaps: []uint8{0x7, 0xc1, 0x80, 0xe0, 0x30, 0xc, 0x3, 0x1, 0xc0, 0x60, 0x18, 0x6, 0x3, 0x80, 0xc0, 0x30, 0xc, 0x7, 0x1, 0xc0, 0x60, 0x18, 0xe, 0x3, 0xe0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x6, Height: 0x10, XAdvance: 0xa, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0xc3, 0x6, 0x18, 0x61, 0x83, 0xc, 0x30, 0xc1, 0x86, 0x18, 0x60, 0xc3}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x9, Height: 0x14, XAdvance: 0x8, XOffset: -1, YOffset: -15, Bitmaps: []uint8{0xf, 0x81, 0xc0, 0xe0, 0x60, 0x30, 0x18, 0x1c, 0xc, 0x6, 0x3, 0x3, 0x81, 0x80, 0xc0, 0x60, 0x70, 0x38, 0x18, 0xc, 0xe, 0x1f, 0x0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0xa, Height: 0x9, XAdvance: 0xe, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0xc, 0x7, 0x81, 0xe0, 0xdc, 0x33, 0x18, 0xc6, 0x1b, 0x6, 0xc0, 0xc0}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xc, Height: 0x1, XAdvance: 0xc, XOffset: 0, YOffset: 4, Bitmaps: []uint8{0xff, 0xf0}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x5, Height: 0x4, XAdvance: 0x8, XOffset: 2, YOffset: -15, Bitmaps: []uint8{0xc7, 0xc, 0x30}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0xc, Height: 0xc, XAdvance: 0xc, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x7, 0x70, 0xce, 0x1c, 0xe3, 0x8e, 0x70, 0xc7, 0xc, 0x71, 0xce, 0x1c, 0xe1, 0x8e, 0x79, 0xe9, 0xa7, 0x1c}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xb, Height: 0x12, XAdvance: 0xc, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x2, 0x7, 0xc0, 0x38, 0x6, 0x1, 0xc0, 0x38, 0x6, 0x71, 0xf7, 0x38, 0xe7, 0x1c, 0xc3, 0xb8, 0x77, 0x1c, 0xe3, 0xb8, 0xe7, 0x18, 0xe6, 0xf, 0x80}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x9, Height: 0xc, XAdvance: 0xa, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x7, 0xc, 0xce, 0x66, 0x7, 0x3, 0x83, 0x81, 0xc0, 0xe0, 0x70, 0xbc, 0x87, 0x80}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xe, Height: 0x12, XAdvance: 0xc, XOffset: 0, YOffset: -16, Bitmaps: []uint8{0x0, 0x8, 0x3, 0xe0, 0x3, 0x80, 0xe, 0x0, 0x70, 0x1, 0xc0, 0x77, 0x3, 0x3c, 0x18, 0xe0, 0xe3, 0x87, 0xe, 0x1c, 0x70, 0x71, 0xc3, 0x87, 0xe, 0x3c, 0x38, 0xe8, 0xe5, 0xa1, 0xe7, 0x0}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x9, Height: 0xc, XAdvance: 0xa, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x7, 0xc, 0xce, 0x66, 0x37, 0x33, 0xbb, 0xb1, 0xe0, 0xe0, 0x70, 0xb8, 0x87, 0x80}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0xe, Height: 0x16, XAdvance: 0xc, XOffset: -2, YOffset: -16, Bitmaps: []uint8{0x0, 0x38, 0x1, 0xb0, 0xc, 0xc0, 0x30, 0x1, 0xc0, 0x7, 0x0, 0x7e, 0x0, 0xe0, 0x3, 0x80, 0xe, 0x0, 0x30, 0x1, 0xc0, 0x7, 0x0, 0x1c, 0x0, 0x70, 0x3, 0x80, 0xe, 0x0, 0x38, 0x0, 0xc0, 0x33, 0x0, 0xd8, 0x1, 0xc0, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xd, Height: 0x10, XAdvance: 0xc, XOffset: -1, YOffset: -10, Bitmaps: []uint8{0x3, 0x80, 0x73, 0xc7, 0x1c, 0x38, 0xe1, 0xcf, 0x6, 0x70, 0x1e, 0x1, 0x0, 0x1c, 0x0, 0xf8, 0x7, 0xf0, 0xc7, 0x8c, 0xc, 0x60, 0x63, 0x86, 0x7, 0xe0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xd, Height: 0x12, XAdvance: 0xd, XOffset: 0, YOffset: -16, Bitmaps: []uint8{0x1, 0x0, 0xf8, 0x1, 0x80, 0x1c, 0x0, 0xe0, 0x7, 0x0, 0x31, 0xc3, 0xbe, 0x1e, 0x70, 0xe3, 0x8f, 0x38, 0x71, 0xc3, 0x8e, 0x1c, 0xe1, 0xc7, 0xe, 0x3a, 0x71, 0xd3, 0xf, 0x0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x6, Height: 0x11, XAdvance: 0x7, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1c, 0x71, 0xc0, 0x0, 0x6f, 0x8e, 0x31, 0xc7, 0x18, 0x63, 0x8e, 0xbc, 0xe0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0xb, Height: 0x15, XAdvance: 0x8, XOffset: -2, YOffset: -15, Bitmaps: []uint8{0x0, 0xe0, 0x1c, 0x3, 0x80, 0x0, 0x0, 0xf, 0x80, 0x70, 0xe, 0x1, 0xc0, 0x70, 0xe, 0x1, 0xc0, 0x38, 0xe, 0x1, 0xc0, 0x38, 0x6, 0x1, 0xc3, 0x38, 0x6e, 0x7, 0x80}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xd, Height: 0x12, XAdvance: 0xc, XOffset: 0, YOffset: -16, Bitmaps: []uint8{0x1, 0x0, 0xf8, 0x1, 0xc0, 0x1c, 0x0, 0xe0, 0x7, 0x0, 0x33, 0xe3, 0x8c, 0x1c, 0xc0, 0xe4, 0x6, 0x40, 0x7e, 0x3, 0xf0, 0x1d, 0x81, 0xce, 0xe, 0x72, 0x71, 0xa3, 0x8e, 0x0}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x6, 0x7c, 0x70, 0xe1, 0xc3, 0xe, 0x1c, 0x38, 0x61, 0xc3, 0x87, 0xc, 0x38, 0x72, 0xe9, 0xe0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x12, Height: 0xc, XAdvance: 0x12, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x3c, 0x73, 0xc7, 0x7d, 0x71, 0xe7, 0x9c, 0xf1, 0xce, 0x3c, 0xf3, 0x8e, 0x39, 0xc3, 0x8e, 0x71, 0xc3, 0x1c, 0x71, 0xc7, 0x1c, 0x71, 0xd7, 0x1c, 0x7b, 0x8e, 0x1c}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x3c, 0xf1, 0xd7, 0x1e, 0x73, 0xce, 0x3c, 0xe3, 0x8e, 0x39, 0xc7, 0x9c, 0x71, 0xc7, 0x1d, 0x71, 0xee, 0x1c}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0xa, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xf, 0x6, 0x63, 0x9d, 0xc7, 0x71, 0xf8, 0x7e, 0x3f, 0x8e, 0xe3, 0xb9, 0xc6, 0x60, 0xf0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xe, Height: 0x10, XAdvance: 0xc, XOffset: -2, YOffset: -10, Bitmaps: []uint8{0xf, 0x38, 0x1f, 0x70, 0x71, 0xc1, 0xc7, 0xe, 0x1c, 0x38, 0xf0, 0xe3, 0x83, 0x8e, 0x1c, 0x70, 0x71, 0xc1, 0xce, 0x7, 0xe0, 0x38, 0x0, 0xe0, 0x3, 0x80, 0x3f, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xc, Height: 0x10, XAdvance: 0xc, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x7, 0x70, 0xce, 0x18, 0xe3, 0x8e, 0x70, 0xe7, 0x1c, 0xf1, 0xce, 0x1c, 0xe3, 0x8e, 0x38, 0xe7, 0x87, 0xb0, 0x7, 0x0, 0x70, 0xf, 0x3, 0xf8}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0xa, Height: 0xb, XAdvance: 0xa, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0xd, 0xdf, 0x71, 0xac, 0xf0, 0x38, 0xe, 0x3, 0x81, 0xc0, 0x70, 0x1c, 0xe, 0x0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x1d, 0x99, 0x8c, 0x46, 0x23, 0x80, 0xe0, 0x70, 0x1c, 0x6, 0x23, 0x19, 0x17, 0x0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x7, Height: 0xf, XAdvance: 0x7, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xc, 0x10, 0xe3, 0xf3, 0x86, 0x1c, 0x38, 0x71, 0xc3, 0x87, 0xe, 0x9e, 0x38, 0x0}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xf8, 0xe3, 0x8e, 0x38, 0xc3, 0x9c, 0x71, 0xc7, 0x18, 0x71, 0x87, 0x38, 0xe3, 0x8e, 0xfa, 0xf3, 0xae, 0x3c}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0xa, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xf0, 0xdc, 0x33, 0xc, 0xc2, 0x31, 0x8c, 0xc3, 0x60, 0xf0, 0x38, 0xc, 0x2, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xf, Height: 0xb, XAdvance: 0x10, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xe0, 0x86, 0xe3, 0xc, 0xc6, 0x19, 0x9c, 0x23, 0x78, 0xc7, 0xf9, 0xe, 0x74, 0x1c, 0xf0, 0x31, 0xc0, 0x43, 0x0, 0x84, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0xd, Height: 0xc, XAdvance: 0xb, XOffset: -1, YOffset: -10, Bitmaps: []uint8{0xe, 0x31, 0xf3, 0x83, 0xa0, 0xe, 0x0, 0x70, 0x3, 0x80, 0x1c, 0x0, 0xe0, 0xb, 0x2, 0x5d, 0x3c, 0xf1, 0xc3, 0x0}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0xb, Height: 0x10, XAdvance: 0xa, XOffset: -1, YOffset: -10, Bitmaps: []uint8{0x4, 0x67, 0x8c, 0x79, 0x87, 0x10, 0xe2, 0x1c, 0x81, 0x90, 0x3a, 0x7, 0x80, 0xf0, 0x1c, 0x3, 0x0, 0x40, 0x8, 0x32, 0x7, 0x80}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x3f, 0xcf, 0xe6, 0x30, 0x8, 0x4, 0x2, 0x1, 0x0, 0xc0, 0x30, 0x1e, 0xf, 0x98, 0x76, 0x7, 0x0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0xb, Height: 0x15, XAdvance: 0x8, XOffset: 0, YOffset: -16, Bitmaps: []uint8{0x1, 0xe0, 0x70, 0x1c, 0x3, 0x80, 0x60, 0x1c, 0x3, 0x80, 0x60, 0xc, 0x3, 0x80, 0xf0, 0x3c, 0x7, 0x0, 0x40, 0xc, 0x1, 0x80, 0x70, 0xe, 0x1, 0xc0, 0x30, 0x3, 0x80}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x2, Height: 0x10, XAdvance: 0x6, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xff}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0xa, Height: 0x15, XAdvance: 0x8, XOffset: -3, YOffset: -16, Bitmaps: []uint8{0x7, 0x0, 0xe0, 0x18, 0x6, 0x1, 0x80, 0xe0, 0x38, 0xc, 0x3, 0x0, 0xc0, 0x10, 0x1f, 0x7, 0x3, 0x80, 0xe0, 0x30, 0xc, 0x7, 0x1, 0x80, 0xe0, 0xe0, 0x0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0xb, Height: 0x4, XAdvance: 0xe, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x38, 0xf, 0xcd, 0x1f, 0x80, 0xe0}},
	},

	YAdvance: 0x1d,
}
