package gogame

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func RGB(r, g, b uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func RGBA(r, g, b, a uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

var (
	BLACK = RGB(0, 0, 0)
	WHITE = RGB(255, 255, 255)
	RED   = RGB(255, 0, 0)
	GREEN = RGB(0, 255, 0)
	BLUE  = RGB(0, 0, 255)
	YELLOW = RGB(255, 255, 0)
)

const (
	NOEVENT = iota
	QUIT
	KEYDOWN
	KEYUP
	MOUSEMOTION
	MOUSEBUTTONDOWN
	MOUSEBUTTONUP
)