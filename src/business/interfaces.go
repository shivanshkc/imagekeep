package business

import "gold/src/service"

var imageService iImageService = service.ImageService

type iImageService interface {
	Resize(image []byte, height int64, width int64) ([]byte, error)
}
