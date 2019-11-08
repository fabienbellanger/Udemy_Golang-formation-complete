package main

import (
	"fmt"

	"github.com/fabienbellanger/Udemy_Golang-formation-complete/Imgproc/filter"
)

func main() {
	// var g filter.Grayscale
	// err := g.Process("img/photo.jpeg", "img/photo_gs.jpeg")
	// fmt.Println(err)

	// var b filter.Blur
	// err = b.Process("img/photo.jpeg", "img/photo_b.jpeg")
	// fmt.Println(err)

	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./img", "ouput", f)
	fmt.Println(t)
}
