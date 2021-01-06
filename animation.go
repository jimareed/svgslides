package svgslides

import (
	"bytes"
	"errors"
	"fmt"
)

const ANIMATION_DURATION = 3.0

type AnimationItem struct {
	ObjId            int     `json:"objId"`
	AnimationId      string  `json:"animationId"`
	BeginAnimationId string  `json:"beginAnimiationId"`
	Duration         float64 `json:"duration"`
}

// Animation
type Animation struct {
	Enabled  bool            `json:"enabled"`
	AutoPlay bool            `json:"autoPlay"`
	Sequence []AnimationItem `json:"sequence"`
}

func getAnimationId(objId int, slideId int) string {
	return fmt.Sprintf("slide%dobj%d", slideId, objId)
}

func getBeginAnimationId(slideId int) string {
	return fmt.Sprintf("slide%dbegin", slideId)
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

	for _, slide := range slides.Slides {
		i := 0
		item := AnimationItem{}
		item.ObjId = slide.TitleObjId
		item.AnimationId = getBeginAnimationId(slide.Id)
		item.BeginAnimationId = getBeginAnimationId(slide.Id)
		item.Duration = ANIMATION_DURATION/2*float64(i) + 1.5
		if slide.Title == "" {
			item.Duration = .1
		}
		sequence = append(sequence, item)

		for _, shape := range slide.Shapes {
			i++
			item := AnimationItem{}
			item.ObjId = shape.Id
			item.AnimationId = getAnimationId(item.ObjId, slide.Id)
			item.BeginAnimationId = getBeginAnimationId(slide.Id)
			item.Duration = ANIMATION_DURATION/2*float64(i) + 1.5
			sequence = append(sequence, item)
		}
	}

	animation.Sequence = sequence
}

func (animation *Animation) render(buffer *bytes.Buffer, config Config, objId int, rederSuffix string) error {

	if animation.Enabled {
		item, _ := animation.getAnimationItem(objId)

		animationId := item.AnimationId

		fmt.Fprintf(buffer, "    <animate id=\"%sstep1\" attributeName=\"opacity\" from=\"0\" to=\"0\" dur=\"%.2f\" begin=\"0s\" repeatCount=\"1\" restart=\"always\" />\n", animationId, item.Duration)
		fmt.Fprintf(buffer, "    <animate id=\"%sstep2\" attributeName=\"opacity\" from=\"0\" to=\"1\" dur=\"%.2f\" begin=\"%sstep1.end\" repeatCount=\"1\" restart=\"always\" />\n", animationId, item.Duration, animationId)
	}

	return nil
}
