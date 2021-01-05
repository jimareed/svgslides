package svgslides

/*
func initSlides(rect1 Point, rect2 Point) *SvgSlides {

	slides := New(SvgSlidesConfig{200, 200, 20, 20})

	slides.AddRect("", rect1.x, rect1.y)
	slides.AddRect("", rect2.x, rect2.y)

	return slides
}

func aproxEqual(p1 Point, p2 Point) bool {

	return (math.Abs(p1.x-p2.x) < .5) && (math.Abs(p1.y-p2.y) < .5)
}

func TestConnectorFromString(t *testing.T) {

	c1 := Connector{1, 2}

	str, err := connectorToString(c1)
	if err != nil {
		t.Log("connectorToString error")
		t.Fail()
	}

	c2, err := connectorFromString(str)
	if err != nil {
		t.Log("connectorFromString error")
		t.Fail()
	}

	if c2.Shape1 != 1 || c2.Shape2 != 2 {
		t.Log("connector To/From string error")
		t.Fail()
	}

}

func TestRightArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{160, 90})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{110, 100}
	expectedP2 := Point{139, 100}
	expectedSlope := 0.0

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}

}

func TestLeftArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{10, 90})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{90, 100}
	expectedP2 := Point{51, 100}
	expectedSlope := 0.0

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}

}

func TestUpArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{90, 10})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{100, 90}
	expectedP2 := Point{100, 51}
	expectedSlope := math.Inf(-1)

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}

}

func TestDownArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{90, 160})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{100, 110}
	expectedP2 := Point{100, 139}
	expectedSlope := math.Inf(1)

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}

}

func TestUpperRightArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{150, 45})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{110, 92.5}
	expectedP2 := Point{133.2, 75.1}
	expectedSlope := -.75

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}
}

func TestUpperLeftArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{30, 45})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{90, 92.5}
	expectedP2 := Point{66.8, 75.1}
	expectedSlope := .75

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}
}

func TestLowerRightArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{150, 170})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{107.5, 110}
	expectedP2 := Point{139.9, 153.2}
	expectedSlope := 1.0/3.0 + 1.0

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if actualP1 != expectedP1 {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if actualP2 != expectedP2 {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}
}

func TestLowerLeftArrow(t *testing.T) {

	d1 := initSlides(Point{90, 90}, Point{30, 170})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{92.5, 110}
	expectedP2 := Point{60.1, 153.2}
	expectedSlope := (1.0/3.0 + 1.0) * -1

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)
	actualSlope := connectorSlope(d1, c1)

	if !aproxEqual(actualP1, expectedP1) {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if !aproxEqual(actualP2, expectedP2) {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
	if actualSlope != expectedSlope {
		t.Log("invalid Slope")
		t.Log(actualSlope)
		t.Fail()
	}
}

func TestConnector(t *testing.T) {

	d1 := initSlides(Point{270, 170}, Point{40, 30})

	c1 := Connector{0, 1}
	d1 = AddConnector(d1, c1)

	expectedP1 := Point{270, 173.91}
	expectedP2 := Point{77.94, 57.0}

	actualP1 := connectorP1(d1, c1)
	actualP2 := connectorP2(d1, c1)

	if !aproxEqual(actualP1, expectedP1) {
		t.Log("invalid P1")
		t.Log(actualP1)
		t.Fail()
	}
	if !aproxEqual(actualP2, expectedP2) {
		t.Log("invalid P2")
		t.Log(actualP2)
		t.Fail()
	}
}
*/
