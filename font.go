package gogame

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type FontObject struct {
	face font.Face
	size int
	name string
}

type FontModule struct{}

var Font = &FontModule{}

// SysFont 对应 pygame.font.SysFont(name, size)
func (f *FontModule) SysFont(name string, size int) *FontObject {
	if size <= 0 {
		size = 16
	}

	path := findSystemFont(name)

	if path == "" {
		panic(fmt.Sprintf(
			"gogame: font %q not found, please install the font or use Font.FontFile()",
			name,
		))
	}

	return f.FontFile(path, size)
}

// FontFile 对应从字体文件加载字体
func (f *FontModule) FontFile(path string, size int) *FontObject {
	if size <= 0 {
		size = 16
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("gogame: cannot read font %q: %v", path, err))
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		panic(fmt.Sprintf("gogame: cannot parse font %q: %v", path, err))
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(fmt.Sprintf("gogame: cannot create font face: %v", err))
	}

	return &FontObject{
		face: face,
		size: size,
		name: filepath.Base(path),
	}
}

// Render 将文字渲染成 Surface
func (f *FontObject) Render(text string, c Color) *Surface {
	if f == nil || f.face == nil {
		return nil
	}

	bounds, _ := font.BoundString(f.face, text)

	width := (bounds.Max.X - bounds.Min.X).Ceil()
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()

	if width < 1 {
		width = 1
	}

	if height < 1 {
		height = f.size
	}

	img := image.NewRGBA(image.Rect(0, 0, width+4, height+4))

	// 文字绘制位置
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{R: c.R, G: c.G, B: c.B, A: c.A}),
		Face: f.face,
		Dot: fixed.Point26_6{
			X: fixed.I(2),
			Y: fixed.I(2 + f.size),
		},
	}

	d.DrawString(text)

	return SurfaceFromImage(ebiten.NewImageFromImage(img))
}

// RenderText 是 Render 的别名
func (f *FontObject) RenderText(text string, c Color) *Surface {
	return f.Render(text, c)
}

func (f *FontObject) GetHeight() int {
	if f == nil {
		return 0
	}
	return f.size
}

func (f *FontObject) GetSize(text string) [2]int {
	if f == nil || f.face == nil {
		return [2]int{0, 0}
	}

	bounds, _ := font.BoundString(f.face, text)

	return [2]int{
		(bounds.Max.X - bounds.Min.X).Ceil(),
		(bounds.Max.Y - bounds.Min.Y).Ceil(),
	}
}

func (f *FontObject) Size(text string) [2]int {
	return f.GetSize(text)
}

// 查找系统字体
func findSystemFont(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	if name == "" {
		return ""
	}

	// Windows
	if runtime.GOOS == "windows" {
		paths := []string{
			`C:\Windows\Fonts`,
		}

		if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
			paths = append(paths,
				filepath.Join(userProfile, "AppData", "Local", "Microsoft", "Windows", "Fonts"),
			)
		}

		return searchFontDirectories(paths, name)
	}

	// Linux
	if runtime.GOOS == "linux" {
		paths := []string{
			"/usr/share/fonts",
			"/usr/local/share/fonts",
		}

		if home := os.Getenv("HOME"); home != "" {
			paths = append(paths,
				filepath.Join(home, ".fonts"),
				filepath.Join(home, ".local", "share", "fonts"),
			)
		}

		return searchFontDirectories(paths, name)
	}

	// macOS
	if runtime.GOOS == "darwin" {
		paths := []string{
			"/System/Library/Fonts",
			"/Library/Fonts",
		}

		if home := os.Getenv("HOME"); home != "" {
			paths = append(paths,
				filepath.Join(home, "Library", "Fonts"),
			)
		}

		return searchFontDirectories(paths, name)
	}

	return ""
}

func searchFontDirectories(paths []string, name string) string {
	extensions := map[string]bool{
		".ttf":  true,
		".otf":  true,
		".ttc":  true,
		".otc":  true,
	}

	for _, root := range paths {
		info, err := os.Stat(root)
		if err != nil || !info.IsDir() {
			continue
		}

		var result string

		_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}

			ext := strings.ToLower(filepath.Ext(path))
			if !extensions[ext] {
				return nil
			}

			base := strings.ToLower(strings.TrimSuffix(filepath.Base(path), ext))

			// 例如：
			// Saira.ttf
			// Saira-Regular.ttf
			// Saira-Bold.ttf
			// Saira_Regular.ttf
			if base == name ||
				strings.HasPrefix(base, name+"-") ||
				strings.HasPrefix(base, name+"_") ||
				strings.Contains(base, name) {

				result = path
				return filepath.SkipDir
			}

			return nil
		})

		if result != "" {
			return result
		}
	}

	return ""
}