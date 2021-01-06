package svgslides

import (
	"bytes"
	"testing"
)

func TestRenderShape(t *testing.T) {

	t.Log("a shape")

	shape := Shape{}
	shape.Id = 1
	shape.X = 240
	shape.Y = 180
	shape.Label = "Shape 1"

	buffer := bytes.NewBuffer([]byte{})
	err := shape.render(buffer, Config{}, Animation{})

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

/*
func TestShapeFromString(t *testing.T) {

	rect1 := Shape{100, 80, 90, 60, "", "", 24, "", "", 0, 0}

	str, err := shapeToString(rect1)
	if err != nil {
		t.Log("shapeToString error")
		t.Fail()
	}

	rect2, err := shapeFromString(str)
	if err != nil {
		t.Log("shapeFromString error")
		t.Fail()
	}

	if rect2.X != 100 || rect2.Y != 80 {
		t.Log("shape To/From string error")
		t.Fail()
	}
}
*/
