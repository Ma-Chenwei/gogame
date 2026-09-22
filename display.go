package gogame

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type DisplayModule struct {
	screen *Surface

	width  int
	height int
	title  string

	mu sync.RWMutex
}

var Display = &DisplayModule{}

func (d *DisplayModule) SetMode(size []int) *Surface {
	if len(size) < 2 {
		panic("gogame: SetMode requires [width, height]")
	}

	return d.SetModeWH(size[0], size[1])
}

func (d *DisplayModule) SetModeWH(width, height int) *Surface {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.width = width
	d.height = height

	if d.title == "" {
		d.title = "gogame"
	}

	d.screen = SurfaceFromImage(
		ebiten.NewImage(width, height),
	)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(d.title)

	return d.screen
}

func (d *DisplayModule) SetCaption(title string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.title = title

	ebiten.SetWindowTitle(title)
}

func (d *DisplayModule) GetSurface() *Surface {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.screen
}

func (d *DisplayModule) GetSize() [2]int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return [2]int{
		d.width,
		d.height,
	}
}

func (d *DisplayModule) Flip() {
}

func (d *DisplayModule) Update() {
}

func (d *DisplayModule) GetRect() Rect {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return Rect{
		X:      0,
		Y:      0,
		Width:  d.width,
		Height: d.height,
	}
}

func (d *DisplayModule) Run(update func(), draw func(*Surface)) error {
	d.mu.RLock()

	if d.screen == nil {
		d.mu.RUnlock()

		return fmt.Errorf(
			"gogame: display.set_mode() has not been called",
		)
	}

	width := d.width
	height := d.height

	d.mu.RUnlock()

	return ebiten.RunGame(
		&displayGame{
			display: d,
			update:  update,
			draw:    draw,
			width:   width,
			height:  height,
		},
	)
}

type displayGame struct {
	display *DisplayModule

	update func()
	draw   func(*Surface)

	width  int
	height int
}

func (g *displayGame) Update() error {
	if !Running() {
		return ebiten.Termination
	}

	if g.update != nil {
		g.update()
	}

	return nil
}

func (g *displayGame) Draw(screen *ebiten.Image) {
	if g.draw != nil {
		g.draw(
			SurfaceFromImage(screen),
		)

		return
	}

	if g.display.screen != nil {
		screen.DrawImage(
			g.display.screen.Image,
			nil,
		)
	}
}

func (g *displayGame) Layout(
	outsideWidth,
	outsideHeight int,
) (int, int) {
	return g.width, g.height
}