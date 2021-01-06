package svgslides

import (
	"bytes"
	"errors"
	"fmt"
)

// SvgSlides
type SvgSlides struct {
	Config         Config    `json:"config"`
	Slides         []Slide   `json:"slides"`
	CurrentSlideId int       `json:"currentSlideId"`
	NextSlideId    int       `json:"nextSlideId"`
	NextObjId      int       `json:"nextObjId"`
	Animation      Animation `json:"animation"`
}

// SvgSlidesConfig
type Config struct {
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	RectWidth  float64 `json:"rectWidth"`
	RectHeight float64 `json:"rectHeight"`
	LabelSize  int     `json:"labelSize"`
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
		config.RectWidth = 180
	}
	if config.RectHeight == 0 {
		config.RectHeight = 120
	}
	if config.LabelSize == 0 {
		config.LabelSize = 24
	}
	slides.Config = config
	slides.CurrentSlideId = 0
	slides.NextSlideId = 1
	slides.NextObjId = 1

	slides.Animation.Enabled = false

	return &slides
}

func (slides *SvgSlides) AddSlide(title string) error {
	slide := Slide{}
	slide.Id = slides.NextSlideId
	slide.Title = title
	slide.TitleObjId = slides.NextObjId
	slides.NextObjId++

	slides.Slides = append(slides.Slides, slide)
	slides.CurrentSlideId = slide.Id
	slides.NextSlideId++
	return nil
}

func (slides *SvgSlides) getSlide(slideId int) (*Slide, error) {

	for i := 0; i < len(slides.Slides); i++ {
		if slideId == slides.Slides[i].Id {
			return &(slides.Slides[i]), nil
		}
	}
	return nil, errors.New("invalid slide")
}

func (slides *SvgSlides) AddRect(label string, x float64, y float64) (*Shape, error) {

	slide, err := slides.getSlide(slides.CurrentSlideId)
	if err != nil {
		return nil, err
	}

	shape, err := slide.addRect(slides.NextObjId, label, x, y, slides.Config)
	slides.NextObjId++

	return shape, err
}

func (slides *SvgSlides) AddConnector(rect1 *Shape, rect2 *Shape) error {

	slide, err := slides.getSlide(slides.CurrentSlideId)
	if err != nil {
		return err
	}

	_, err = slide.addConnector(slides.NextObjId, rect1, rect2)
	slides.NextObjId++

	return err
}

func (slides *SvgSlides) AddAnimation(autoPlay bool) error {

	slides.Animation.Enabled = true
	slides.Animation.AutoPlay = autoPlay

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

	slides.Animation.updateSequence(slides)

	for _, slide := range slides.Slides {
		slide.render(buffer, slides.Config, slides.Animation)
	}
	fmt.Fprintf(buffer, "</svg>\n")

	return nil
}
