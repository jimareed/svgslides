package svgslides

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

const arrowHeadLength = 21

type Point struct {
	x, y float64
}

// Shape
type Connector struct {
	Shape1 int `json:"shape1"`
	Shape2 int `json:"shape2"`
}

func connectorSlope(d *SvgSlides, c Connector) float64 {

	if len(d.Shapes) <= c.Shape1 || len(d.Shapes) <= c.Shape2 {
		return 0
	}

	return slope(d.Shapes[c.Shape1].X, d.Shapes[c.Shape1].Y, d.Shapes[c.Shape2].X, d.Shapes[c.Shape2].Y)
}

func connectorFromString(input string) (Connector, error) {

	r := strings.NewReader(input)
	c := Connector{}
	err := json.NewDecoder(r).Decode(&c)

	return c, err
}

func connectorToString(c Connector) (string, error) {

	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func connectorToSvg(d *SvgSlides, c Connector, transitionId int) string {

	svg := ""

	p1 := connectorP1(d, c)
	p2 := connectorP2(d, c)
	svg += fmt.Sprintf(
		"<line class=\"transition%d\" x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"black\" stroke-width=\"4\" marker-end=\"url(#arrowhead)\"></line>\n",
		transitionId, p1.x, p1.y, p2.x, p2.y)

	return svg
}

func slope(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func arrowHeadXLength(slope float64) float64 {
	return arrowHeadLength / math.Sqrt(slope*slope+1)
}

func connectorP1(d *SvgSlides, c Connector) Point {
	p := Point{0.0, 0.0}
	p1 := Point{d.Shapes[c.Shape1].X, d.Shapes[c.Shape1].Y}
	p2 := Point{d.Shapes[c.Shape2].X, d.Shapes[c.Shape2].Y}

	s := slope(p1.x, p1.y, p2.x, p2.y)

	if s == math.Inf(1) || s == math.Inf(-1) {
		p.x = p1.x + d.RectWidth/2
		if p1.y < p2.y {
			p.y = p1.y + d.RectHeight
		} else {
			p.y = p1.y
		}
	} else {
		if math.Abs(s) <= slope(0.0, 0.0, d.RectWidth, d.RectHeight) {
			if p1.x < p2.x {
				// right side
				p.x = p1.x + d.RectWidth
				p.y = p1.y + d.RectHeight/2 + d.RectWidth/2*s
			} else {
				// left side
				p.x = p1.x
				p.y = p1.y + d.RectHeight/2 - d.RectWidth/2*s
			}
		} else {
			if p1.y > p2.y {
				// top side
				p.x = p1.x + d.RectWidth/2 - (d.RectHeight/2)/s
				p.y = p1.y
			} else {
				// botton side
				p.x = p1.x + d.RectWidth/2 + (d.RectHeight/2)/s
				p.y = p1.y + d.RectHeight
			}
		}
	}
	return p
}

func connectorP2(d *SvgSlides, c Connector) Point {
	p := Point{0, 0}

	p1 := Point{d.Shapes[c.Shape1].X, d.Shapes[c.Shape1].Y}
	p2 := Point{d.Shapes[c.Shape2].X, d.Shapes[c.Shape2].Y}

	s := slope(p1.x, p1.y, p2.x, p2.y)

	if s == math.Inf(1) || s == math.Inf(-1) {
		p.x = p2.x + d.RectWidth/2
		if p1.y < p2.y {
			p.y = p2.y - arrowHeadLength
		} else {
			p.y = p2.y + d.RectHeight + arrowHeadLength
		}
	} else {
		arrowHeadX := arrowHeadXLength(s)
		arrowHeadY := arrowHeadX * s

		if math.Abs(s) <= slope(0, 0, d.RectWidth, d.RectHeight) {
			if p1.x < p2.x {
				// right side
				p.x = p2.x
				p.y = p2.y + d.RectHeight/2 - d.RectWidth/2*s
				p.x -= arrowHeadX
				p.y -= arrowHeadY
			} else {
				// left side
				p.x = p2.x + d.RectWidth
				p.y = p2.y + d.RectHeight/2 + d.RectWidth/2*s
				p.x += arrowHeadX
				p.y += arrowHeadY
			}
		} else {
			if p1.y > p2.y {
				// top side
				p.x = p2.x + d.RectWidth/2 + d.RectHeight/2/s
				p.y = p2.y + d.RectHeight
				if p1.x < p2.x {
					arrowHeadX = arrowHeadX * -1
				}
				p.x += arrowHeadX
				p.y += math.Abs(arrowHeadY)
			} else {
				// botton side
				p.x = p2.x + d.RectWidth/2 - d.RectHeight/2/s
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
