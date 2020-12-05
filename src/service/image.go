package service

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
)

// ImageService : Implements the Image Service methods.
var ImageService = &imageService{}

type imageService struct{}

func (i *imageService) Resize(imageBytes []byte, height int64, width int64) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	resizedImage := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	buf := bytes.NewBuffer(nil)
	if err := jpeg.Encode(buf, resizedImage, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
