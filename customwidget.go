package customwidget

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	theme "github.com/Nicify/theme"
)

func ImageButton(texture *g.Texture, width float32, height float32, onClick func()) g.Layout {
	return g.Layout{
		g.Custom(func() {
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 0)
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
		}),
		g.ImageButton(texture, width, height, onClick),
		g.Custom(func() {
			imgui.PopStyleVarV(2)
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
	return g.Layout{
		g.Custom(func() {
			if imgui.DPIScale == 1 {
				g.PushFont(lDPIFont)
			} else {
				g.PushFont(hDPIFont)
			}
		}),
		layout,
		g.Custom(func() {
			g.PopFont()
		}),
	}
}

func RadioButton(items []string, selected *int32, onChange func()) g.Layout {
	useLayoutFlat := theme.UseLayoutFlat()
	useStyleButtonGhost := theme.UseStyleButtonGhost()
	return g.Layout{
		g.Custom(func() {
			length := len(items)
			for idx, item := range items {
				currentSelected := int32(idx) == *selected
				useLayoutFlat.Push()
				if !currentSelected {
					useStyleButtonGhost.Push()
				}
				if imgui.Button(item) {
					*selected = int32(idx)
					if onChange != nil {
						onChange()
					}
				}
				if idx < length-1 {
					imgui.SameLineV(0, 0)
				}
				if !currentSelected {
					useStyleButtonGhost.Pop()
				}
				useLayoutFlat.Pop()
			}
		}),
	}
}
