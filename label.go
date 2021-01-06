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
}

func (label *Label) render(buffer *bytes.Buffer, config Config, animation Animation, objId int) error {

	if label.Value == "" {
		return nil
	}

	items := strings.Split(label.Value, "\n")

	offset := (len(items) - 1) * (label.Size / 2)

	for i, item := range items {
		y := label.Y - float64(offset) + float64(i*label.Size)
		fmt.Fprintf(buffer, "   <text x=\"%.2f\" y=\"%.2f\" fill=\"black\" dominant-baseline=\"middle\" text-anchor=\"middle\" font-size=\"%dpx\">%s\n",
			label.X, y, label.Size, item)
		animation.render(buffer, config, objId, "label")
		fmt.Fprintf(buffer, "   </text>\n")
	}

	return nil
}
