package main

import (
	"bytes"
	"fmt"

	"github.com/jimareed/svgslides-go"
)

func main() {
	slides := svgslides.New(svgslides.Config{})
	slides.AddSlide("slide 1")
	slides.AddText("This is slide 1", 678, 516)
	slides.AddSlide("slide 2")
	slides.AddText("Followed by slide 2", 678, 516)
	slides.AddSlide("slide 3")
	slides.AddText("Last slide", 678, 516)
	slides.AddAnimation(true)

	buffer := bytes.NewBuffer([]byte{})
	slides.Render(buffer)
	fmt.Println(buffer)
}
