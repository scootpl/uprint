package uprint

import (
	"image/color"
)

type device interface {
	SetPixel(x int16, y int16, c color.RGBA)
}

func New(d device, fname string) *text {
	return &text{
		d: d,
		f: fonts[fname],
	}
}

type text struct {
	d device
	f font
}

func (t *text) isSet(n uint16, bit int) bool {
	val := n & (1 << (16 - bit))
	return (val > 0)
}

func (t *text) insertChar(char byte, x, y int, c color.RGBA) {
	char -= 32 // "space" ascii code
	var c1 color.RGBA

	for yi, b := range t.f.data[char] {
		for xi := 0; xi < t.f.width; xi++ {
			if !t.isSet(b, xi) {
				c1 = color.RGBA{}
			} else {
				c1 = c
			}
			t.d.SetPixel(int16(x+xi), int16(y+yi), c1)
		}
	}
}

func (t *text) Print(s string, x, y int, c color.RGBA) {
	for _, char := range []byte(s) {
		t.insertChar(char, x, y, c)
		x += t.f.width
	}
}
