package svgslides

import "bytes"

// SvgSlides
type SvgSlides struct {
	Width      float64     `json:"width"`
	Height     float64     `json:"height"`
	RectWidth  float64     `json:"rectWidth"`
	RectHeight float64     `json:"rectHeight"`
	Shapes     []Shape     `json:"shapes"`
	Connectors []Connector `json:"connectors"`
	//	Animation  Animation   `json:"animation"`
}

// SvgSlidesConfig
type Config struct {
	Width      float64
	Height     float64
	RectWidth  float64
	RectHeight float64
}

func New(config Config) *SvgSlides {
	slides := SvgSlides{}

	if config.Width == 0 {
		config.Width = 1024
	}
	if config.Height == 0 {
		config.Height = 768
	}
	if config.RectWidth == 0 {
		config.RectWidth = 120
	}
	if config.RectHeight == 0 {
		config.RectHeight = 80
	}
	slides.Width = config.Width
	slides.Height = config.Height
	slides.RectWidth = config.RectWidth
	slides.RectHeight = config.RectHeight

	return &slides
}

func (slides *SvgSlides) NewSlide(title string) error {
	return nil
}

func (slides *SvgSlides) AddRect(label string, x float64, y float64) (Shape, error) {

	shape := Shape{x, y, slides.RectWidth, slides.RectHeight, "rect", "", 0, "", "", 0, 0}
	slides.Shapes = append(slides.Shapes, shape)
	return shape, nil
}

func (slides *SvgSlides) AddConnector(rect1 Shape, rect2 Shape) error {

	return nil
}

func (slides *SvgSlides) AddAnimation(timeInSec int, autoPlay bool) error {
	return nil
}

func (slides *SvgSlides) Render(buffer *bytes.Buffer) error {
	return nil
}
