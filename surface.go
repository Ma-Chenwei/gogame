package gogame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Surface struct {
	Image *ebiten.Image
}

func NewSurface(size [2]int) *Surface {
	width := size[0]
	height := size[1]

	if width <= 0 {
		width = 1
	}

	if height <= 0 {
		height = 1
	}

	return &Surface{
		Image: ebiten.NewImage(width, height),
	}
}

func NewSurfaceWH(width, height int) *Surface {
	return NewSurface([2]int{width, height})
}

func SurfaceFromImage(image *ebiten.Image) *Surface {
	return &Surface{
		Image: image,
	}
}

func (s *Surface) GetSize() [2]int {
	if s == nil || s.Image == nil {
		return [2]int{0, 0}
	}

	bounds := s.Image.Bounds()

	return [2]int{
		bounds.Dx(),
		bounds.Dy(),
	}
}

func (s *Surface) GetRect() Rect {
	size := s.GetSize()

	return Rect{
		X:      0,
		Y:      0,
		Width:  size[0],
		Height: size[1],
	}
}

func (s *Surface) Fill(c Color) {
	if s == nil || s.Image == nil {
		return
	}

	s.Image.Fill(color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: c.A,
	})
}

func (s *Surface) Blit(src *Surface, rect Rect) {
	if s == nil || s.Image == nil {
		return
	}

	if src == nil || src.Image == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(
		float64(rect.X),
		float64(rect.Y),
	)

	s.Image.DrawImage(
		src.Image,
		op,
	)
}

func (s *Surface) BlitAt(src *Surface, x, y int) {
	if src == nil {
		return
	}

	s.Blit(
		src,
		Rect{
			X:      x,
			Y:      y,
			Width:  src.GetSize()[0],
			Height: src.GetSize()[1],
		},
	)
}