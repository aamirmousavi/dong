package image

import (
	"image"
	"mime/multipart"

	"github.com/aamirmousavi/dong/utils/crop"
)

func imageCrop(src multipart.File) (*image.Image, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}
	return crop.CropSize(&img, [2]uint{360, 360})
}
