package svgslides

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {

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
