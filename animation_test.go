package svgslides

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRenderAnimation(t *testing.T) {

	t.Log("an animation")

	slides := New(Config{})
	slides.AddSlide("slide 1")

	buffer := bytes.NewBuffer([]byte{})

	slides.Animation.updateSequence(slides)
	err := slides.Animation.render(buffer, slides.Config, slides.Slides[0].TitleObjId, "")

	if err == nil {
		t.Log(" should render without errors", checkMark)
	} else {
		t.Fatal(" should render without errors", xMark, err)
	}

	if len(buffer.Bytes()) == 0 {
		t.Log(" should have a length of 0 if disabled", checkMark)
	} else {
		t.Fatal(" should have a length of 0 if disabled", xMark)
	}

	buffer = bytes.NewBuffer([]byte{})

	slides.AddAnimation(true)

	slides.Animation.updateSequence(slides)
	err = slides.Animation.render(buffer, slides.Config, slides.Slides[0].TitleObjId, "")

	if len(buffer.Bytes()) > 0 {
		t.Log(" should have a length greater than 0", checkMark)
	} else {
		t.Fatal(" should have a length greater than 0", xMark)
	}

	fmt.Printf(string(buffer.Bytes()))
}
