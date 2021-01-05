package svgslides

import (
	"bytes"
	"fmt"
)

// SvgSlides
type SvgSlides struct {
	Config         Config  `json:"config"`
	Slides         []Slide `json:"slides"`
	CurrentSlideId int     `json:"currentSlideId"`
	NextSlideId    int     `json:"nextSlideId"`
	NextObjId      int     `json:"nextObjId"`
	//	Animation  Animation   `json:"animation"`
}

// SvgSlidesConfig
type Config struct {
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	RectWidth  float64 `json:"rectWidth"`
	RectHeight float64 `json:"rectHeight"`
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
	slides.Config = config
	slides.CurrentSlideId = 0
	slides.NextSlideId = 1
	slides.NextObjId = 1

	return &slides
}

func (slides *SvgSlides) AddSlide(title string) error {
	slide := Slide{}
	slide.Id = slides.NextSlideId
	slide.Title = title

	slides.Slides = append(slides.Slides, slide)
	slides.CurrentSlideId = slide.Id
	slides.NextSlideId++
	return nil
}

func (slides *SvgSlides) AddRect(label string, x float64, y float64) (Shape, error) {

	shape := Shape{slides.NextObjId, x, y, slides.Config.RectWidth, slides.Config.RectHeight, "rect", "", 0, "", "", 0, 0}
	slides.NextObjId++
	//	slides.Shapes = append(slides.Shapes, shape)
	return shape, nil
}

func (slides *SvgSlides) AddConnector(rect1 Shape, rect2 Shape) error {

	return nil
}

func (slides *SvgSlides) AddAnimation(autoPlay bool) error {
	return nil
}

func (slides *SvgSlides) Render(buffer *bytes.Buffer) error {

	err := slides.render(buffer)

	return err
}

func (slides *SvgSlides) render(buffer *bytes.Buffer) error {

	fmt.Fprintf(buffer, "<svg width=\"%.2f\" height=\"%.2f\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" >\n", slides.Config.Width, slides.Config.Height)

	if len(slides.Slides) > 0 {
		fmt.Fprintf(buffer, " <use id=\"slide%d\" xlink:href=\"#slide%d-def\" x=\"0\" y=\"0\" />\n", slides.Slides[0].Id, slides.Slides[0].Id)
	}

	for _, slide := range slides.Slides {
		slide.render(buffer, slides.Config)
	}
	fmt.Fprintf(buffer, "</svg>\n")

	return nil
}
