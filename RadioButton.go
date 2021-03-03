package customwidget

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	theme "github.com/Nicify/theme"
)

type RadioButtonWidget struct {
	items    []string
	selected *int32
	onChange func()
}

func RadioButton(items []string, selected *int32) *RadioButtonWidget {
	return &RadioButtonWidget{
		items:    items,
		selected: selected,
	}
}

func (rb *RadioButtonWidget) OnChange(onChange func()) *RadioButtonWidget {
	rb.onChange = onChange
	return rb
}

func (rb *RadioButtonWidget) Build() {
	useLayoutFlat := theme.UseLayoutFlat()
	useLayoutFlat.Push()
	imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(color.RGBA{0, 0, 0, 0}))
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
	imgui.PushStyleColor(imgui.StyleColorButtonActive, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
	imgui.PushStyleColor(imgui.StyleColorText, g.ToVec4Color(color.RGBA{255, 255, 255, 255}))
	length := len(rb.items)
	for idx, item := range rb.items {
		currentSelected := int32(idx) == *rb.selected
		if currentSelected {
			imgui.PushStyleColor(imgui.StyleColorButton, g.ToVec4Color(color.RGBA{37, 122, 211, 255}))
		}
		if imgui.Button(item) {
			equal := *rb.selected == int32(idx)
			if !equal {
				*rb.selected = int32(idx)
				if rb.onChange != nil {
					rb.onChange()
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
}
