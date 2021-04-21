package customwidget

import (
	g "github.com/AllenDang/giu"
	imgui "github.com/AllenDang/imgui-go"
	theme "github.com/Nicify/theme"
)

type ImageButtonWidget struct {
	texture *g.Texture
	width   float32
	height  float32
	palette *theme.Palette
	onClick func()
}

func ImageButton(texture *g.Texture) *ImageButtonWidget {
	return &ImageButtonWidget{
		texture: texture,
	}
}

func (ib *ImageButtonWidget) Size(width float32, height float32) *ImageButtonWidget {
	ib.width = width
	ib.height = height
	return ib
}

func (ib *ImageButtonWidget) Palette(palette *theme.Palette) *ImageButtonWidget {
	ib.palette = palette
	return ib
}

func (ib *ImageButtonWidget) OnClick(onClick func()) *ImageButtonWidget {
	ib.onClick = onClick
	return ib
}

func (ib *ImageButtonWidget) Build() {
	if ib.palette != nil {
		imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(ib.palette.Tint))
		imgui.PushStyleColor(imgui.StyleColorButtonHovered, g.ToVec4Color(ib.palette.Hover))
		imgui.PushStyleColor(imgui.StyleColorButtonActive, g.ToVec4Color(ib.palette.Active))
	}
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
	g.ImageButton(ib.texture).Size(ib.width, ib.height).OnClick(ib.onClick).Build()
	imgui.PopStyleVar()
	if ib.palette != nil {
		imgui.PopStyleColorV(3)
	}
}
