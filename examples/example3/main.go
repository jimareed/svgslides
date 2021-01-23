package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jimareed/svgslides-go"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	slides := svgslides.New(svgslides.Config{})

	err = slides.FromString(text)
	if err != nil {
		log.Fatalln(err)
	}
	slides.AddAnimation(true)

	buffer := bytes.NewBuffer([]byte{})
	slides.Render(buffer)
	fmt.Println(buffer)
}
