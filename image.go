package gogame

import (
	"fmt"
	"image"
	"image/png"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

// ImageModule 对应 pygame.image。
type ImageModule struct{}

var Image = &ImageModule{}

// Load 加载图片。
//
// 对应 Pygame:
//
//	self.photo = pygame.image.load(pic_path)
//
// Go:
//
//	self.Photo = gogame.Image.Load(picPath)
//
// 支持 PNG、JPEG 等 Go image 标准库支持的格式。
func (i *ImageModule) Load(path string) *Surface {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf(
			"gogame: cannot load image %q: %v",
			path,
			err,
		))
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Sprintf(
			"gogame: cannot decode image %q: %v",
			path,
			err,
		))
	}

	return SurfaceFromImage(
		ebiten.NewImageFromImage(img),
	)
}

// LoadWithError 加载图片，但不直接 panic。
// 如果加载失败，返回 nil 和 error。
//
// 用法:
//
//	img, err := gogame.Image.LoadWithError("player.png")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (i *ImageModule) LoadWithError(path string) (*Surface, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return SurfaceFromImage(
		ebiten.NewImageFromImage(img),
	), nil
}

// LoadRGBA 返回 Go 标准库的 image.Image。
// 适合需要直接操作像素的情况。
func (i *ImageModule) LoadRGBA(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf(
			"gogame: cannot load image %q: %v",
			path,
			err,
		))
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Sprintf(
			"gogame: cannot decode image %q: %v",
			path,
			err,
		))
	}

	return img
}

// SavePNG 将 Surface 保存为 PNG。
//
// 对应类似:
//
//	pygame.image.save(surface, "output.png")
func (i *ImageModule) SavePNG(surface *Surface, path string) error {
	if surface == nil {
		return fmt.Errorf("gogame: cannot save nil surface")
	}

	if surface.Image == nil {
		return fmt.Errorf("gogame: surface contains nil image")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	bounds := surface.Image.Bounds()

	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(
				x,
				y,
				surface.Image.At(x, y),
			)
		}
	}

	return png.Encode(file, rgba)
}

// GetSize 获取图片尺寸。
//
// 用法:
//
//	size := gogame.Image.GetSize("player.png")
func (i *ImageModule) GetSize(path string) [2]int {
	surface, err := i.LoadWithError(path)
	if err != nil {
		return [2]int{0, 0}
	}

	return surface.GetSize()
}