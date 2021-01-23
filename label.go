package svgslides

import (
	"bytes"
	"fmt"
	"strings"
)

// Label
type Label struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Value string  `json:"value"`
	Size  int     `json:"size"`
	Style string  `json:"style"`
}

func encodeLabel(s string) string {
	s = strings.Replace(s, "&", "&amp;", -1)
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
	s = strings.Replace(s, `"`, "&quot;", -1)
	s = strings.Replace(s, "'", "&apos;", -1)

	return s
}

func (label *Label) render(buffer *bytes.Buffer, config Config, animation Animation, objId int) error {

	if label.Value == "" {
		return nil
	}

	dominantBaseline := "middle"
	textAnchor := "middle"

	if label.Style == "left" {
		dominantBaseline = "auto"
		textAnchor = "start"
	}

	items := strings.Split(label.Value, "\n")

	offset := (len(items) - 1) * (label.Size / 2)

	for i, item := range items {
		text := encodeLabel(item)
		y := label.Y - float64(offset) + float64(i*label.Size)
		fmt.Fprintf(buffer, "   <text x=\"%.2f\" y=\"%.2f\" fill=\"black\" dominant-baseline=\"%s\" text-anchor=\"%s\" font-size=\"%dpx\">%s\n",
			label.X, y, dominantBaseline, textAnchor, label.Size, text)
		animation.render(buffer, config, objId, "label")
		fmt.Fprintf(buffer, "   </text>\n")
	}

	return nil
}
