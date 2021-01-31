package svgslides

import (
	"bytes"
	"errors"
	"fmt"
)

const ANIMATION_DURATION = 3.0

type AnimationItem struct {
	ObjId         int     `json:"objId"`
	SlideId       int     `json:"slideId"`
	AnimationId   string  `json:"animationId"`
	Duration      float64 `json:"duration"`
	BeginDuration float64 `json:"beginDuration"`
	LastSlide     bool    `json:"lastSlide"`
}

// Animation
type Animation struct {
	Enabled  bool            `json:"enabled"`
	AutoPlay bool            `json:"autoPlay"`
	Duration float64         `json:"duration"`
	Sequence []AnimationItem `json:"sequence"`
}

func getAnimationId(objId int, slideId int) string {
	return fmt.Sprintf("slide%dobj%d", slideId, objId)
}

func getBeginAnimationId(slideId int) string {
	return fmt.Sprintf("slide%dbegin", slideId)
}

func getEndAnimationId(slideId int) string {
	return fmt.Sprintf("slide%dend", slideId)
}

func (animation *Animation) getAnimationItem(objId int) (*AnimationItem, error) {

	for i := 0; i < len(animation.Sequence); i++ {
		if objId == animation.Sequence[i].ObjId {
			return &(animation.Sequence[i]), nil
		}
	}
	return nil, errors.New("can't find animation item")
}

func (animation *Animation) updateSequence(slides *SvgSlides) {
	sequence := []AnimationItem{}

	if animation.Duration == 0 {
		animation.Duration = ANIMATION_DURATION
	}
	beginDuration := 0.0
	nextDuration := 0.0

	for slideIndex, slide := range slides.Slides {
		i := 0
		item := AnimationItem{}
		item.ObjId = slide.TitleObjId
		item.SlideId = slide.Id
		item.AnimationId = getBeginAnimationId(slide.Id)
		item.Duration = animation.Duration/2*float64(i) + animation.Duration/2
		item.BeginDuration = beginDuration
		item.LastSlide = slideIndex == (len(slides.Slides) - 1)
		if slide.Title == "" {
			item.Duration = .1
		}
		nextDuration = float64(item.Duration)
		sequence = append(sequence, item)

		for shapeIndex, shape := range slide.Shapes {
			i++
			item := AnimationItem{}
			item.ObjId = shape.Id
			item.SlideId = slide.Id
			item.AnimationId = getAnimationId(item.ObjId, slide.Id)
			if shapeIndex == (len(slide.Shapes) - 1) {
				item.AnimationId = getEndAnimationId(slide.Id)
			}
			item.Duration = animation.Duration/2*float64(i) + animation.Duration/2
			item.BeginDuration = beginDuration
			item.LastSlide = slideIndex == (len(slides.Slides) - 1)
			nextDuration = item.Duration
			sequence = append(sequence, item)
		}

		beginDuration = beginDuration + nextDuration*2 + animation.Duration
		nextDuration = 0.0
	}

	animation.Sequence = sequence
}

func (animation *Animation) render(buffer *bytes.Buffer, config Config, objId int, rederSuffix string) error {

	if animation.Enabled {
		item, _ := animation.getAnimationItem(objId)

		animationId := item.AnimationId

		fmt.Fprintf(buffer, "    <animate id=\"%sstep1\" attributeName=\"opacity\" from=\"0\" to=\"0\" dur=\"%.2f\" begin=\"0s\" repeatCount=\"1\" restart=\"always\" />\n", animationId, item.Duration+item.BeginDuration)
		fmt.Fprintf(buffer, "    <animate id=\"%sstep2\" attributeName=\"opacity\" from=\"0\" to=\"1\" dur=\"%.2f\" begin=\"%sstep1.end\" repeatCount=\"1\" restart=\"always\" />\n", animationId, item.Duration, animationId)
		if !item.LastSlide {
			fmt.Fprintf(buffer, "    <animate id=\"%sstep3\" attributeName=\"opacity\" from=\"1\" to=\"0\" dur=\"1.50\" begin=\"%sstep2.end\" repeatCount=\"0\" fill=\"freeze\" />\n", animationId, getEndAnimationId(item.SlideId))
		}
	}

	return nil
}
