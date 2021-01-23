package svgslides

import (
	"testing"
)

func TestImport(t *testing.T) {

	str := `
{"width":1024,"height":768,"rectWidth":180,"rectHeight":120,
"shapes":[
	{"x":150,"y":100,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":""},
	{"x":400,"y":300,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":""}
],
"connectors":[
	{"shape1":0,"shape2":1}
],
"transitions":[
	{"duration":3},
	{"duration":6}
]
}`

	t.Log("a developer")

	slides := New(Config{})
	err := slides.Import(str, "drawing")
	if err == nil {
		t.Log(" should be able to a slidedeck in jimareed/drawing format.", checkMark)
	} else {
		t.Fatal(" should be able to a slidedeck in jimareed/drawing format.", xMark)
	}

	if len(slides.Slides) == 1 {
		t.Log(" which should contain a single slide.", checkMark)
	} else {
		t.Fatal(" which should contain a single slide.", xMark, len(slides.Slides))
	}

	if len(slides.Slides[0].Shapes) == 2 {
		t.Log(" and two rectangles.", checkMark)
	} else {
		t.Fatal(" and two rectangles.", xMark, len(slides.Slides[0].Shapes))
	}
}

func TestImportDrawing(t *testing.T) {

	str := `
{"width":1024,"height":768,"rectWidth":40,"rectHeight":40,
 "shapes":[
     {"x":410,"y":230,"width":180,"height":120,"type":"text","desc":"Large text","size":24,"style":"","slide":""},
     {"x":210,"y":270,"width":180,"height":120,"type":"text","desc":"Small text","size":12,"style":"","slide":""}
]
}`

	t.Log("a developer")

	slides := New(Config{})
	err := slides.Import(str, "drawing")
	if err != nil {
		t.Fatal(" should be able to import a drawing with multiple text sizes.", xMark)
	}
	if len(slides.Slides) != 1 {
		t.Fatal(" which should contain a single slide.", xMark, len(slides.Slides))
	}
	if len(slides.Slides[0].Shapes) != 2 {
		t.Fatal(" and two shapes.", xMark, len(slides.Slides[0].Shapes))
	}

	if slides.Slides[0].Shapes[0].LabelSize == 24 {
		t.Log(" with correct sizes.", checkMark)
	} else {
		t.Fatal(" with correct sizes.", xMark, slides.Slides[0].Shapes[0].LabelSize)
	}
}
