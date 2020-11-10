package customwidget

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	theme "github.com/Nicify/theme"
)

func ImageButton(texture *g.Texture, width float32, height float32, onClick func()) g.Layout {
	return g.Layout{
		g.Custom(func() {
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
		}),
		g.ImageButton(texture, width, height, onClick),
		g.Custom(func() {
			imgui.PopStyleVar()
		}),
	}
}

func ImageButtonV(texture *g.Texture, width float32, height float32, palette theme.Palette, onClick func()) g.Layout {
	return g.Layout{
		g.Custom(func() {
			imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(palette.Tint))
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, g.ToVec4Color(palette.Hover))
			imgui.PushStyleColor(imgui.StyleColorButtonActive, g.ToVec4Color(palette.Active))
		}),
		ImageButton(texture, width, height, onClick),
		g.Custom(func() {
			imgui.PopStyleColorV(3)
		}),
	}
}

func WithHiDPIFont(hDPIFont imgui.Font, lDPIFont imgui.Font, layout g.Layout) g.Layout {
	font := hDPIFont
	if imgui.DPIScale == 1 {
		font = lDPIFont
	}
	useFont := theme.UseFont(font)
	return g.Layout{
		g.Custom(useFont.Push),
		layout,
		g.Custom(useFont.Pop),
	}
}

func RadioButton(items []string, selected *int32, onChange func()) g.Layout {
	useLayoutFlat := theme.UseLayoutFlat()
	return g.Layout{
		g.Custom(func() {
			useLayoutFlat.Push()
			imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(color.RGBA{0, 0, 0, 0}))
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
			imgui.PushStyleColor(imgui.StyleColorButtonActive, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
			imgui.PushStyleColor(imgui.StyleColorText, g.ToVec4Color(color.RGBA{255, 255, 255, 255}))
			length := len(items)
			for idx, item := range items {
				currentSelected := int32(idx) == *selected
				if currentSelected {
					imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
				}
				if imgui.Button(item) {
					equal := *selected == int32(idx)
					if !equal {
						*selected = int32(idx)
						if onChange != nil {
							onChange()
						}
					}
				}
				if idx < length-1 {
					imgui.SameLineV(0, 0)
				}
				if currentSelected {
					imgui.PopStyleColorV(1)
				}
			}
			imgui.PopStyleColorV(4)
			useLayoutFlat.Pop()
		}),
	}
}
