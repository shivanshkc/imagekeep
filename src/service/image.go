package service

// ImageService : Implements the Image Service methods.
var ImageService = &imageService{}

type imageService struct{}

func (i *imageService) Resize(image []byte, height int64, width int64) ([]byte, error) {
	return nil, nil
}
