package freemono

import (
	"tinygo.org/x/tinyfont"
)

var Oblique18pt7b = tinyfont.Font{
	BBox: [4]int8{26, 30, -1, -22},
	Glyphs: []tinyfont.Glypher{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x15, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x7, Height: 0x16, XAdvance: 0x15, XOffset: 9, YOffset: -21, Bitmaps: []uint8{0x0, 0x1c, 0x38, 0x70, 0xc1, 0x83, 0x6, 0x18, 0x30, 0x60, 0xc1, 0x2, 0x4, 0x0, 0x0, 0x1, 0xc7, 0x8f, 0x1c, 0x0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0xd, Height: 0xa, XAdvance: 0x15, XOffset: 7, YOffset: -20, Bitmaps: []uint8{0x78, 0x7b, 0xc3, 0xfc, 0x3d, 0xe1, 0xef, 0xf, 0x70, 0x73, 0x83, 0x98, 0x18, 0xc0, 0xc6, 0x6, 0x0}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xf, Height: 0x18, XAdvance: 0x15, XOffset: 5, YOffset: -21, Bitmaps: []uint8{0x0, 0x8c, 0x1, 0x18, 0x6, 0x20, 0x8, 0x40, 0x11, 0x80, 0x62, 0x0, 0xc4, 0x1, 0x18, 0x2, 0x30, 0x7f, 0xfc, 0x10, 0x80, 0x23, 0x0, 0xc4, 0x1, 0x88, 0x3f, 0xff, 0x4, 0x60, 0x18, 0x80, 0x21, 0x0, 0x46, 0x1, 0x88, 0x3, 0x10, 0x4, 0x60, 0x8, 0xc0, 0x31, 0x0}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x10, Height: 0x1a, XAdvance: 0x15, XOffset: 4, YOffset: -22, Bitmaps: []uint8{0x0, 0x30, 0x0, 0x20, 0x0, 0x20, 0x0, 0xf9, 0x3, 0xf, 0x6, 0x3, 0x4, 0x3, 0x8, 0x0, 0x8, 0x0, 0x8, 0x0, 0x4, 0x0, 0x3, 0xc0, 0x0, 0x78, 0x0, 0xc, 0x0, 0x4, 0x0, 0x4, 0x40, 0x4, 0x40, 0x8, 0x40, 0x18, 0xf0, 0x60, 0x9f, 0x80, 0x2, 0x0, 0x6, 0x0, 0x4, 0x0, 0x4, 0x0, 0x4, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x10, Height: 0x15, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x3, 0xc0, 0xc, 0x60, 0x8, 0x20, 0x10, 0x20, 0x10, 0x20, 0x10, 0x40, 0x18, 0x80, 0xf, 0x0, 0x0, 0xf, 0x0, 0x78, 0x7, 0xc0, 0x3c, 0x0, 0xe0, 0x0, 0x1, 0xe0, 0x2, 0x18, 0x4, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x10, 0xc, 0x20, 0x7, 0xc0}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xd, Height: 0x12, XAdvance: 0x15, XOffset: 5, YOffset: -17, Bitmaps: []uint8{0x1, 0xf0, 0x11, 0x81, 0x0, 0x10, 0x0, 0x80, 0x4, 0x0, 0x20, 0x1, 0x80, 0x4, 0x0, 0xf0, 0x9, 0x86, 0x84, 0x48, 0x32, 0x40, 0xa2, 0x7, 0x10, 0x30, 0x43, 0x81, 0xe7, 0x80}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x5, Height: 0xa, XAdvance: 0x15, XOffset: 12, YOffset: -20, Bitmaps: []uint8{0x7b, 0xfd, 0xef, 0x73, 0x98, 0xc6, 0x0}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x8, Height: 0x19, XAdvance: 0x15, XOffset: 12, YOffset: -20, Bitmaps: []uint8{0x1, 0x2, 0x6, 0xc, 0xc, 0x18, 0x10, 0x30, 0x30, 0x60, 0x60, 0x60, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0x40, 0x60, 0x60, 0x20}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x8, Height: 0x19, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x4, 0x6, 0x6, 0x2, 0x2, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x6, 0x6, 0x6, 0xc, 0xc, 0xc, 0x18, 0x10, 0x30, 0x60, 0x40, 0xc0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0xe, Height: 0xb, XAdvance: 0x15, XOffset: 7, YOffset: -19, Bitmaps: []uint8{0x1, 0x0, 0x4, 0x0, 0x10, 0x0, 0xc6, 0xe3, 0xf8, 0x7e, 0x0, 0x70, 0x3, 0x40, 0x19, 0x80, 0xc2, 0x6, 0xc, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0xf, Height: 0x11, XAdvance: 0x15, XOffset: 5, YOffset: -17, Bitmaps: []uint8{0x0, 0xc0, 0x1, 0x0, 0x2, 0x0, 0x4, 0x0, 0x8, 0x0, 0x20, 0x0, 0x40, 0x0, 0x80, 0xff, 0xfe, 0x2, 0x0, 0x8, 0x0, 0x10, 0x0, 0x20, 0x0, 0x40, 0x0, 0x80, 0x2, 0x0, 0x4, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x9, Height: 0xa, XAdvance: 0x15, XOffset: 4, YOffset: -4, Bitmaps: []uint8{0xf, 0x87, 0x87, 0x83, 0x83, 0xc1, 0xc1, 0xc0, 0xc0, 0xe0, 0x60, 0x0}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x10, Height: 0x1, XAdvance: 0x15, XOffset: 5, YOffset: -9, Bitmaps: []uint8{0xff, 0xff}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x5, Height: 0x5, XAdvance: 0x15, XOffset: 8, YOffset: -4, Bitmaps: []uint8{0x77, 0xff, 0xf7, 0x0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x13, Height: 0x1a, XAdvance: 0x15, XOffset: 3, YOffset: -22, Bitmaps: []uint8{0x0, 0x0, 0x60, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0xc0, 0x0, 0x30, 0x0, 0x4, 0x0, 0x1, 0x80, 0x0, 0x60, 0x0, 0x8, 0x0, 0x3, 0x0, 0x0, 0xc0, 0x0, 0x10, 0x0, 0x6, 0x0, 0x1, 0x80, 0x0, 0x20, 0x0, 0xc, 0x0, 0x3, 0x0, 0x0, 0x40, 0x0, 0x18, 0x0, 0x6, 0x0, 0x0, 0x80, 0x0, 0x20, 0x0, 0xc, 0x0, 0x3, 0x0, 0x0, 0x40, 0x0, 0x8, 0x0, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0xe, Height: 0x15, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x1, 0xf0, 0x18, 0x60, 0x80, 0x86, 0x1, 0x10, 0x4, 0x80, 0x12, 0x0, 0x50, 0x1, 0x40, 0xd, 0x0, 0x24, 0x0, 0xa0, 0x2, 0x80, 0x1a, 0x0, 0x48, 0x1, 0x20, 0xc, 0x80, 0x22, 0x1, 0x84, 0xc, 0x18, 0x60, 0x3e, 0x0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0xd, Height: 0x15, XAdvance: 0x15, XOffset: 4, YOffset: -20, Bitmaps: []uint8{0x0, 0x60, 0x7, 0x0, 0x68, 0x6, 0x40, 0xe4, 0x4, 0x20, 0x1, 0x0, 0x8, 0x0, 0x40, 0x4, 0x0, 0x20, 0x1, 0x0, 0x8, 0x0, 0x80, 0x4, 0x0, 0x20, 0x1, 0x0, 0x8, 0x0, 0x80, 0x4, 0xf, 0xff, 0x80}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x11, Height: 0x15, XAdvance: 0x15, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0x0, 0x3c, 0x0, 0x61, 0x80, 0x40, 0x40, 0x40, 0x10, 0x60, 0x8, 0x0, 0x4, 0x0, 0x2, 0x0, 0x2, 0x0, 0x3, 0x0, 0x3, 0x0, 0x7, 0x0, 0x7, 0x0, 0x6, 0x0, 0x6, 0x0, 0xe, 0x0, 0xe, 0x0, 0xc, 0x0, 0xc, 0x0, 0x1c, 0x1, 0x1c, 0x0, 0x8f, 0xff, 0xc0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x10, Height: 0x15, XAdvance: 0x15, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0x0, 0xfc, 0x3, 0x6, 0x6, 0x3, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x2, 0x0, 0x2, 0x0, 0xc, 0x0, 0xf0, 0x0, 0x18, 0x0, 0x4, 0x0, 0x2, 0x0, 0x2, 0x0, 0x2, 0x0, 0x2, 0x0, 0x4, 0x0, 0x4, 0x40, 0x18, 0x70, 0x30, 0xf, 0xc0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0xe, Height: 0x15, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x0, 0x1c, 0x0, 0xd0, 0x6, 0x80, 0x32, 0x0, 0x88, 0x4, 0x20, 0x30, 0x81, 0x84, 0x4, 0x10, 0x20, 0x41, 0x81, 0xc, 0x8, 0x60, 0x21, 0x0, 0x8f, 0xff, 0x80, 0x18, 0x0, 0x40, 0x1, 0x0, 0x4, 0x0, 0x10, 0x7, 0xe0}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x11, Height: 0x15, XAdvance: 0x15, XOffset: 4, YOffset: -20, Bitmaps: []uint8{0x3, 0xff, 0x3, 0x0, 0x1, 0x80, 0x0, 0x80, 0x0, 0x40, 0x0, 0x20, 0x0, 0x30, 0x0, 0x1b, 0xe0, 0xe, 0xc, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x40, 0x0, 0x20, 0x0, 0x10, 0x0, 0x8, 0x0, 0x8, 0x0, 0x4, 0x60, 0x4, 0x18, 0x4, 0x6, 0xc, 0x0, 0xf8, 0x0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x10, Height: 0x15, XAdvance: 0x15, XOffset: 6, YOffset: -20, Bitmaps: []uint8{0x0, 0x3f, 0x0, 0xc0, 0x3, 0x0, 0x4, 0x0, 0x8, 0x0, 0x10, 0x0, 0x30, 0x0, 0x20, 0x0, 0x40, 0x0, 0x43, 0xe0, 0x4c, 0x30, 0xb0, 0x18, 0xe0, 0x8, 0xc0, 0x8, 0x80, 0x8, 0x80, 0x8, 0x80, 0x10, 0xc0, 0x10, 0x40, 0x20, 0x20, 0xc0, 0x1f, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0xd, Height: 0x15, XAdvance: 0x15, XOffset: 8, YOffset: -20, Bitmaps: []uint8{0xff, 0xfc, 0x0, 0xe0, 0x4, 0x0, 0x60, 0x2, 0x0, 0x30, 0x1, 0x0, 0x18, 0x0, 0x80, 0xc, 0x0, 0x40, 0x6, 0x0, 0x20, 0x3, 0x0, 0x10, 0x1, 0x80, 0x8, 0x0, 0xc0, 0x4, 0x0, 0x60, 0x2, 0x0, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0xf, Height: 0x15, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x0, 0xf0, 0x6, 0x18, 0x10, 0x18, 0x40, 0x11, 0x0, 0x22, 0x0, 0x44, 0x0, 0x88, 0x2, 0x18, 0x8, 0x18, 0x60, 0x1f, 0x80, 0xc1, 0x82, 0x1, 0x88, 0x1, 0x20, 0x2, 0x40, 0x4, 0x80, 0x9, 0x0, 0x23, 0x0, 0x83, 0x6, 0x1, 0xf0, 0x0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0xf, Height: 0x15, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x0, 0xf0, 0x6, 0x18, 0x10, 0x10, 0x40, 0x30, 0x80, 0x22, 0x0, 0x44, 0x0, 0x88, 0x3, 0x10, 0xe, 0x30, 0x34, 0x30, 0xd0, 0x3e, 0x20, 0x0, 0x40, 0x1, 0x0, 0x2, 0x0, 0x8, 0x0, 0x20, 0x0, 0xc0, 0x2, 0x0, 0x18, 0xf, 0xc0, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x7, Height: 0xf, XAdvance: 0x15, XOffset: 8, YOffset: -14, Bitmaps: []uint8{0x1c, 0x7c, 0xf9, 0xf1, 0xc0, 0x0, 0x0, 0x0, 0x1, 0xc7, 0xcf, 0x9f, 0x1c, 0x0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0xb, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1, 0xc0, 0x7c, 0xf, 0x81, 0xf0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3e, 0x7, 0x81, 0xe0, 0x3c, 0xf, 0x1, 0xc0, 0x70, 0xe, 0x3, 0x80, 0x60, 0x0}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x11, Height: 0x10, XAdvance: 0x15, XOffset: 5, YOffset: -17, Bitmaps: []uint8{0x0, 0x1, 0x80, 0x3, 0x80, 0x7, 0x0, 0xe, 0x0, 0x1c, 0x0, 0x38, 0x0, 0x70, 0x0, 0xe0, 0x0, 0xe0, 0x0, 0x1c, 0x0, 0x7, 0x0, 0x0, 0xe0, 0x0, 0x38, 0x0, 0x7, 0x0, 0x0, 0xe0, 0x0, 0x38}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x13, Height: 0x6, XAdvance: 0x15, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0x7f, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xff, 0xff, 0x80}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x12, Height: 0x10, XAdvance: 0x15, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x18, 0x0, 0x3, 0x80, 0x0, 0x38, 0x0, 0x7, 0x0, 0x0, 0x70, 0x0, 0xe, 0x0, 0x0, 0xe0, 0x0, 0xe, 0x0, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x3, 0x80, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0xc, Height: 0x14, XAdvance: 0x15, XOffset: 8, YOffset: -19, Bitmaps: []uint8{0x1f, 0xce, 0x6, 0x80, 0x38, 0x1, 0x80, 0x10, 0x1, 0x0, 0x20, 0x4, 0x1, 0x80, 0xf0, 0x18, 0x1, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x78, 0xf, 0x80, 0xf8, 0x7, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0xf, Height: 0x17, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x1, 0xf0, 0xc, 0x30, 0x30, 0x30, 0x40, 0x21, 0x0, 0x44, 0x0, 0x88, 0x1, 0x10, 0x1e, 0x40, 0xc4, 0x86, 0x11, 0x8, 0x22, 0x20, 0x48, 0x40, 0x90, 0x82, 0x21, 0x84, 0x40, 0xfc, 0x80, 0x1, 0x0, 0x2, 0x0, 0x4, 0x0, 0x4, 0x0, 0xc, 0x18, 0x7, 0xc0, 0x0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x15, Height: 0x14, XAdvance: 0x15, XOffset: 0, YOffset: -19, Bitmaps: []uint8{0x1, 0xfe, 0x0, 0x0, 0x68, 0x0, 0x6, 0x40, 0x0, 0x32, 0x0, 0x3, 0x10, 0x0, 0x10, 0x80, 0x1, 0x84, 0x0, 0x18, 0x10, 0x0, 0xc0, 0x80, 0xc, 0x4, 0x0, 0x60, 0x20, 0x6, 0x1, 0x0, 0x3f, 0xfc, 0x2, 0x0, 0x20, 0x10, 0x1, 0x1, 0x0, 0x8, 0x8, 0x0, 0x40, 0x80, 0x2, 0xc, 0x0, 0x9, 0xfc, 0x7, 0xf0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xff, 0x0, 0x40, 0x60, 0x20, 0xc, 0x8, 0x1, 0x2, 0x0, 0x40, 0x80, 0x10, 0x40, 0x8, 0x10, 0x6, 0x4, 0x3, 0x1, 0xff, 0x80, 0x40, 0x38, 0x20, 0x2, 0x8, 0x0, 0x42, 0x0, 0x10, 0x80, 0x4, 0x40, 0x1, 0x10, 0x0, 0x84, 0x0, 0x41, 0x0, 0x23, 0xff, 0xf0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x0, 0xfc, 0x40, 0xc1, 0xf0, 0xc0, 0x1c, 0x60, 0x6, 0x10, 0x0, 0x88, 0x0, 0x24, 0x0, 0x1, 0x0, 0x0, 0x40, 0x0, 0x30, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x20, 0x0, 0x8, 0x0, 0x3, 0x0, 0x0, 0x40, 0x6, 0x8, 0x3, 0x1, 0x83, 0x80, 0x3f, 0x0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xfe, 0x0, 0x80, 0xc0, 0x20, 0x18, 0x10, 0x2, 0x4, 0x0, 0x41, 0x0, 0x10, 0x40, 0x4, 0x20, 0x1, 0x8, 0x0, 0x42, 0x0, 0x10, 0x80, 0x8, 0x20, 0x2, 0x10, 0x0, 0x84, 0x0, 0x21, 0x0, 0x10, 0x40, 0x8, 0x20, 0x6, 0x8, 0x3, 0x2, 0x1, 0x83, 0xff, 0x80}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x14, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xff, 0xe0, 0x10, 0x2, 0x2, 0x0, 0x60, 0x20, 0x6, 0x2, 0x0, 0x60, 0x20, 0x0, 0x4, 0x0, 0x0, 0x40, 0x80, 0x4, 0x10, 0x0, 0x7f, 0x0, 0x4, 0x10, 0x0, 0x81, 0x0, 0x8, 0x0, 0x0, 0x80, 0x0, 0x8, 0x0, 0x81, 0x0, 0x8, 0x10, 0x0, 0x81, 0x0, 0x18, 0x10, 0x1, 0x8f, 0xff, 0xf0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x14, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xff, 0xf0, 0x10, 0x3, 0x2, 0x0, 0x30, 0x20, 0x3, 0x2, 0x0, 0x20, 0x20, 0x0, 0x4, 0x0, 0x0, 0x40, 0x80, 0x4, 0x10, 0x0, 0x7f, 0x0, 0x4, 0x10, 0x0, 0x81, 0x0, 0x8, 0x0, 0x0, 0x80, 0x0, 0x8, 0x0, 0x1, 0x0, 0x0, 0x10, 0x0, 0x1, 0x0, 0x0, 0x10, 0x0, 0xf, 0xf8, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x0, 0xfe, 0x40, 0xc0, 0xf0, 0x40, 0x1c, 0x20, 0x3, 0x10, 0x0, 0x88, 0x0, 0x2, 0x0, 0x1, 0x0, 0x0, 0x40, 0x0, 0x10, 0x0, 0x8, 0x0, 0x2, 0x1, 0xfe, 0x80, 0x2, 0x20, 0x0, 0x88, 0x0, 0x22, 0x0, 0x8, 0x40, 0x4, 0x18, 0x1, 0x3, 0x81, 0xc0, 0x3f, 0x80}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x15, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0x7, 0xe1, 0xf8, 0x8, 0x2, 0x0, 0x80, 0x10, 0x4, 0x0, 0x80, 0x20, 0x4, 0x1, 0x0, 0x20, 0x18, 0x2, 0x0, 0x80, 0x10, 0x4, 0x0, 0x80, 0x3f, 0xfc, 0x1, 0x0, 0x60, 0x10, 0x2, 0x0, 0x80, 0x10, 0x4, 0x0, 0x80, 0x20, 0x4, 0x2, 0x0, 0x40, 0x10, 0x2, 0x0, 0x80, 0x10, 0x4, 0x0, 0x81, 0xf8, 0x3f, 0x0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x11, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0xf, 0xff, 0x80, 0x10, 0x0, 0x8, 0x0, 0x8, 0x0, 0x4, 0x0, 0x2, 0x0, 0x1, 0x0, 0x0, 0x80, 0x0, 0x80, 0x0, 0x40, 0x0, 0x20, 0x0, 0x10, 0x0, 0x8, 0x0, 0x8, 0x0, 0x4, 0x0, 0x2, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x80, 0x1f, 0xff, 0x0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x14, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x0, 0xff, 0xf0, 0x0, 0x20, 0x0, 0x2, 0x0, 0x0, 0x20, 0x0, 0x2, 0x0, 0x0, 0x20, 0x0, 0x4, 0x0, 0x0, 0x40, 0x0, 0x4, 0x0, 0x0, 0x40, 0x0, 0xc, 0x4, 0x0, 0x80, 0x40, 0x8, 0x8, 0x0, 0x80, 0x80, 0x8, 0x8, 0x1, 0x0, 0x80, 0x10, 0xc, 0x2, 0x0, 0x60, 0xc0, 0x1, 0xf0, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x15, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xe1, 0xf8, 0x8, 0x3, 0x0, 0x80, 0x60, 0x4, 0x6, 0x0, 0x20, 0x60, 0x1, 0x6, 0x0, 0x10, 0xc0, 0x0, 0x8c, 0x0, 0x4, 0xc0, 0x0, 0x2f, 0x80, 0x1, 0x8e, 0x0, 0x18, 0x30, 0x0, 0x80, 0xc0, 0x4, 0x6, 0x0, 0x20, 0x10, 0x2, 0x0, 0xc0, 0x10, 0x6, 0x0, 0x80, 0x30, 0x4, 0x0, 0x81, 0xfc, 0x7, 0x80}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0x7, 0xfc, 0x0, 0x10, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x20, 0x0, 0x8, 0x0, 0x4, 0x0, 0x1, 0x0, 0x0, 0x40, 0x0, 0x10, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0x80, 0x10, 0x20, 0x4, 0x8, 0x1, 0x4, 0x0, 0x81, 0x0, 0x20, 0x40, 0xb, 0xff, 0xfe}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x18, Height: 0x14, XAdvance: 0x15, XOffset: 1, YOffset: -19, Bitmaps: []uint8{0xf, 0x0, 0x1e, 0x3, 0x0, 0x38, 0x5, 0x0, 0x68, 0x4, 0x80, 0x68, 0x4, 0x80, 0xc8, 0x4, 0x80, 0x90, 0x4, 0x81, 0x90, 0x8, 0x43, 0x10, 0x8, 0x42, 0x10, 0x8, 0x46, 0x10, 0x8, 0x4c, 0x20, 0x10, 0x2c, 0x20, 0x10, 0x38, 0x20, 0x10, 0x30, 0x20, 0x10, 0x0, 0x40, 0x10, 0x0, 0x40, 0x20, 0x0, 0x40, 0x20, 0x0, 0x40, 0x20, 0x0, 0x40, 0xfc, 0x7, 0xe0}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x16, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0x1f, 0x1, 0xfc, 0xc, 0x0, 0x80, 0x78, 0x2, 0x1, 0xe0, 0x18, 0x4, 0x80, 0x60, 0x13, 0x1, 0x0, 0x4c, 0x4, 0x3, 0x18, 0x10, 0xc, 0x60, 0xc0, 0x20, 0x83, 0x0, 0x83, 0x8, 0x6, 0xc, 0x20, 0x18, 0x18, 0x80, 0x40, 0x66, 0x1, 0x0, 0x98, 0x4, 0x3, 0x40, 0x30, 0xd, 0x0, 0xc0, 0x14, 0x2, 0x0, 0x70, 0x3f, 0x80, 0xc0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x11, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x0, 0xf8, 0x1, 0x83, 0x1, 0x0, 0xc1, 0x0, 0x21, 0x0, 0x19, 0x0, 0x4, 0x80, 0x2, 0x80, 0x1, 0x40, 0x0, 0xc0, 0x0, 0x60, 0x0, 0x30, 0x0, 0x28, 0x0, 0x14, 0x0, 0x12, 0x0, 0x9, 0x80, 0x8, 0x40, 0x8, 0x30, 0x8, 0xc, 0x18, 0x1, 0xf0, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xfe, 0x0, 0x40, 0x60, 0x20, 0xc, 0x8, 0x1, 0x2, 0x0, 0x40, 0x80, 0x10, 0x40, 0x4, 0x10, 0x2, 0x4, 0x1, 0x1, 0x1, 0x80, 0x7f, 0x80, 0x20, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x40, 0x0, 0x10, 0x0, 0x4, 0x0, 0x1, 0x0, 0x3, 0xfe, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x11, Height: 0x18, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x0, 0xf8, 0x1, 0x83, 0x1, 0x0, 0xc1, 0x0, 0x21, 0x0, 0x19, 0x0, 0x5, 0x0, 0x2, 0x80, 0x1, 0x40, 0x0, 0xc0, 0x0, 0x60, 0x0, 0x30, 0x0, 0x28, 0x0, 0x14, 0x0, 0x12, 0x0, 0x9, 0x80, 0x8, 0x40, 0x8, 0x30, 0x8, 0xc, 0x18, 0x3, 0xf0, 0x0, 0xc0, 0x1, 0xc0, 0x1, 0xfe, 0x18, 0xc0, 0xf0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xfe, 0x0, 0x40, 0x60, 0x20, 0xc, 0x8, 0x1, 0x2, 0x0, 0x40, 0x80, 0x10, 0x40, 0x4, 0x10, 0x2, 0x4, 0x1, 0x1, 0x1, 0x80, 0x7f, 0x80, 0x20, 0x60, 0x8, 0xc, 0x2, 0x3, 0x80, 0x80, 0x60, 0x40, 0x18, 0x10, 0x3, 0x4, 0x0, 0xc1, 0x0, 0x1b, 0xf8, 0x7}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 3, YOffset: -19, Bitmaps: []uint8{0x0, 0x7e, 0x40, 0x60, 0xf0, 0x20, 0x1c, 0x10, 0x2, 0x8, 0x0, 0x82, 0x0, 0x0, 0x80, 0x0, 0x30, 0x0, 0x6, 0x0, 0x0, 0xf8, 0x0, 0x3, 0xc0, 0x0, 0x18, 0x0, 0x1, 0x0, 0x0, 0x44, 0x0, 0x11, 0x0, 0x4, 0x40, 0x2, 0x38, 0x1, 0xb, 0x81, 0x82, 0x3f, 0x80}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0x11, Height: 0x14, XAdvance: 0x15, XOffset: 5, YOffset: -19, Bitmaps: []uint8{0x3f, 0xff, 0xa0, 0x20, 0x50, 0x10, 0x28, 0x8, 0x24, 0x8, 0x10, 0x4, 0x0, 0x2, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x80, 0x0, 0x40, 0x0, 0x20, 0x0, 0x10, 0x0, 0x10, 0x0, 0x8, 0x0, 0x4, 0x0, 0x2, 0x0, 0x2, 0x0, 0x1, 0x0, 0x1f, 0xfc, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 5, YOffset: -19, Bitmaps: []uint8{0x7e, 0xf, 0xc4, 0x0, 0x42, 0x0, 0x10, 0x80, 0x8, 0x20, 0x2, 0x8, 0x0, 0x82, 0x0, 0x21, 0x0, 0x8, 0x40, 0x4, 0x10, 0x1, 0x4, 0x0, 0x41, 0x0, 0x10, 0x80, 0xc, 0x20, 0x2, 0x8, 0x0, 0x82, 0x0, 0x60, 0x80, 0x10, 0x10, 0x8, 0x6, 0xc, 0x0, 0x7c, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0x15, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0xfe, 0x3, 0xf9, 0x80, 0x2, 0xc, 0x0, 0x30, 0x20, 0x1, 0x1, 0x0, 0x10, 0x8, 0x1, 0x80, 0x60, 0x8, 0x3, 0x0, 0xc0, 0x18, 0x4, 0x0, 0x40, 0x60, 0x2, 0x6, 0x0, 0x10, 0x20, 0x0, 0xc3, 0x0, 0x6, 0x10, 0x0, 0x31, 0x80, 0x0, 0x88, 0x0, 0x4, 0x80, 0x0, 0x2c, 0x0, 0x1, 0xc0, 0x0, 0xe, 0x0, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x14, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0x7f, 0x7, 0xf2, 0x0, 0x4, 0x20, 0x0, 0xc2, 0x0, 0x8, 0x20, 0xc0, 0x82, 0xc, 0x18, 0x21, 0xa1, 0x2, 0x1a, 0x10, 0x23, 0x23, 0x4, 0x32, 0x30, 0x46, 0x22, 0x4, 0x62, 0x60, 0x4c, 0x26, 0x4, 0xc2, 0x40, 0x58, 0x24, 0x5, 0x82, 0xc0, 0x70, 0x28, 0x7, 0x2, 0x80, 0xe0, 0x38, 0xe, 0x3, 0x0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x15, Height: 0x14, XAdvance: 0x15, XOffset: 2, YOffset: -19, Bitmaps: []uint8{0xf, 0xc1, 0xf8, 0x30, 0x3, 0x0, 0xc0, 0x30, 0x6, 0x3, 0x0, 0x18, 0x10, 0x0, 0xc1, 0x0, 0x3, 0x18, 0x0, 0x9, 0x80, 0x0, 0x78, 0x0, 0x1, 0x80, 0x0, 0x1c, 0x0, 0x1, 0xa0, 0x0, 0x19, 0x80, 0x1, 0x84, 0x0, 0x18, 0x30, 0x1, 0x80, 0xc0, 0x8, 0x6, 0x0, 0x80, 0x18, 0x8, 0x0, 0xc1, 0xf8, 0x3f, 0x80}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0x12, Height: 0x14, XAdvance: 0x15, XOffset: 5, YOffset: -19, Bitmaps: []uint8{0x7e, 0xf, 0xc4, 0x0, 0xc1, 0x80, 0x60, 0x20, 0x30, 0xc, 0x8, 0x3, 0x4, 0x0, 0x43, 0x0, 0x19, 0x80, 0x2, 0xc0, 0x0, 0xe0, 0x0, 0x10, 0x0, 0x4, 0x0, 0x1, 0x0, 0x0, 0x80, 0x0, 0x20, 0x0, 0x8, 0x0, 0x2, 0x0, 0x1, 0x0, 0x0, 0x40, 0x3, 0xff, 0x80}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x11, Height: 0x14, XAdvance: 0x15, XOffset: 4, YOffset: -19, Bitmaps: []uint8{0xf, 0xff, 0x86, 0x0, 0x82, 0x0, 0x81, 0x0, 0xc1, 0x80, 0xc0, 0xc0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0x40, 0x0, 0x40, 0x0, 0x60, 0x0, 0x60, 0x0, 0x60, 0x0, 0x60, 0x10, 0x60, 0x18, 0x20, 0x8, 0x20, 0x4, 0x20, 0x2, 0x30, 0x3, 0x1f, 0xff, 0x80}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0xb, Height: 0x19, XAdvance: 0x15, XOffset: 9, YOffset: -20, Bitmaps: []uint8{0x7, 0xe0, 0x80, 0x10, 0x2, 0x0, 0xc0, 0x18, 0x2, 0x0, 0x40, 0x18, 0x3, 0x0, 0x40, 0x8, 0x1, 0x0, 0x60, 0xc, 0x1, 0x0, 0x20, 0x4, 0x1, 0x80, 0x30, 0x4, 0x0, 0x80, 0x10, 0x6, 0x0, 0xfc, 0x0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x8, Height: 0x1b, XAdvance: 0x15, XOffset: 9, YOffset: -22, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x40, 0x40, 0x40, 0x20, 0x20, 0x20, 0x20, 0x10, 0x10, 0x10, 0x10, 0x8, 0x8, 0x8, 0x8, 0x4, 0x4, 0x4, 0x4, 0x2, 0x2, 0x2, 0x2, 0x0}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0xb, Height: 0x19, XAdvance: 0x15, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0x7, 0xe0, 0xc, 0x1, 0x0, 0x20, 0x4, 0x1, 0x80, 0x30, 0x4, 0x0, 0x80, 0x30, 0x6, 0x0, 0x80, 0x10, 0x2, 0x0, 0xc0, 0x18, 0x2, 0x0, 0x40, 0x18, 0x3, 0x0, 0x40, 0x8, 0x3, 0x0, 0x60, 0xf8, 0x0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0xd, Height: 0x9, XAdvance: 0x15, XOffset: 7, YOffset: -20, Bitmaps: []uint8{0x1, 0x0, 0x1c, 0x1, 0xb0, 0x19, 0x81, 0x86, 0x18, 0x11, 0x80, 0xd8, 0x3, 0x80, 0x18}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x15, Height: 0x1, XAdvance: 0x15, XOffset: -1, YOffset: 4, Bitmaps: []uint8{0xff, 0xff, 0xf8}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x5, Height: 0x5, XAdvance: 0x15, XOffset: 9, YOffset: -21, Bitmaps: []uint8{0xc7, 0x1c, 0x71, 0x80}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x10, Height: 0xf, XAdvance: 0x15, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x3, 0xf8, 0xc, 0xc, 0x0, 0x2, 0x0, 0x2, 0x0, 0x2, 0x0, 0x2, 0x7, 0xfc, 0x18, 0xc, 0x20, 0x4, 0x40, 0x4, 0x80, 0x4, 0x80, 0x8, 0x80, 0x38, 0xc0, 0xe8, 0x3f, 0xf}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x13, Height: 0x15, XAdvance: 0x15, XOffset: 1, YOffset: -20, Bitmaps: []uint8{0xf, 0x0, 0x0, 0x20, 0x0, 0x4, 0x0, 0x1, 0x80, 0x0, 0x30, 0x0, 0x4, 0x0, 0x0, 0x87, 0xc0, 0x13, 0xc, 0x6, 0x80, 0x40, 0xe0, 0xc, 0x18, 0x0, 0x82, 0x0, 0x10, 0xc0, 0x2, 0x10, 0x0, 0x42, 0x0, 0x8, 0x40, 0x2, 0x8, 0x0, 0x43, 0x80, 0x10, 0x70, 0x4, 0x9, 0x83, 0xf, 0x1f, 0x80}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x11, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1, 0xfc, 0x83, 0x3, 0xc6, 0x0, 0xe4, 0x0, 0x22, 0x0, 0x12, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x80, 0x0, 0x40, 0x0, 0x20, 0x0, 0x18, 0x0, 0x64, 0x0, 0x61, 0x81, 0xc0, 0x7f, 0x0}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0x12, Height: 0x15, XAdvance: 0x15, XOffset: 4, YOffset: -20, Bitmaps: []uint8{0x0, 0x3, 0xc0, 0x0, 0x30, 0x0, 0xc, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x60, 0x3f, 0x18, 0x10, 0x64, 0x18, 0xd, 0x8, 0x1, 0xc2, 0x0, 0x71, 0x0, 0xc, 0x80, 0x2, 0x20, 0x0, 0x88, 0x0, 0x62, 0x0, 0x18, 0x80, 0xe, 0x20, 0x3, 0x4, 0x3, 0x40, 0xc1, 0xb0, 0x1f, 0x8f, 0x0}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x10, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1, 0xf0, 0xe, 0xc, 0x18, 0x6, 0x30, 0x2, 0x60, 0x1, 0x40, 0x1, 0xc0, 0x1, 0xff, 0xff, 0x80, 0x0, 0x80, 0x0, 0x80, 0x0, 0x40, 0x0, 0x60, 0x6, 0x30, 0x1c, 0xf, 0xe0}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x13, Height: 0x15, XAdvance: 0x15, XOffset: 4, YOffset: -20, Bitmaps: []uint8{0x0, 0x1f, 0xe0, 0xc, 0x0, 0x3, 0x0, 0x0, 0x40, 0x0, 0x8, 0x0, 0x2, 0x0, 0x7, 0xff, 0xc0, 0x8, 0x0, 0x1, 0x0, 0x0, 0x20, 0x0, 0x8, 0x0, 0x1, 0x0, 0x0, 0x20, 0x0, 0x4, 0x0, 0x0, 0x80, 0x0, 0x20, 0x0, 0x4, 0x0, 0x0, 0x80, 0x0, 0x10, 0x0, 0x4, 0x0, 0xf, 0xff, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x13, Height: 0x16, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x3, 0xe3, 0xe1, 0x83, 0x60, 0x40, 0x38, 0x10, 0x3, 0x4, 0x0, 0x60, 0x80, 0xc, 0x20, 0x1, 0x84, 0x0, 0x20, 0x80, 0x4, 0x10, 0x1, 0x82, 0x0, 0x30, 0x60, 0xc, 0x4, 0x2, 0x80, 0x61, 0x90, 0x7, 0xc6, 0x0, 0x0, 0xc0, 0x0, 0x10, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x30, 0x0, 0xc, 0x0, 0xfe, 0x0, 0x0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x12, Height: 0x15, XAdvance: 0x15, XOffset: 2, YOffset: -20, Bitmaps: []uint8{0xf, 0x0, 0x0, 0x40, 0x0, 0x10, 0x0, 0x8, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x23, 0xe0, 0xb, 0xc, 0x5, 0x0, 0x81, 0x80, 0x20, 0x40, 0x8, 0x10, 0x2, 0x8, 0x0, 0x82, 0x0, 0x60, 0x80, 0x18, 0x20, 0x6, 0x10, 0x1, 0x84, 0x0, 0x61, 0x0, 0x30, 0x40, 0xc, 0xfc, 0x1f, 0xc0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0xf, Height: 0x16, XAdvance: 0x15, XOffset: 3, YOffset: -21, Bitmaps: []uint8{0x0, 0x30, 0x0, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x0, 0x20, 0x0, 0x40, 0x1, 0x0, 0x2, 0x0, 0x4, 0x0, 0x8, 0x0, 0x10, 0x0, 0x40, 0x0, 0x80, 0x1, 0x0, 0x2, 0x0, 0x8, 0x0, 0x10, 0x1f, 0xff, 0x80}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0xf, Height: 0x1d, XAdvance: 0x15, XOffset: 3, YOffset: -21, Bitmaps: []uint8{0x0, 0x6, 0x0, 0xc, 0x0, 0x18, 0x0, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xfe, 0x0, 0x4, 0x0, 0x8, 0x0, 0x10, 0x0, 0x20, 0x0, 0x80, 0x1, 0x0, 0x2, 0x0, 0x4, 0x0, 0x8, 0x0, 0x20, 0x0, 0x40, 0x0, 0x80, 0x1, 0x0, 0x6, 0x0, 0x8, 0x0, 0x10, 0x0, 0x20, 0x0, 0x80, 0x3, 0x0, 0xc, 0xf, 0xe0, 0x0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x12, Height: 0x15, XAdvance: 0x15, XOffset: 2, YOffset: -20, Bitmaps: []uint8{0x7, 0x80, 0x0, 0x60, 0x0, 0x10, 0x0, 0x4, 0x0, 0x1, 0x0, 0x0, 0xc0, 0x0, 0x30, 0xfc, 0x8, 0x18, 0x2, 0xc, 0x0, 0x8c, 0x0, 0x66, 0x0, 0x1b, 0x0, 0x5, 0x80, 0x1, 0xb0, 0x0, 0x46, 0x0, 0x31, 0xc0, 0xc, 0x30, 0x2, 0x6, 0x0, 0x80, 0xc0, 0x60, 0x30, 0xf8, 0x1f, 0x80}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0xf, Height: 0x15, XAdvance: 0x15, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0x1, 0xf8, 0x0, 0x20, 0x0, 0x40, 0x0, 0x80, 0x1, 0x0, 0x4, 0x0, 0x8, 0x0, 0x10, 0x0, 0x20, 0x0, 0x80, 0x1, 0x0, 0x2, 0x0, 0x4, 0x0, 0x8, 0x0, 0x20, 0x0, 0x40, 0x0, 0x80, 0x1, 0x0, 0x4, 0x0, 0x8, 0xf, 0xff, 0xc0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x14, Height: 0xf, XAdvance: 0x15, XOffset: 1, YOffset: -14, Bitmaps: []uint8{0x1c, 0xf1, 0xe0, 0xf1, 0xe3, 0xe, 0x1c, 0x10, 0xc1, 0x81, 0x8, 0x10, 0x30, 0x81, 0x3, 0x18, 0x10, 0x21, 0x83, 0x2, 0x10, 0x30, 0x21, 0x2, 0x6, 0x10, 0x20, 0x63, 0x2, 0x4, 0x30, 0x60, 0x42, 0x6, 0x4, 0xf8, 0x70, 0xf0}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x11, Height: 0xf, XAdvance: 0x15, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0xe, 0x3e, 0x1, 0x60, 0x81, 0xc0, 0x20, 0xc0, 0x10, 0x40, 0x8, 0x20, 0x4, 0x30, 0x2, 0x10, 0x2, 0x8, 0x1, 0x4, 0x0, 0x82, 0x0, 0x42, 0x0, 0x21, 0x0, 0x20, 0x80, 0x13, 0xf0, 0x3e}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x10, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1, 0xf0, 0x6, 0xc, 0x18, 0x6, 0x20, 0x3, 0x60, 0x1, 0x40, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x2, 0x80, 0x6, 0xc0, 0x4, 0x40, 0x18, 0x30, 0x60, 0x1f, 0x80}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x14, Height: 0x16, XAdvance: 0x15, XOffset: 0, YOffset: -14, Bitmaps: []uint8{0xf, 0x1f, 0x80, 0x16, 0xc, 0x1, 0xc0, 0x20, 0x30, 0x3, 0x3, 0x0, 0x10, 0x20, 0x1, 0x2, 0x0, 0x10, 0x40, 0x1, 0x4, 0x0, 0x10, 0x40, 0x2, 0x6, 0x0, 0x60, 0x60, 0x4, 0xb, 0x0, 0x80, 0x98, 0x30, 0x8, 0xfc, 0x0, 0x80, 0x0, 0x8, 0x0, 0x1, 0x0, 0x0, 0x10, 0x0, 0x1, 0x0, 0x0, 0x10, 0x0, 0xf, 0xf0, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x13, Height: 0x16, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x3, 0xf1, 0xe1, 0x83, 0x20, 0x40, 0x34, 0x10, 0x3, 0x84, 0x0, 0x30, 0x80, 0x4, 0x20, 0x0, 0x84, 0x0, 0x10, 0x80, 0x6, 0x10, 0x0, 0xc2, 0x0, 0x30, 0x60, 0xe, 0x4, 0x3, 0x40, 0x60, 0xc8, 0x7, 0xe2, 0x0, 0x0, 0x40, 0x0, 0x8, 0x0, 0x1, 0x0, 0x0, 0x20, 0x0, 0x8, 0x0, 0x1, 0x0, 0x3, 0xfc, 0x0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x13, Height: 0xf, XAdvance: 0x15, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0xf, 0x87, 0xc0, 0x23, 0x8, 0x4, 0xc0, 0x0, 0xe0, 0x0, 0x18, 0x0, 0x2, 0x0, 0x0, 0x80, 0x0, 0x10, 0x0, 0x2, 0x0, 0x0, 0x40, 0x0, 0x10, 0x0, 0x2, 0x0, 0x0, 0x40, 0x0, 0x8, 0x0, 0x3f, 0xfe, 0x0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0xf, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1, 0xfa, 0xc, 0x1c, 0x20, 0x8, 0x80, 0x11, 0x0, 0x3, 0x0, 0x3, 0xf8, 0x0, 0x7c, 0x0, 0xc, 0x0, 0x9, 0x0, 0x16, 0x0, 0x2c, 0x0, 0x9e, 0x6, 0x27, 0xf0, 0x0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0xd, Height: 0x14, XAdvance: 0x15, XOffset: 5, YOffset: -19, Bitmaps: []uint8{0x8, 0x0, 0x40, 0x2, 0x0, 0x10, 0x0, 0x80, 0x7f, 0xfc, 0x40, 0x2, 0x0, 0x10, 0x0, 0x80, 0x8, 0x0, 0x40, 0x2, 0x0, 0x10, 0x1, 0x0, 0x8, 0x0, 0x40, 0x2, 0x0, 0xd8, 0x1c, 0x3f, 0x0}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xf, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0xf0, 0x1e, 0x20, 0x4, 0x80, 0x9, 0x0, 0x12, 0x0, 0x24, 0x0, 0xc8, 0x1, 0x20, 0x2, 0x40, 0x4, 0x80, 0x9, 0x0, 0x12, 0x0, 0x64, 0x3, 0x8c, 0x1d, 0xf, 0xc3, 0x80}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x13, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0xfe, 0xf, 0xe6, 0x0, 0x20, 0x40, 0x8, 0x8, 0x3, 0x1, 0x80, 0x40, 0x30, 0x18, 0x6, 0x2, 0x0, 0x40, 0x80, 0x8, 0x30, 0x1, 0x84, 0x0, 0x31, 0x80, 0x2, 0x20, 0x0, 0x48, 0x0, 0x9, 0x0, 0x1, 0xc0, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x11, Height: 0xf, XAdvance: 0x15, XOffset: 5, YOffset: -14, Bitmaps: []uint8{0xf8, 0xf, 0xa0, 0x1, 0x90, 0x0, 0x88, 0x40, 0xc4, 0x30, 0x42, 0x18, 0x61, 0x1a, 0x20, 0x8d, 0x10, 0x4c, 0x98, 0x26, 0x48, 0x16, 0x2c, 0xb, 0x14, 0x7, 0xa, 0x3, 0x7, 0x1, 0x81, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x13, Height: 0xf, XAdvance: 0x15, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0xf, 0x83, 0xe0, 0xc0, 0x18, 0xc, 0xc, 0x1, 0x83, 0x0, 0x18, 0xc0, 0x1, 0xb0, 0x0, 0x1c, 0x0, 0x3, 0x0, 0x0, 0xf0, 0x0, 0x63, 0x0, 0x18, 0x30, 0x6, 0x6, 0x1, 0x80, 0x60, 0x60, 0x6, 0x3f, 0x7, 0xe0}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x15, Height: 0x16, XAdvance: 0x15, XOffset: 1, YOffset: -14, Bitmaps: []uint8{0xf, 0xc0, 0xf8, 0x30, 0x1, 0x0, 0x80, 0x18, 0x4, 0x0, 0x80, 0x30, 0xc, 0x1, 0x80, 0xc0, 0x4, 0x4, 0x0, 0x30, 0x60, 0x1, 0x86, 0x0, 0x4, 0x20, 0x0, 0x23, 0x0, 0x1, 0xb0, 0x0, 0xd, 0x0, 0x0, 0x38, 0x0, 0x1, 0x80, 0x0, 0x8, 0x0, 0x0, 0xc0, 0x0, 0x4, 0x0, 0x0, 0x60, 0x0, 0x6, 0x0, 0x0, 0x20, 0x0, 0x7f, 0xe0, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x10, Height: 0xf, XAdvance: 0x15, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x1f, 0xff, 0x10, 0x6, 0x10, 0xc, 0x10, 0x18, 0x0, 0x30, 0x0, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x3, 0x0, 0x6, 0x0, 0xc, 0x0, 0x18, 0x4, 0x30, 0xc, 0x60, 0xc, 0xff, 0xf8}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0xb, Height: 0x19, XAdvance: 0x15, XOffset: 8, YOffset: -20, Bitmaps: []uint8{0x0, 0xe0, 0x20, 0x8, 0x1, 0x0, 0x20, 0x4, 0x1, 0x0, 0x20, 0x4, 0x0, 0x80, 0x20, 0x8, 0xe, 0x0, 0x60, 0x4, 0x0, 0x80, 0x10, 0x2, 0x0, 0x40, 0x8, 0x2, 0x0, 0x40, 0x8, 0x1, 0x0, 0x18, 0x0}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x6, Height: 0x18, XAdvance: 0x15, XOffset: 9, YOffset: -19, Bitmaps: []uint8{0x0, 0x10, 0xc3, 0x8, 0x20, 0x86, 0x18, 0x41, 0x4, 0x30, 0xc2, 0x8, 0x21, 0x86, 0x10, 0x43, 0xc, 0x20}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0xa, Height: 0x19, XAdvance: 0x15, XOffset: 6, YOffset: -20, Bitmaps: []uint8{0x6, 0x0, 0x40, 0x10, 0x4, 0x1, 0x0, 0x40, 0x10, 0x4, 0x2, 0x0, 0x80, 0x20, 0xc, 0x1, 0xc0, 0xc0, 0x40, 0x10, 0x4, 0x3, 0x0, 0x80, 0x20, 0x8, 0x2, 0x1, 0x0, 0xc0, 0xe0, 0x0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0xf, Height: 0x5, XAdvance: 0x15, XOffset: 5, YOffset: -11, Bitmaps: []uint8{0x1e, 0x2, 0x66, 0xd, 0x86, 0x16, 0x6, 0x48, 0x7, 0x0}},
	},

	YAdvance: 0x23,
}
