package crop

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func CropManual(
	ori *image.Image,
	size [2]uint,
	reletiveSize [2]int,
	reletivePosition [2][2]int,
) (*image.Image, error) {
	ratioDemension := [2]float64{
		float64((*ori).Bounds().Dx()) / float64(reletiveSize[0]),
		float64((*ori).Bounds().Dy()) / float64(reletiveSize[1]),
	}
	firstPoint := image.Point{
		X: int(float64(reletivePosition[0][0]) * ratioDemension[0]),
		Y: int(float64(reletivePosition[0][1]) * ratioDemension[1]),
	}
	lastPoint := image.Point{
		X: int(float64(reletivePosition[0][0]+reletivePosition[1][0]) * ratioDemension[0]),
		Y: int(float64(reletivePosition[0][1]+reletivePosition[1][1]) * ratioDemension[1]),
	}
	croppedImage, err := cutter.Crop(*ori, cutter.Config{
		Width:  int(lastPoint.X - firstPoint.X),
		Height: int(lastPoint.Y - firstPoint.Y),
		Anchor: image.Point{
			X: firstPoint.X,
			Y: firstPoint.Y,
		},
	})
	if err != nil {
		return nil, err
	}
	croppedTragert := resize.Resize(size[0], size[1], croppedImage, resize.Lanczos2)
	return &croppedTragert, nil
}

func CropManualOriginalImage(
	ori *image.Image,
	size [2]uint,
	reletiveSize [2]int,
	reletivePosition [2][2]int,
) (*image.Image, error) {
	ratioDemension := [2]float64{
		float64((*ori).Bounds().Dx()) / float64(reletiveSize[0]),
		float64((*ori).Bounds().Dy()) / float64(reletiveSize[1]),
	}
	firstPoint := image.Point{
		X: int(float64(reletivePosition[0][0]) * ratioDemension[0]),
		Y: int(float64(reletivePosition[0][1]) * ratioDemension[1]),
	}
	lastPoint := image.Point{
		X: int(float64(reletivePosition[0][0]+reletivePosition[1][0]) * ratioDemension[0]),
		Y: int(float64(reletivePosition[0][1]+reletivePosition[1][1]) * ratioDemension[1]),
	}
	croppedImage, err := cutter.Crop(*ori, cutter.Config{
		Width:  int(lastPoint.X - firstPoint.X),
		Height: int(lastPoint.Y - firstPoint.Y),
		Anchor: image.Point{
			X: firstPoint.X,
			Y: firstPoint.Y,
		},
	})
	if err != nil {
		return nil, err
	}
	return &croppedImage, nil
}
