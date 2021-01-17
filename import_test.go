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
{"width":1024,"height":768,"rectWidth":180,"rectHeight":120,
"shapes":[
	{"x":380,"y":240,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
	{"x":150,"y":100,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":""},
	{"x":610,"y":100,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":""},
	{"x":150,"y":370,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":"instructions"},
	{"x":610,"y":370,"width":180,"height":120,"type":"rect","desc":"","size":0,"style":"","slide":"examples"},
	{"x":396,"y":310,"width":180,"height":120,"type":"text","desc":"Slide","size":36,"style":"","slide":""},
	{"x":470,"y":310,"width":180,"height":120,"type":"text","desc":"show","size":36,"style":"","slide":""},
	{"x":196,"y":150,"width":180,"height":120,"type":"text","desc":"Interactive","size":20,"style":"","slide":""},
	{"x":186,"y":180,"width":180,"height":120,"type":"text","desc":"presentations","size":20,"style":"","slide":""},
	{"x":679,"y":150,"width":180,"height":120,"type":"text","desc":"With","size":20,"style":"","slide":""},
	{"x":662,"y":180,"width":180,"height":120,"type":"text","desc":"animation","size":20,"style":"","slide":""},
	{"x":190,"y":435,"width":180,"height":120,"type":"text","desc":"Instructions","size":20,"style":"","slide":""},
	{"x":660,"y":435,"width":180,"height":120,"type":"text","desc":"Examples","size":20,"style":"","slide":""}
],
"connectors":[
	{"shape1":0,"shape2":1},
	{"shape1":0,"shape2":2},
	{"shape1":0,"shape2":3},
	{"shape1":0,"shape2":4}
],
"transitions":[
	{"duration":0},
	{"duration":12},
	{"duration":16},
	{"duration":20},
	{"duration":24},
	{"duration":4},
	{"duration":8},
	{"duration":12},
	{"duration":12},
	{"duration":16},
	{"duration":16},
	{"duration":20},
	{"duration":24}
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
}
