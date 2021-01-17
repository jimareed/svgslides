package svgslides

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Shape
type Shape struct {
	Id        int     `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Label     string  `json:"label"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Type      string  `json:"type"`
	Size      int     `json:"size"`
	LabelSize int     `json:"labelSize"`
	Style     string  `json:"style"`
	X2        float64 `json:"x2"`
	Y2        float64 `json:"y2"`
}

func shapeFromString(input string) (Shape, error) {

	r := strings.NewReader(input)
	shape := Shape{}
	err := json.NewDecoder(r).Decode(&shape)

	return shape, err
}

func shapeToString(rect Shape) (string, error) {

	b, err := json.Marshal(rect)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func shapeToSvg(shape Shape, transitionId int) string {

	svg := ""

	if shape.Type == "text" {
		svg += fmt.Sprintf(
			"<text class=\"transition%d\" x=\"%f\" y=\"%f\" fill=\"black\" font-size=\"%dpx\">%s</text>\n",
			transitionId, shape.X, shape.Y, shape.Size, shape.Label)
	} else if shape.Type == "circle" {
		svg += fmt.Sprintf(
			"<circle class=\"transition%d\" cx=\"%f\" cy=\"%f\" r=\"%f\" stroke=\"black\" fill=\"transparent\" stroke-width=\"4\" \"></circle>\n",
			transitionId, shape.X, shape.Y, shape.Width/2)
	} else if shape.Type == "line" {
		svg += fmt.Sprintf(
			"<line class=\"transition%d\" x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"black\" stroke-width=\"4\" \"></line>\n",
			transitionId, shape.X, shape.Y, shape.X2, shape.Y2)
	} else if shape.Type == "rect" {
		strokeWidth := 4
		if shape.Style == "hidden" {
			strokeWidth = 0
		}
		onClick := ""
		//		if shape.Slide != "" {
		//			onClick = fmt.Sprintf("onclick=\"location.href='%s'\"", shape.Slide)
		//		}
		svg += fmt.Sprintf(
			"<rect class=\"transition%d\" x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" id=\"1\" stroke=\"black\" fill=\"transparent\" stroke-width=\"%d\" %s\"></rect>\n",
			transitionId, shape.X, shape.Y, shape.Width, shape.Height, strokeWidth, onClick)
	}

	return svg
}

func (shape *Shape) render(buffer *bytes.Buffer, config Config, animation Animation) error {

	switch shape.Type {
	case "rect":
		strokeWidth := 4
		if shape.Style == "hidden" {
			strokeWidth = 0
		}
		fmt.Fprintf(buffer, "   <rect x=\"%.2f\" y=\"%.2f\" width=\"%.2f\" height=\"%.2f\" id=\"%d\" stroke=\"black\" fill=\"transparent\" stroke-width=\"%d\">\n",
			shape.X, shape.Y, config.RectWidth, config.RectHeight, shape.Id, strokeWidth)
		animation.render(buffer, config, shape.Id, "")
		fmt.Fprintf(buffer, "   </rect>\n")
	case "text":
	default:
	}

	if shape.Label != "" {
		x := shape.X + shape.Width/2
		y := shape.Y + shape.Height/2
		labelSize := shape.LabelSize
		if labelSize == 0 {
			labelSize = config.LabelSize
		}
		label := Label{x, y, shape.Label, labelSize}
		label.render(buffer, config, animation, shape.Id)
	}

	return nil
}
