package main

import (
	"fmt"
	"image/color"
	"runtime"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	c "github.com/Nicify/customwidget"
	theme "github.com/Nicify/theme"
	"github.com/flopp/go-findfont"
)

var (
	fontDefault imgui.Font
	fontConsola imgui.Font

	items        []string
	itemSelected int32
)

func loop() {
	useLayoutFluent := theme.UseLayoutFluent()
	defer useLayoutFluent.Pop()
	useLayoutFluent.Push()
	g.SingleWindow("Examples", g.Layout{
		g.Spacing(),
		g.Label("WithHiDPIFont"),
		g.Child("WithHiDPIFont", true, 0, 100, 0, g.Layout{
			c.WithHiDPIFont(fontConsola, fontDefault, g.Layout{
				g.LabelWrapped("The quick brown fox jumps over the lazy dog"),
			}),
		}),
		g.Spacing(),
		g.Label("RadioButton"),
		g.Child("RadioButton", true, 0, 100, 0, g.Layout{
			g.Line(
				g.Label("Fruit:"),
				c.RadioButton(items, &itemSelected, func() {
					fmt.Printf("select %s\n", items[itemSelected])
				}),
			),
		}),
	})
}

func loadFont() {
	fonts := g.Context.IO().Fonts()
	ranges := fonts.GlyphRangesDefault()
	fontPath, err := findfont.Find("./assets/tamzen8x16b.ttf")
	if err == nil {
		fontDefault = fonts.AddFontFromFileTTFV(fontPath, 16, imgui.DefaultFontConfig, ranges)
	}
	fontPath, err = findfont.Find("Consola.ttf")
	if err == nil {
		fontConsola = fonts.AddFontFromFileTTFV(fontPath, 16, imgui.DefaultFontConfig, ranges)
	}
}

func init() {
	runtime.LockOSThread()

	items = []string{"Apple", "Pear", "Orange"}
}

func main() {
	w := g.NewMasterWindow("Examples", 320, 480, 0, loadFont)
	w.SetBgColor(color.RGBA{255, 255, 255, 255})
	// imgui.StyleColorsDark()
	style := imgui.CurrentStyle()
	theme.SetThemeFluentDark(&style)
	w.Main(loop)
}
