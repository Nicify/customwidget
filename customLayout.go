package customwidget

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	theme "github.com/Nicify/theme"
)

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
