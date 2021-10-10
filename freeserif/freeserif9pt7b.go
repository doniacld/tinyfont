package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var Regular9pt7b = tinyfont.Font{
	BBox: [4]int8{17, 17, 0, -12},
	Glyphs: []tinyfont.Glypher{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x5, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x2, Height: 0xc, XAdvance: 0x6, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xff, 0xea, 0x3}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x5, Height: 0x4, XAdvance: 0x7, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xde, 0xf7, 0x20}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x11, 0x9, 0x4, 0x82, 0x4f, 0xf9, 0x10, 0x89, 0xff, 0x24, 0x12, 0x9, 0xc, 0x80}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x8, Height: 0xe, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x7c, 0xd6, 0xd2, 0xd0, 0xf0, 0x38, 0x1e, 0x17, 0x93, 0x93, 0xd6, 0x7c, 0x10}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0xd, Height: 0xc, XAdvance: 0xf, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x38, 0x43, 0x3c, 0x39, 0x21, 0x8a, 0xc, 0x50, 0x65, 0x39, 0xcb, 0x20, 0xb9, 0x5, 0x88, 0x4c, 0x44, 0x64, 0x21, 0xc0}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xd, Height: 0xd, XAdvance: 0xe, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xe, 0x0, 0xc8, 0x6, 0x40, 0x32, 0x1, 0xa0, 0x7, 0x78, 0x31, 0x87, 0x88, 0x46, 0x86, 0x34, 0x30, 0xc1, 0xc7, 0x17, 0xcf, 0x0}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x2, Height: 0x4, XAdvance: 0x4, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xfe}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x5, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x8, 0x88, 0x84, 0x63, 0x18, 0xc6, 0x10, 0x82, 0x8, 0x20}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x5, Height: 0xf, XAdvance: 0x6, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x82, 0x8, 0x21, 0xc, 0x63, 0x18, 0xc4, 0x22, 0x22, 0x0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x6, Height: 0x8, XAdvance: 0x9, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0x63, 0x9a, 0xdc, 0x72, 0xb6, 0x8}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x9, Height: 0x9, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x8, 0x4, 0x2, 0x1, 0xf, 0xf8, 0x40, 0x20, 0x10, 0x8, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x2, Height: 0x3, XAdvance: 0x4, XOffset: 2, YOffset: 0, Bitmaps: []uint8{0xd8}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x4, Height: 0x1, XAdvance: 0x6, XOffset: 1, YOffset: -3, Bitmaps: []uint8{0xf0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x2, Height: 0x2, XAdvance: 0x5, XOffset: 1, YOffset: -1, Bitmaps: []uint8{0xf0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x8, 0x84, 0x22, 0x10, 0x8c, 0x42, 0x31, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x1c, 0x31, 0x98, 0xd8, 0x3c, 0x1e, 0xf, 0x7, 0x83, 0xc1, 0xe0, 0xd8, 0xc4, 0x61, 0xc0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x5, Height: 0xd, XAdvance: 0x9, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x13, 0x8c, 0x63, 0x18, 0xc6, 0x31, 0x8c, 0x67, 0x80}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3c, 0x4e, 0x86, 0x6, 0x6, 0x4, 0xc, 0x8, 0x10, 0x20, 0x41, 0xfe}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3c, 0xc6, 0x6, 0x4, 0x1c, 0x3e, 0x7, 0x3, 0x3, 0x3, 0x6, 0xf8}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x7, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x4, 0x18, 0x71, 0x64, 0xc9, 0xa3, 0x46, 0xfe, 0x18, 0x30, 0x60}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf, 0x10, 0x20, 0x3c, 0xe, 0x7, 0x3, 0x3, 0x3, 0x2, 0x4, 0xf8}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x8, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x7, 0x1c, 0x30, 0x60, 0x60, 0xdc, 0xe6, 0xc3, 0xc3, 0xc3, 0x43, 0x66, 0x3c}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x7f, 0x82, 0x2, 0x2, 0x4, 0x4, 0x4, 0x8, 0x8, 0x8, 0x10, 0x10}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x7, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3c, 0x8f, 0x1e, 0x3e, 0x4f, 0x6, 0x36, 0xc7, 0x8f, 0x1b, 0x33, 0xc0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x8, Height: 0xe, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3c, 0x66, 0xc2, 0xc3, 0xc3, 0xc3, 0xc3, 0x63, 0x3f, 0x6, 0x6, 0xc, 0x38, 0x60}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x2, Height: 0x8, XAdvance: 0x5, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xf0, 0xf}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x3, Height: 0xa, XAdvance: 0x5, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xd8, 0x0, 0x3, 0x28}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x9, Height: 0x9, XAdvance: 0xa, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x1, 0x87, 0xe, 0x1c, 0xc, 0x3, 0x80, 0x70, 0xe, 0x0, 0x80}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x9, Height: 0x5, XAdvance: 0xa, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0xff, 0x80, 0x0, 0x0, 0xf, 0xf8}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xa, Height: 0x9, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x80, 0x1c, 0x1, 0xc0, 0x1c, 0x1, 0xc0, 0xe0, 0xe0, 0xe0, 0xc0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x7, Height: 0xd, XAdvance: 0x8, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x79, 0x1a, 0x18, 0x30, 0x60, 0x83, 0x4, 0x10, 0x20, 0x40, 0x3, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0xc, Height: 0xd, XAdvance: 0x10, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xf, 0x83, 0x8c, 0x60, 0x26, 0x2, 0xc7, 0x9c, 0xc9, 0xd8, 0x9d, 0x99, 0xd9, 0x26, 0xec, 0x60, 0x3, 0x4, 0xf, 0x80}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xd, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x2, 0x0, 0x10, 0x1, 0xc0, 0x16, 0x0, 0x98, 0x4, 0xc0, 0x43, 0x3, 0xf8, 0x20, 0x61, 0x3, 0x18, 0x1d, 0xe1, 0xf0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xb, Height: 0xc, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xff, 0x86, 0x1c, 0xc1, 0x98, 0x33, 0xc, 0x7e, 0xc, 0x31, 0x83, 0x30, 0x66, 0xc, 0xc3, 0x7f, 0xc0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1f, 0x26, 0x1d, 0x81, 0xe0, 0x1c, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0xc, 0x0, 0xc1, 0x8f, 0xc0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xff, 0x3, 0x1c, 0x30, 0x63, 0x7, 0x30, 0x33, 0x3, 0x30, 0x33, 0x3, 0x30, 0x33, 0x6, 0x30, 0xcf, 0xf0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xa, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0x98, 0x26, 0x1, 0x80, 0x61, 0x1f, 0xc6, 0x11, 0x80, 0x60, 0x18, 0x16, 0xf, 0xfe}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x9, Height: 0xc, XAdvance: 0xa, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xb0, 0x58, 0xc, 0x6, 0x13, 0xf9, 0x84, 0xc0, 0x60, 0x30, 0x18, 0x1e, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1f, 0x23, 0xe, 0x60, 0x26, 0x0, 0xc0, 0xc, 0xf, 0xc0, 0x6c, 0x6, 0xc0, 0x66, 0x6, 0x30, 0x60, 0xf8}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf1, 0xec, 0x19, 0x83, 0x30, 0x66, 0xc, 0xff, 0x98, 0x33, 0x6, 0x60, 0xcc, 0x19, 0x83, 0x78, 0xf0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x4, Height: 0xc, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf6, 0x66, 0x66, 0x66, 0x66, 0x6f}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x6, Height: 0xc, XAdvance: 0x7, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3c, 0x61, 0x86, 0x18, 0x61, 0x86, 0x18, 0x6d, 0xbc}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf3, 0xe6, 0x8, 0x61, 0x6, 0x20, 0x64, 0x7, 0x80, 0x6c, 0x6, 0x60, 0x63, 0x6, 0x18, 0x60, 0xcf, 0x3f}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xa, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf0, 0x18, 0x6, 0x1, 0x80, 0x60, 0x18, 0x6, 0x1, 0x80, 0x60, 0x18, 0x16, 0xb, 0xfe}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0xf, Height: 0xc, XAdvance: 0x10, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf0, 0xe, 0x70, 0x38, 0xe0, 0x71, 0xe1, 0x62, 0xc2, 0xc5, 0xc9, 0x89, 0x93, 0x13, 0x26, 0x23, 0x8c, 0x47, 0x18, 0x84, 0x33, 0x88, 0xf0}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xe0, 0xee, 0x9, 0xc1, 0x2c, 0x25, 0xc4, 0x9c, 0x91, 0x92, 0x1a, 0x41, 0xc8, 0x19, 0x3, 0x70, 0x20}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0xb, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0x6, 0x31, 0x83, 0x20, 0x2c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x68, 0x9, 0x83, 0x18, 0xc1, 0xf0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x9, Height: 0xc, XAdvance: 0xa, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xfe, 0x31, 0x98, 0x6c, 0x36, 0x1b, 0x19, 0xf8, 0xc0, 0x60, 0x30, 0x18, 0x1e, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0xb, Height: 0x10, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0x6, 0x31, 0x83, 0x20, 0x2c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x68, 0x19, 0x83, 0x18, 0xc0, 0xe0, 0xe, 0x0, 0xe0, 0x7}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xfe, 0xc, 0x61, 0x86, 0x30, 0xc6, 0x18, 0xc6, 0x1f, 0x83, 0x70, 0x67, 0xc, 0x71, 0x87, 0x78, 0x70}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x9, Height: 0xc, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1d, 0x31, 0x98, 0x4c, 0x7, 0x80, 0xe0, 0x1c, 0x7, 0x1, 0xa0, 0xd8, 0xcb, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xb, Height: 0xc, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xff, 0xf8, 0xce, 0x18, 0x83, 0x0, 0x60, 0xc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0x18, 0x7, 0x80}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf0, 0xec, 0x9, 0x81, 0x30, 0x26, 0x4, 0xc0, 0x98, 0x13, 0x2, 0x60, 0x4c, 0x8, 0xc2, 0xf, 0x80}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf8, 0x77, 0x2, 0x30, 0x23, 0x4, 0x18, 0x41, 0x84, 0xc, 0x80, 0xc8, 0x7, 0x0, 0x70, 0x2, 0x0, 0x20}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x11, Height: 0xc, XAdvance: 0x11, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfb, 0xe7, 0xb0, 0xc0, 0x8c, 0x20, 0x86, 0x18, 0x41, 0x8c, 0x40, 0xcb, 0x20, 0x65, 0x90, 0x1a, 0x70, 0xe, 0x38, 0x3, 0x1c, 0x1, 0x4, 0x0, 0x82, 0x0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xd, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfc, 0xf9, 0x83, 0x6, 0x10, 0x19, 0x0, 0xd0, 0x3, 0x0, 0x1c, 0x1, 0x30, 0x11, 0xc1, 0x86, 0x8, 0x19, 0xe3, 0xf0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf8, 0xf6, 0x6, 0x30, 0x41, 0x88, 0x1d, 0x0, 0xd0, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0xf0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xb, Height: 0xc, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xcc, 0x11, 0x6, 0x1, 0x80, 0x70, 0xc, 0x3, 0x0, 0xe0, 0x38, 0x6, 0x5, 0xc1, 0x7f, 0xe0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x3, Height: 0xf, XAdvance: 0x6, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xfb, 0x6d, 0xb6, 0xdb, 0x6d, 0xb8}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x82, 0x10, 0x82, 0x10, 0x86, 0x10, 0x86, 0x10}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x3, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xed, 0xb6, 0xdb, 0x6d, 0xb6, 0xf8}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x8, Height: 0x7, XAdvance: 0x8, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x18, 0x1c, 0x34, 0x26, 0x62, 0x42, 0xc1}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x9, Height: 0x1, XAdvance: 0x9, XOffset: 0, YOffset: 2, Bitmaps: []uint8{0xff, 0x80}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x4, Height: 0x3, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x84, 0x20}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x7, Height: 0x8, XAdvance: 0x8, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x79, 0x98, 0x30, 0xe6, 0xd9, 0xb3, 0x3f}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x20, 0x70, 0x18, 0xc, 0x6, 0x3, 0x71, 0xcc, 0xc3, 0x61, 0xb0, 0xd8, 0x6c, 0x63, 0xe0}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x7, Height: 0x8, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3c, 0xcf, 0x6, 0xc, 0x18, 0x18, 0x9e}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x1, 0x3, 0x80, 0xc0, 0x60, 0x31, 0xd9, 0x9d, 0x86, 0xc3, 0x61, 0xb0, 0xcc, 0x63, 0xf0}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x8, Height: 0x8, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3c, 0x46, 0xfe, 0xc0, 0xc0, 0xe1, 0x62, 0x3c}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x7, Height: 0xd, XAdvance: 0x7, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1e, 0x41, 0x83, 0x6, 0x1e, 0x18, 0x30, 0x60, 0xc1, 0x83, 0xf, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xa, Height: 0xc, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3c, 0x19, 0xf6, 0x31, 0x8c, 0x1e, 0x8, 0x4, 0x1, 0xfc, 0x40, 0xb0, 0x2e, 0x11, 0xf8}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x20, 0x70, 0x18, 0xc, 0x6, 0x3, 0x71, 0xcc, 0xc6, 0x63, 0x31, 0x98, 0xcc, 0x6f, 0x78}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x4, Height: 0xb, XAdvance: 0x5, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x60, 0x2, 0xe6, 0x66, 0x66, 0xf0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x5, Height: 0xf, XAdvance: 0x6, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x18, 0x0, 0x33, 0x8c, 0x63, 0x18, 0xc6, 0x31, 0x8b, 0x80}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x20, 0x70, 0x18, 0xc, 0x6, 0x3, 0x3d, 0x88, 0xd8, 0x78, 0x36, 0x19, 0x8c, 0x6f, 0x78}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x4, Height: 0xd, XAdvance: 0x5, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x2e, 0x66, 0x66, 0x66, 0x66, 0x66, 0xf0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0xe, Height: 0x8, XAdvance: 0xe, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xee, 0x71, 0xce, 0x66, 0x31, 0x98, 0xc6, 0x63, 0x19, 0x8c, 0x66, 0x31, 0xbd, 0xef}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x9, Height: 0x8, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xee, 0x39, 0x98, 0xcc, 0x66, 0x33, 0x19, 0x8d, 0xef}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x9, Height: 0x8, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3e, 0x31, 0xb0, 0x78, 0x3c, 0x1e, 0xd, 0x8c, 0x7c}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xee, 0x39, 0x98, 0x6c, 0x36, 0x1b, 0xd, 0x8c, 0xfc, 0x60, 0x30, 0x18, 0x1e, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3d, 0x31, 0xb0, 0xd8, 0x6c, 0x36, 0x1b, 0x8c, 0xfe, 0x3, 0x1, 0x80, 0xc0, 0xf0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x6, Height: 0x8, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x6d, 0xc6, 0x18, 0x61, 0x86, 0x3c}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x6, Height: 0x8, XAdvance: 0x7, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x76, 0x38, 0x58, 0x3e, 0x38, 0xfe}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x5, Height: 0x9, XAdvance: 0x5, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x27, 0x98, 0xc6, 0x31, 0x8c, 0x38}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x9, Height: 0x8, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xe7, 0x31, 0x98, 0xcc, 0x66, 0x33, 0x19, 0x8c, 0x7f}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x8, Height: 0x8, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xf3, 0x61, 0x22, 0x32, 0x14, 0x1c, 0x8, 0x8}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xc, Height: 0x8, XAdvance: 0xc, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xef, 0x36, 0x61, 0x62, 0x22, 0x32, 0x35, 0x41, 0x9c, 0x18, 0x81, 0x8}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x9, Height: 0x8, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xf7, 0x12, 0xe, 0x3, 0x1, 0xc1, 0x21, 0x9, 0xcf}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x8, Height: 0xc, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xf3, 0x61, 0x62, 0x32, 0x34, 0x14, 0x1c, 0x8, 0x8, 0x8, 0x10, 0xe0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x7, Height: 0x8, XAdvance: 0x7, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xfd, 0x18, 0x60, 0x83, 0xc, 0x70, 0xfe}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x5, Height: 0x10, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x19, 0x8c, 0x63, 0x18, 0xc4, 0x61, 0x8c, 0x63, 0x18, 0xc3}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0xc, XAdvance: 0x4, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xf0}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x5, Height: 0x10, XAdvance: 0x9, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0xc3, 0x18, 0xc6, 0x31, 0x84, 0x33, 0x18, 0xc6, 0x31, 0x98}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x9, Height: 0x3, XAdvance: 0x9, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x70, 0x24, 0xc1, 0xc0}},
	},

	YAdvance: 0x16,
}
