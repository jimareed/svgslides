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
			newShape, err = slides.AddRect("", shape.X, shape.Y, shape.Width, shape.Height)
			newShape.Style = shape.Style
		case "circle":
			newShape, err = slides.AddCircle("", shape.X, shape.Y, shape.Width)
		case "line":
			newShape, err = slides.AddLine("", shape.X, shape.Y, shape.X2, shape.Y2)
		case "text":
			newShape, err = slides.AddText(shape.Desc, shape.X, shape.Y)
			newShape.Style = "left"
			newShape.LabelSize = shape.Size
		}

		newShape, err = slides.Slides[0].getShape(newShape.Id) // find shape and update as required

		switch shape.Type {
		case "rect":
			newShape.Style = shape.Style
		case "text":
			newShape.Style = "left"
			newShape.LabelSize = shape.Size
		}
	}

	for _, connector := range d.Connectors {

		if connector.Shape1 < len(slides.Slides[0].Shapes) && connector.Shape2 < len(slides.Slides[0].Shapes) {
			slides.AddConnector(&slides.Slides[0].Shapes[connector.Shape1], &slides.Slides[0].Shapes[connector.Shape2])
		}
	}

	return nil
}
