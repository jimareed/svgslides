package svgslides

import (
	"bytes"
	"fmt"
)

// Slide
type Slide struct {
	Id         int         `json:"id"`
	Title      string      `json:"title"`
	Shapes     []Shape     `json:"shapes"`
	Connectors []Connector `json:"connectors"`
}

func (slide *Slide) addRect(id int, label string, x float64, y float64, config Config) (*Shape, error) {

	shape := Shape{}
	shape.Id = id
	shape.Label = label
	shape.X = x
	shape.Y = y
	shape.Width = config.RectWidth
	shape.Height = config.RectHeight

	slide.Shapes = append(slide.Shapes, shape)

	return &shape, nil
}

func (slide *Slide) render(buffer *bytes.Buffer, config Config) error {

	fmt.Fprintf(buffer, " <def>\n")
	fmt.Fprintf(buffer, "  <g id=\"slide%d-def\">\n", slide.Id)
	fmt.Fprintf(buffer, "   <text x=\"50%%\" y=\"50%%\" dominant-baseline=\"middle\" text-anchor=\"middle\" fill=\"black\" font-size=\"32px\">%s</text>\n", slide.Title)

	for _, shape := range slide.Shapes {
		shape.render(buffer, config)
	}

	fmt.Fprintf(buffer, "  </g>\n")
	fmt.Fprintf(buffer, " </def>\n")

	return nil
}
