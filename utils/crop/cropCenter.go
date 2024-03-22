package crop

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

// crop or resize image to size (centered)
func CropSize(ori *image.Image, size [2]uint) (*image.Image, error) {
	croppedTragert := *ori
	var err error

	originalRatio := float64((*ori).Bounds().Dx()) / float64((*ori).Bounds().Dy())
	ratio := float64(size[0]) / float64(size[1])
	// if image and size are in same ratio, just resize image to smaller size
	if ratio == originalRatio {
		croppedTragert = resize.Resize(size[0], size[1], croppedTragert, resize.Lanczos3)
	} else {
		biggerSize := resizeToBig(
			[2]uint{uint((*ori).Bounds().Dx()), uint((*ori).Bounds().Dy())},
			size,
		)
		croppedTragert = resize.Resize(biggerSize[0], biggerSize[1], croppedTragert, resize.Lanczos3)
	}
	croppedTragert, err = cutter.Crop(croppedTragert, cutter.Config{
		Width:  int(size[0]),
		Height: int(size[1]),
		Mode:   cutter.Centered,
	})
	if err != nil {
		return nil, err
	}
	return &croppedTragert, nil
}

// crop or resize image to sizes (centered)
func CropSizes(ori *image.Image, sizes [][2]uint) ([]*image.Image, error) {
	images := make([]*image.Image, 0)
	for _, size := range sizes {
		croppedImage, err := CropSize(ori, size)
		if err != nil {
			return nil, err
		}
		images = append(images, croppedImage)
	}
	return images, nil
}
