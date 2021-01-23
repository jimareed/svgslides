package svgslides

import (
	"bytes"
	"math"
	"testing"
)

func aproxEqual(p1 Point, p2 Point) bool {

	return (math.Abs(p1.x-p2.x) < .5) && (math.Abs(p1.y-p2.y) < .5)
}

func initRects(p1 Point, p2 Point) (*Shape, *Shape) {
	slides := New(Config{200, 200, 20, 20, 12})

	slides.AddSlide("")
	rect1, _ := slides.AddRect("", p1.x, p1.y, 20, 20)
	rect2, _ := slides.AddRect("", p2.x, p2.y, 20, 20)

	return rect1, rect2
}

func TestRenderConnector(t *testing.T) {

	t.Log("a connector")

	rect1, rect2 := initRects(Point{90, 90}, Point{160, 90})

	connector := Connector{1, rect1.Id, rect2.Id}

	buffer := bytes.NewBuffer([]byte{})
	err := connector.render(buffer, Config{}, Animation{}, rect1, rect2)

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

func compareResults(t *testing.T, expectedP1 Point, expectedP2 Point, expectedSlope float64, actualP1 Point, actualP2 Point, actualSlope float64) {
	if actualP1 == expectedP1 {
		t.Log(" should match expected P1 coordinates", checkMark)
	} else {
		t.Fatalf(" should match expected P1 coordinates %s %v", xMark, actualP1)
	}

	if actualP2 == expectedP2 {
		t.Log(" should match expected P2 coordinates", checkMark)
	} else {
		t.Fatalf(" should match expected P2 coordinates %s %v", xMark, actualP2)
	}

	if actualSlope == expectedSlope {
		t.Log(" should match expected slope", checkMark)
	} else {
		t.Fatalf(" should match expected slope %s %f", xMark, actualSlope)
	}
}

func TestRightArrow(t *testing.T) {

	t.Log("a connector")

	rect1, rect2 := initRects(Point{90, 90}, Point{160, 90})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{110, 100}
	expectedP2 := Point{139, 100}
	expectedSlope := 0.0

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestLeftArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{10, 90})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{90, 100}
	expectedP2 := Point{51, 100}
	expectedSlope := 0.0

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestUpArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{90, 10})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{100, 90}
	expectedP2 := Point{100, 51}
	expectedSlope := math.Inf(-1)

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestDownArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{90, 160})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{100, 110}
	expectedP2 := Point{100, 139}
	expectedSlope := math.Inf(1)

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestUpperRightArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{150, 45})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{110, 92.5}
	expectedP2 := Point{133.2, 75.1}
	expectedSlope := -.75

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestUpperLeftArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{30, 45})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{90, 92.5}
	expectedP2 := Point{66.8, 75.1}
	expectedSlope := .75

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestLowerRightArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{150, 170})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{107.5, 110}
	expectedP2 := Point{139.9, 153.2}
	expectedSlope := 1.0/3.0 + 1.0

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestLowerLeftArrow(t *testing.T) {

	rect1, rect2 := initRects(Point{90, 90}, Point{30, 170})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{92.5, 110}
	expectedP2 := Point{60.1, 153.2}
	expectedSlope := (1.0/3.0 + 1.0) * -1

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}

func TestConnector(t *testing.T) {

	rect1, rect2 := initRects(Point{270, 170}, Point{40, 30})

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{270, 173.91}
	expectedP2 := Point{77.94, 57.0}

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)

	if aproxEqual(actualP1, expectedP1) {
		t.Log(" should match expected P1 coordinates", checkMark)
	} else {
		t.Fatalf(" should match expected P1 coordinates %s %v", xMark, actualP1)
	}

	if aproxEqual(actualP2, expectedP2) {
		t.Log(" should match expected P2 coordinates", checkMark)
	} else {
		t.Fatalf(" should match expected P2 coordinates %s %v", xMark, actualP2)
	}
}

func TestExample1Connector(t *testing.T) {

	slides := New(Config{1024, 768, 180, 120, 24})

	slides.AddSlide("")
	rect1, _ := slides.AddRect("", 166, 132, 180, 120)
	rect2, _ := slides.AddRect("", 166, 516, 180, 120)

	connector := Connector{1, rect1.Id, rect2.Id}

	expectedP1 := Point{256, 252}
	expectedP2 := Point{256, 495}
	expectedSlope := math.Inf(1)

	actualP1 := connector.getP1(rect1, rect2)
	actualP2 := connector.getP2(rect1, rect2)
	actualSlope := connector.slope(rect1, rect2)

	compareResults(t, expectedP1, expectedP2, expectedSlope, actualP1, actualP2, actualSlope)
}
