package svgslides

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderSlide(t *testing.T) {

	t.Log("a slide")

	slide := Slide{}
	slide.Id = 1
	slide.Title = "slide 1"

	buffer := bytes.NewBuffer([]byte{})
	err := slide.render(buffer, Config{})

	if err == nil {
		t.Log(" should render without errors", checkMark)
	} else {
		t.Fatal(" should render without errors", xMark, err)
	}

	if len(buffer.Bytes()) > 0 {
		t.Log(" should have a length greater than 0", checkMark)
	} else {
		t.Fatal(" should have a length greater than 0", xMark)
	}

}

func TestRenderSlideWithRect(t *testing.T) {

	t.Log("a slide")

	slide := Slide{}
	slide.Id = 1
	slide.Title = "slide 1"

	_, err := slide.addRect(1, "rect label", 50, 50, Config{})

	if err == nil {
		t.Log(" should add a rect without errors", checkMark)
	} else {
		t.Fatal(" should add a rect without errors", xMark, err)
	}

	buffer := bytes.NewBuffer([]byte{})
	err = slide.render(buffer, Config{})

	if err != nil {
		t.Fatal(" should render without errors", xMark, err)
	}

	str := string(buffer.Bytes())

	if strings.Contains(str, "<rect ") {
		t.Log(" should render a rect", checkMark)
	} else {
		t.Fatal(" should render a rect", xMark)
	}

}
