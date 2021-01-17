package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/jimareed/svgslides-go"
)

const importStr = `
{"width":1024,"height":768,"rectWidth":360,"rectHeight":60,
 "shapes":[
     {"x":300,"y":250,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":90,"y":140,"width":360,"height":60,"type":"rect","desc":"Self Awareness","size":0,"style":"hidden","slide":""},
     {"x":110,"y":60,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":220,"y":30,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":420,"y":20,"width":360,"height":60,"type":"rect","desc":"Social Awareness","size":0,"style":"hidden","slide":""},
     {"x":520,"y":80,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":81,"y":335,"width":360,"height":60,"type":"rect","desc":"Self Management","size":0,"style":"hidden","slide":""},
     {"x":82,"y":410,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":81,"y":515,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":178,"y":480,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":240,"y":550,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":580,"y":340,"width":360,"height":60,"type":"rect","desc":"Relationship Management","size":0,"style":"hidden","slide":""},
     {"x":540,"y":410,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":510,"y":470,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":455,"y":520,"width":360,"height":60,"type":"rect","desc":"","size":0,"style":"hidden","slide":""},
     {"x":305,"y":285,"width":180,"height":120,"type":"text","desc":"Emotional","size":36,"style":"","slide":""},
     {"x":470,"y":285,"width":180,"height":120,"type":"text","desc":"Intelligence","size":36,"style":"","slide":""},
     {"x":175,"y":180,"width":180,"height":130,"type":"text","desc":"Self Awareness","size":20,"style":"","slide":""},
     {"x":195,"y":95,"width":180,"height":130,"type":"text","desc":"Self Assessment","size":20,"style":"","slide":""},
     {"x":345,"y":60,"width":180,"height":130,"type":"text","desc":"Self Confidence","size":20,"style":"","slide":""},
     {"x":560,"y":60,"width":180,"height":130,"type":"text","desc":"Empathy","size":20,"style":"","slide":""},
     {"x":640,"y":120,"width":180,"height":130,"type":"text","desc":"Organizational Awareness","size":20,"style":"","slide":""},
     {"x":220,"y":355,"width":180,"height":130,"type":"text","desc":"Self Control","size":20,"style":"","slide":""},
     {"x":185,"y":435,"width":180,"height":130,"type":"text","desc":"Trustworthiness","size":20,"style":"","slide":""},
     {"x":205,"y":540,"width":180,"height":130,"type":"text","desc":"Motivation","size":20,"style":"","slide":""},
     {"x":325,"y":510,"width":180,"height":130,"type":"text","desc":"Disipline","size":20,"style":"","slide":""},
     {"x":375,"y":575,"width":180,"height":130,"type":"text","desc":"Optimism","size":20,"style":"","slide":""},
     {"x":680,"y":360,"width":180,"height":130,"type":"text","desc":"Influence","size":20,"style":"","slide":""},
     {"x":665,"y":440,"width":180,"height":130,"type":"text","desc":"Conflict Management","size":20,"style":"","slide":""},
     {"x":660,"y":500,"width":180,"height":130,"type":"text","desc":"Relationships","size":20,"style":"","slide":""},
     {"x":515,"y":550,"width":180,"height":130,"type":"text","desc":"Teamwork & Collaboration","size":20,"style":"","slide":""}
],
"connectors":[
     {"shape1":0,"shape2":1},
     {"shape1":0,"shape2":2},
     {"shape1":0,"shape2":3},
     {"shape1":0,"shape2":4},
     {"shape1":0,"shape2":5},
     {"shape1":0,"shape2":6},
     {"shape1":0,"shape2":7},
     {"shape1":0,"shape2":8},
     {"shape1":0,"shape2":9},
     {"shape1":0,"shape2":10},
     {"shape1":0,"shape2":11},
     {"shape1":0,"shape2":12},
     {"shape1":0,"shape2":13},
     {"shape1":0,"shape2":14}
],
"transitions":[
     {"duration":0},
     {"duration":8},
     {"duration":10},
     {"duration":12},
     {"duration":24},
     {"duration":26},
     {"duration":14},
     {"duration":16},
     {"duration":18},
     {"duration":20},
     {"duration":22},
     {"duration":28},
     {"duration":30},
     {"duration":32},
     {"duration":34},
     {"duration":4},
     {"duration":4},
     {"duration":8},
     {"duration":10},
     {"duration":12},
     {"duration":24},
     {"duration":26},
     {"duration":14},
     {"duration":16},
     {"duration":18},
     {"duration":20},
     {"duration":22},
     {"duration":28},
     {"duration":30},
     {"duration":32},
     {"duration":34}
],
"nextPage":{"name":"emotional-intelligence-2","delay":28}
}`

func main() {

	// converting the examples used in jimareed/slideshow-editor to svgslides format (work in progress)
	slides := svgslides.New(svgslides.Config{})
	err := slides.Import(importStr, "drawing")
	if err != nil {
		log.Fatalln(err)
	}
	slides.AddAnimation(true)

	buffer := bytes.NewBuffer([]byte{})
	slides.Render(buffer)
	fmt.Println(buffer)
}
