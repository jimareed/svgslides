# svgslides

Go library to generate simple svg presentations

<p  align="center">
    <img src="./examples/example1/output.svg" width="50%" height="50%" alt="svgslides output"/>
</p>

## usage

```golang
package main

import (
    "fmt"
    "bytes"
    "github.com/jimareed/svgslides"
)

func main() {
    slides := svgslides.New(svgslides.Config{})
    slides.AddSlide("svgslides")
    rect1, _ := slides.AddRect("Go library", 310, 200)
    rect2, _ := slides.AddRect("Simple presentations", 310, 200)
    rect3, _ := slides.AddRect("With animation", 310, 200)
    rect4, _ := slides.AddRect("In SVG|format", 310, 200)
    slides.AddConnector(rect1, rect2)
    slides.AddConnector(rect2, rect3)
    slides.AddConnector(rect3, rect4)
    slides.AddAnimation(true)

    buffer := bytes.NewBuffer([]byte{})
    slides.Render(buffer)
    fmt.Println(buffer)
}
```
