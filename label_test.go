package svgslides

import (
	"bytes"
	"testing"
)

func TestRenderLabel(t *testing.T) {

	t.Log("a label")

	label := Label{100, 100, "label", 20, ""}

	buffer := bytes.NewBuffer([]byte{})
	err := label.render(buffer, Config{}, Animation{}, 1)

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
