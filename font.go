package uprint

type font struct {
	width  int
	height int
	data   [][]uint16
}

var (
	fonts = map[string]font{
		"font6x8": {
			width:  6,
			height: 8,
			data:   font6x8,
		},
		"font7x10": {
			width:  7,
			height: 10,
			data:   font7x10,
		},
		"font11x18": {
			width:  11,
			height: 18,
			data:   font11x18,
		},
		"font16x26": {
			width:  16,
			height: 26,
			data:   font16x26,
		},
	}
)
