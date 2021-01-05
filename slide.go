package svgslides

// Slide
type Slide struct {
	Shapes     []Shape     `json:"shapes"`
	Connectors []Connector `json:"connectors"`
}
