package main

import (
	"fmt"
	"time"

	"github.com/fabienbellanger/Udemy_Golang-formation-complete/Imgproc/filter"
	"github.com/fabienbellanger/Udemy_Golang-formation-complete/Imgproc/task"
)

func main() {
	// var g filter.Grayscale
	// err := g.Process("img/photo.jpeg", "img/photo_gs.jpeg")
	// fmt.Println(err)

	// var b filter.Blur
	// err = b.Process("img/photo.jpeg", "img/photo_b.jpeg")
	// fmt.Println(err)

	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)

	start := time.Now()
	t.Process()
	elapsted := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsted)
}
