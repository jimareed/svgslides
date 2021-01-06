package svgslides

import (
	"bytes"
	"errors"
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

func (slide *Slide) addConnector(id int, rect1 *Shape, rect2 *Shape) (*Connector, error) {

	connector := Connector{id, rect1.Id, rect2.Id}

	slide.Connectors = append(slide.Connectors, connector)

	return &connector, nil
}

func (slide *Slide) getShape(shapeId int) (*Shape, error) {

	for i := 0; i < len(slide.Shapes); i++ {
		if shapeId == slide.Shapes[i].Id {
			return &(slide.Shapes[i]), nil
		}
	}
	return nil, errors.New("invalid shape")
}

func (slide *Slide) render(buffer *bytes.Buffer, config Config) error {

	fmt.Fprintf(buffer, " <defs>\n")
	fmt.Fprintf(buffer, "  <g id=\"slide%d-def\">\n", slide.Id)
	fmt.Fprintf(buffer, "   <text x=\"50%%\" y=\"50%%\" dominant-baseline=\"middle\" text-anchor=\"middle\" fill=\"black\" font-size=\"32px\">%s</text>\n", slide.Title)

	for _, shape := range slide.Shapes {
		shape.render(buffer, config)
	}

	for _, connector := range slide.Connectors {
		rect1, err := slide.getShape(connector.ShapeId1)
		if err != nil {
			return err
		}
		rect2, err := slide.getShape(connector.ShapeId2)
		if err != nil {
			return err
		}
		connector.render(buffer, rect1, rect2)
	}

	fmt.Fprintf(buffer, "  </g>\n")
	fmt.Fprintf(buffer, "  <marker id=\"arrowhead\" markerWidth=\"5\" markerHeight=\"3.5\" refX=\"0\" refY=\"1.75\" orient=\"auto\">\n")
	fmt.Fprintf(buffer, "   <polygon points=\"0 0, 5 1.75 0 3.5\"></polygon>\n")
	fmt.Fprintf(buffer, "  </marker>\n")
	fmt.Fprintf(buffer, " </defs>\n")

	return nil
}
