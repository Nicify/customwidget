package customwidget

import (
	g "github.com/AllenDang/giu"
)

type ImageWidget struct {
	texture *g.Texture
	width   float32
	height  float32
}

func Image(texture *g.Texture) *ImageButtonWidget {
	return &ImageButtonWidget{
		texture: texture,
	}
}

func (i *ImageWidget) Size(width float32, height float32) *ImageWidget {
	i.width = width
	i.height = height
	return i
}

func (i *ImageWidget) Build() {
	if i.texture == nil {
		g.Dummy(i.width, i.height).Build()
		return
	}
	g.Image(i.texture).Size(i.width, i.height).Build()
}
