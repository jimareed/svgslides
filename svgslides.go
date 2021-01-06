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

func (slides *SvgSlides) AddText(label string, x float64, y float64) (*Shape, error) {

	slide, err := slides.getSlide(slides.CurrentSlideId)
	if err != nil {
		return nil, err
	}

	shape, err := slide.addText(slides.NextObjId, label, x, y, slides.Config)
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

	fmt.Fprintf(buffer, "<svg id=\"canvas\" width=\"%.2f\" height=\"%.2f\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" >\n", slides.Config.Width, slides.Config.Height)

	if len(slides.Slides) > 0 {
		fmt.Fprintf(buffer, " <use id=\"slide%d\" xlink:href=\"#slide%d-def\" x=\"0\" y=\"0\" />\n", slides.Slides[0].Id, slides.Slides[0].Id)
	}

	if len(slides.Slides) > 1 {
		fmt.Fprintf(buffer, " <polygon class=\"nextslide-transition\" points=\"990 740, 975 753 990 766\"    stroke=\"lightgrey\" fill=\"lightgrey\" onmousedown=\"previousSlide(evt)\" onmouseover=\"evt.target.setAttribute('fill', 'black');\" onmouseout=\"evt.target.setAttribute('fill','lightgrey');\"></polygon>\n")
		fmt.Fprintf(buffer, " <polygon class=\"nextslide-transition\" points=\"1000 740, 1015 753 1000 766\" stroke=\"lightgrey\" fill=\"lightgrey\" onmousedown=\"nextSlide(evt)\"     onmouseover=\"evt.target.setAttribute('fill', 'black');\" onmouseout=\"evt.target.setAttribute('fill','lightgrey');\"></polygon>\n")
	}

	slides.Animation.updateSequence(slides)

	for _, slide := range slides.Slides {
		slide.render(buffer, slides.Config, slides.Animation)
	}

	if len(slides.Slides) > 1 {
		slides.renderPageButtons(buffer)
	}

	fmt.Fprintf(buffer, "</svg>\n")

	return nil
}

func (slides *SvgSlides) renderPageButtons(buffer *bytes.Buffer) error {

	fmt.Fprintf(buffer, " <script type=\"text/javascript\">\n")
	fmt.Fprintf(buffer, "  var svgNS = \"http://www.w3.org/2000/svg\";\n")
	fmt.Fprintf(buffer, "  var xlinkNS = \"http://www.w3.org/1999/xlink\";\n")
	fmt.Fprintf(buffer, "  var currentIndex = 1;\n")
	fmt.Fprintf(buffer, "  var maxIndex = %d;\n", len(slides.Slides))

	fmt.Fprintf(buffer, "  function previousSlide(evt) {\n")

	fmt.Fprintf(buffer, "  prevIndex = currentIndex - 1;\n")
	fmt.Fprintf(buffer, "  if (prevIndex === 0) {\n")
	fmt.Fprintf(buffer, "   prevIndex = 1;\n")
	fmt.Fprintf(buffer, "  }\n")

	fmt.Fprintf(buffer, "   canvas = document.getElementById('canvas');\n")
	fmt.Fprintf(buffer, "   prevslide = document.getElementById('slide' + currentIndex);\n")
	fmt.Fprintf(buffer, "   nextslide = document.createElementNS(svgNS, \"use\");\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"id\", \"slide\" + prevIndex);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"x\", 0);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"y\", 0);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(xlinkNS, \"xlink:href\", \"#\" + nextslide.id + \"-def\");\n")
	fmt.Fprintf(buffer, "   canvas.removeChild(prevslide)\n")
	fmt.Fprintf(buffer, "   canvas.appendChild(nextslide)\n")

	fmt.Fprintf(buffer, "   var animations = Array.from(svg.querySelectorAll(\"animate\"));\n")
	fmt.Fprintf(buffer, "   animations.forEach(function (item, index) {\n")
	fmt.Fprintf(buffer, "    if (item.id.includes('step1')) {\n")
	fmt.Fprintf(buffer, "     item.beginElement();\n")
	fmt.Fprintf(buffer, "    }\n")
	fmt.Fprintf(buffer, "   });\n")
	fmt.Fprintf(buffer, "   currentIndex = prevIndex\n")
	fmt.Fprintf(buffer, "  }\n")

	fmt.Fprintf(buffer, "  function nextSlide(evt) {\n")
	fmt.Fprintf(buffer, "   nextIndex = currentIndex + 1;\n")
	fmt.Fprintf(buffer, "   if (nextIndex > maxIndex) {\n")
	fmt.Fprintf(buffer, "    nextIndex = maxIndex;\n")
	fmt.Fprintf(buffer, "   }\n")
	fmt.Fprintf(buffer, "   svg = document.getElementById('canvas');\n")
	fmt.Fprintf(buffer, "   prevslide = document.getElementById('slide' + currentIndex);\n")
	fmt.Fprintf(buffer, "   nextslide = document.createElementNS(svgNS, \"use\");\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"id\", \"slide\" + nextIndex);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"x\", 0);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(null, \"y\", 0);\n")
	fmt.Fprintf(buffer, "   nextslide.setAttributeNS(xlinkNS, \"xlink:href\", \"#\" + nextslide.id + \"-def\");\n")
	fmt.Fprintf(buffer, "   svg.removeChild(prevslide)\n")
	fmt.Fprintf(buffer, "   svg.appendChild(nextslide)\n")

	fmt.Fprintf(buffer, "   var animations = Array.from(svg.querySelectorAll(\"animate\"));\n")
	fmt.Fprintf(buffer, "   animations.forEach(function (item, index) {\n")
	fmt.Fprintf(buffer, "    if (item.id.includes('step1')) {\n")
	fmt.Fprintf(buffer, "     item.beginElement();\n")
	fmt.Fprintf(buffer, "    }\n")
	fmt.Fprintf(buffer, "   });\n")
	fmt.Fprintf(buffer, "   currentIndex = nextIndex\n")
	fmt.Fprintf(buffer, "  }\n")
	fmt.Fprintf(buffer, " </script>\n")

	fmt.Fprintf(buffer, " <style>\n")
	fmt.Fprintf(buffer, "  .nextslide-transition {	animation-name: transitionOpacity;	animation-duration: 6s;	animation-iteration-count: 1;}\n")

	fmt.Fprintf(buffer, "  @keyframes transitionOpacity {\n")
	fmt.Fprintf(buffer, "   0%%   { opacity: 0; }\n")
	fmt.Fprintf(buffer, "   50%%   { opacity: 0; }\n")
	fmt.Fprintf(buffer, "   100%% { opacity: 1; }\n")
	fmt.Fprintf(buffer, "  }\n")
	fmt.Fprintf(buffer, " </style>\n")

	return nil
}
