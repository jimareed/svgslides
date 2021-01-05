package svgslides

import (
	"bytes"
	"testing"

	svg "github.com/h2non/go-is-svg"
)

const checkMark = "\u2713"
const xMark = "\u2717"

func TestAddSlide(t *testing.T) {

	t.Log("a developer")

	slides := New(Config{})
	err := slides.AddSlide("svgslides")

	if err == nil {
		t.Log(" should be able to create a new slide", checkMark)
	} else {
		t.Fatal(" should be able to create a new slide", xMark, err)
	}

	buffer := bytes.NewBuffer([]byte{})
	err = slides.Render(buffer)

	if err == nil {
		t.Log(" which should render without errors", checkMark)
	} else {
		t.Fatal(" which should render without errors", xMark, err)
	}

	if svg.Is(buffer.Bytes()) {
		t.Log(" as a valid svg", checkMark)
	} else {
		t.Fatal(" as a valid svg", xMark)
	}
}
