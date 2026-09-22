package gogame

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func NewRect(x, y, width, height int) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) GetSize() [2]int {
	return [2]int{
		r.Width,
		r.Height,
	}
}

func (r Rect) GetRect() Rect {
	return r
}

func (r Rect) Top() int {
	return r.Y
}

func (r Rect) Bottom() int {
	return r.Y + r.Height
}

func (r Rect) Left() int {
	return r.X
}

func (r Rect) Right() int {
	return r.X + r.Width
}

func (r Rect) CenterX() int {
	return r.X + r.Width/2
}

func (r Rect) CenterY() int {
	return r.Y + r.Height/2
}

func (r Rect) CollidePoint(x, y int) bool {
	return x >= r.X &&
		x < r.X+r.Width &&
		y >= r.Y &&
		y < r.Y+r.Height
}

func (r Rect) Colliderect(other Rect) bool {
	return r.X < other.X+other.Width &&
		r.X+r.Width > other.X &&
		r.Y < other.Y+other.Height &&
		r.Y+r.Height > other.Y
}

func (r Rect) Move(x, y int) Rect {
	return Rect{
		X:      r.X + x,
		Y:      r.Y + y,
		Width:  r.Width,
		Height: r.Height,
	}
}

func (r *Rect) MoveIP(x, y int) {
	r.X += x
	r.Y += y
}