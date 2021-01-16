package svgslides

import (
	"bytes"
	"fmt"
	"math"
)

const arrowHeadLength = 21

type Point struct {
	x, y float64
}

// Connector
type Connector struct {
	Id       int `json:"id"`
	ShapeId1 int `json:"shapeId1"`
	ShapeId2 int `json:"shapeId2"`
}

func arrowHeadXLength(slope float64) float64 {
	return arrowHeadLength / math.Sqrt(slope*slope+1)
}

func slope(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func (connector *Connector) slope(shape1 *Shape, shape2 *Shape) float64 {

	return (shape2.Y - shape1.Y) / (shape2.X - shape1.X)
}

func (connector *Connector) getP1(shape1 *Shape, shape2 *Shape) Point {

	p := Point{0.0, 0.0}
	p1 := Point{shape1.X, shape1.Y}
	p2 := Point{shape2.X, shape2.Y}

	s := connector.slope(shape1, shape2)

	if s == math.Inf(1) || s == math.Inf(-1) {
		p.x = p1.x + shape1.Width/2
		if p1.y < p2.y {
			p.y = p1.y + shape1.Height
		} else {
			p.y = p1.y
		}
	} else {
		if math.Abs(s) <= slope(0.0, 0.0, shape1.Width, shape1.Height) {
			if p1.x < p2.x {
				// right side
				p.x = p1.x + shape1.Width
				p.y = p1.y + shape1.Height/2 + shape1.Width/2*s
			} else {
				// left side
				p.x = p1.x
				p.y = p1.y + shape1.Height/2 - shape1.Width/2*s
			}
		} else {
			if p1.y > p2.y {
				// top side
				p.x = p1.x + shape1.Width/2 - (shape1.Height/2)/s
				p.y = p1.y
			} else {
				// botton side
				p.x = p1.x + shape1.Width/2 + (shape1.Height/2)/s
				p.y = p1.y + shape1.Height
			}
		}
	}
	return p
}

func (connector *Connector) getP2(shape1 *Shape, shape2 *Shape) Point {

	p := Point{0, 0}

	p1 := Point{shape1.X, shape1.Y}
	p2 := Point{shape2.X, shape2.Y}

	s := slope(p1.x, p1.y, p2.x, p2.y)

	if s == math.Inf(1) || s == math.Inf(-1) {
		p.x = p2.x + shape1.Width/2
		if p1.y < p2.y {
			p.y = p2.y - arrowHeadLength
		} else {
			p.y = p2.y + shape1.Height + arrowHeadLength
		}
	} else {
		arrowHeadX := arrowHeadXLength(s)
		arrowHeadY := arrowHeadX * s

		if math.Abs(s) <= slope(0, 0, shape1.Width, shape1.Height) {
			if p1.x < p2.x {
				// right side
				p.x = p2.x
				p.y = p2.y + shape1.Height/2 - shape1.Width/2*s
				p.x -= arrowHeadX
				p.y -= arrowHeadY
			} else {
				// left side
				p.x = p2.x + shape1.Width
				p.y = p2.y + shape1.Height/2 + shape1.Width/2*s
				p.x += arrowHeadX
				p.y += arrowHeadY
			}
		} else {
			if p1.y > p2.y {
				// top side
				p.x = p2.x + shape1.Width/2 + shape1.Height/2/s
				p.y = p2.y + shape1.Height
				if p1.x < p2.x {
					arrowHeadX = arrowHeadX * -1
				}
				p.x += arrowHeadX
				p.y += math.Abs(arrowHeadY)
			} else {
				// botton side
				p.x = p2.x + shape1.Width/2 - shape1.Height/2/s
				p.y = p2.y
				if p1.x < p2.x {
					arrowHeadX = arrowHeadX * -1
				}
				p.x += arrowHeadX
				p.y -= math.Abs(arrowHeadY)
			}
		}
	}

	return p
}

func (connector *Connector) render(buffer *bytes.Buffer, config Config, animation Animation, shape1 *Shape, shape2 *Shape) error {

	p1 := connector.getP1(shape1, shape2)
	p2 := connector.getP2(shape1, shape2)

	fmt.Fprintf(buffer, "   <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"black\" stroke-width=\"4\" marker-end=\"url(#arrowhead)\">\n",
		p1.x, p1.y, p2.x, p2.y)
	animation.render(buffer, config, shape2.Id, "")
	fmt.Fprintf(buffer, "   </line>\n")
	return nil
}

/*


func connectorP2(d *SvgSlides, c Connector) Point {
	p := Point{0, 0}

	p1 := Point{d.Shapes[c.Shape1].X, d.Shapes[c.Shape1].Y}
	p2 := Point{d.Shapes[c.Shape2].X, d.Shapes[c.Shape2].Y}

	s := slope(p1.x, p1.y, p2.x, p2.y)

	if s == math.Inf(1) || s == math.Inf(-1) {
		p.x = p2.x + shape1.Width/2
		if p1.y < p2.y {
			p.y = p2.y - arrowHeadLength
		} else {
			p.y = p2.y + shape1.Height + arrowHeadLength
		}
	} else {
		arrowHeadX := arrowHeadXLength(s)
		arrowHeadY := arrowHeadX * s

		if math.Abs(s) <= slope(0, 0, shape1.Width, shape1.Height) {
			if p1.x < p2.x {
				// right side
				p.x = p2.x
				p.y = p2.y + shape1.Height/2 - shape1.Width/2*s
				p.x -= arrowHeadX
				p.y -= arrowHeadY
			} else {
				// left side
				p.x = p2.x + shape1.Width
				p.y = p2.y + shape1.Height/2 + shape1.Width/2*s
				p.x += arrowHeadX
				p.y += arrowHeadY
			}
		} else {
			if p1.y > p2.y {
				// top side
				p.x = p2.x + shape1.Width/2 + shape1.Height/2/s
				p.y = p2.y + shape1.Height
				if p1.x < p2.x {
					arrowHeadX = arrowHeadX * -1
				}
				p.x += arrowHeadX
				p.y += math.Abs(arrowHeadY)
			} else {
				// botton side
				p.x = p2.x + shape1.Width/2 - shape1.Height/2/s
				p.y = p2.y
				if p1.x < p2.x {
					arrowHeadX = arrowHeadX * -1
				}
				p.x += arrowHeadX
				p.y -= math.Abs(arrowHeadY)
			}
		}
	}

	return p
}

*/
