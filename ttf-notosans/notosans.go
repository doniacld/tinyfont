package notosans

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

type NotoSans12ptGlyph struct {
	Offset uint32
	Rune   rune
}

func (g NotoSans12ptGlyph) Draw(display drivers.Displayer, x int16, y int16, c1 color.RGBA) {
	info := g.Info()
	buf := g.Bytes()
	bmp := buf[5:]

	for j := int16(0); j < int16(info.Height); j++ {
		for i := int16(0); i < int16(info.Width); i++ {
			//display.SetPixel(x+int16(info.XOffset)+i, y+int16(info.YOffset)+j, color.RGBA{0x00, 0x00, 0x00, 0xFF})

			shiftNum := 0
			idx := i + j*int16(info.Width)
			switch idx % 4 {
			case 0:
				shiftNum = 6
			case 1:
				shiftNum = 4
			case 2:
				shiftNum = 2
			case 3:
				shiftNum = 0
			}

			c := (bmp[idx/4] >> shiftNum) & 0x03
			c2 := color.RGBA{}

			//fmt.Printf("+%d+%d+%x ", idx, shiftNum, c)
			//fmt.Printf("%x ", c)
			//if i == int16(info.Width)-1 {
			//	fmt.Printf("\n")
			//}
			switch c {
			case 0:
				c2 = color.RGBA{
					R: 0xFF,
					G: 0xFF,
					B: 0xFF,
					A: 0xFF,
				}
			case 1:
				c2 = color.RGBA{
					R: c1.R + (0xFF-c1.R)/8*7,
					G: c1.G + (0xFF-c1.G)/8*7,
					B: c1.B + (0xFF-c1.B)/8*7,
					A: c1.A,
				}
			case 2:
				c2 = color.RGBA{
					R: c1.R + (0xFF-c1.R)/4,
					G: c1.G + (0xFF-c1.G)/4,
					B: c1.B + (0xFF-c1.B)/4,
					A: c1.A,
				}
			case 3:
				c2 = color.RGBA{
					R: c1.R,
					G: c1.G,
					B: c1.B,
					A: c1.A,
				}
			default:
				panic("err")
			}
			//c = 0xFF - c

			display.SetPixel(x+int16(info.XOffset)+i, y+int16(info.YOffset)+j, c2)
		}
	}
}

func (g NotoSans12ptGlyph) Info() *tinyfont.GlyphInfo {
	buf := g.Bytes()
	return &tinyfont.GlyphInfo{
		Rune:     g.Rune,
		Width:    buf[0],
		Height:   buf[1],
		XAdvance: buf[2],
		XOffset:  int8(buf[3]),
		YOffset:  int8(buf[4]),
	}
}
