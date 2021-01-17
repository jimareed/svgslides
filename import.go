package svgslides

import (
	"github.com/jimareed/drawing"
)

func (slides *SvgSlides) Import(str string, format string) error {

	d, err := drawing.FromString(str)
	if err != nil {
		return err
	}

	slides.Config.Width = d.Width
	slides.Config.Height = d.Height
	slides.Config.RectWidth = d.RectWidth
	slides.Config.RectHeight = d.RectHeight

	slides.AddSlide("")

	for _, shape := range d.Shapes {

		var newShape *Shape = nil

		switch shape.Type {
		case "rect":
			newShape, err = slides.AddRect("", shape.X, shape.Y)
		case "text":
			newShape, err = slides.AddText(shape.Desc, shape.X, shape.Y)
		}

		if newShape != nil && err == nil {
			newShape.Style = shape.Style
		}
	}

	return nil
}
