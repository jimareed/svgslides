package main

import (
	"bytes"
	"fmt"

	"github.com/jimareed/svgslides-go"
)

func main() {
	slides := svgslides.New(svgslides.Config{})
	slides.AddSlide("svgslides")
	rect1, _ := slides.AddRect("Go library", 166, 132, 180, 120)
	rect2, _ := slides.AddRect("Simple\nPresentations", 166, 516, 180, 120)
	rect3, _ := slides.AddRect("With\nAnimation", 678, 516, 180, 120)
	rect4, _ := slides.AddRect("In SVG\nFormat", 678, 132, 180, 120)
	slides.AddConnector(rect1, rect2)
	slides.AddConnector(rect2, rect3)
	slides.AddConnector(rect3, rect4)
	slides.AddAnimation(true)

	buffer := bytes.NewBuffer([]byte{})
	slides.Render(buffer)
	fmt.Println(buffer)
}
