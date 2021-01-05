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

func (slide *Slide) render(buffer *bytes.Buffer, config Config) error {

	fmt.Fprintf(buffer, " <def>\n")
	fmt.Fprintf(buffer, "  <g id=\"slide%d-def\">\n", slide.Id)
	fmt.Fprintf(buffer, "   <text x=\"30\" y=\"40\" fill=\"black\" font-size=\"32px\">%s</text>\n", slide.Title)
	fmt.Fprintf(buffer, "  </g>\n")
	fmt.Fprintf(buffer, " </def>\n")

	return nil
}
