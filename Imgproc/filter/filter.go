package filter

import (
	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Grayscale struct{}
type Blur struct{}

func (g Grayscale) Process(srcPath, dstPath string) error {
	// Open image
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Create a grayscale version of the image
	dst := imaging.Grayscale(src)

	// Save the resulting image as JPEG
	err = imaging.Save(dst, dstPath)
	if err != nil {
		return err
	}
	return nil
}

func (b Blur) Process(srcPath, dstPath string) error {
	// Open image
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Create a blurred version of the image
	dst := imaging.Blur(src, 5)

	// Save the resulting image as JPEG
	err = imaging.Save(dst, dstPath)
	if err != nil {
		return err
	}
	return nil
}
