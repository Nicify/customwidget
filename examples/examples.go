package main

import (
	"fmt"
	"image/color"
	"runtime"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
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
	g.SingleWindow("Examples").Layout(
		g.Spacing(),
		g.Label("WithHiDPIFont"),
		g.Child("WithHiDPIFont").Border(true).Size(0, 100).Layout(
			c.WithHiDPIFont(fontConsola, fontDefault, g.Layout{
				g.Label("The quick brown fox jumps over the lazy dog").Wrapped(true),
			}),
		),
		g.Spacing(),
		g.Label("RadioButton"),
		g.Child("RadioButton").Border(true).Size(0, 100).Layout(
			g.Line(
				g.Label("Fruit:"),
				c.RadioButton(items, &itemSelected).OnChange(func() {
					fmt.Printf("select %s\n", items[itemSelected])
				}),
			),
		),
	)
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
	w.Run(loop)
}
