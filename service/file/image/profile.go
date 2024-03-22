package image

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func Profile(
	pic *multipart.FileHeader,
	userId string,
) (string, error) {
	fileName, err := generateFileName("/profile", fmt.Sprintf("/%s", userId))
	if err != nil {
		return "", err
	}
	return saveImage(pic, fileName)
}

func saveImage(
	pic *multipart.FileHeader,
	addr string,
) (string, error) {
	src, err := pic.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	extention := strings.ToLower(filepath.Ext(pic.Filename))

	switch pic.Header.Get("Content-Type") {
	case "image/jpeg", "image/png":
		{
			image, err := imageCrop(src)
			if err != nil {
				return "", err
			}
			addr += extention
			out, err := os.Create(addr)
			if err != nil {
				return "", err
			}
			defer out.Close()
			switch extention {
			case ".png":
				err = png.Encode(out, *image)
			case ".jpg", ".jpeg":
				err = jpeg.Encode(out, *image, nil)
			}
			if err != nil {
				return "", err
			}
			return addr, nil
		}
	default:
		{
			return "", fmt.Errorf("invalid content type")
		}
	}
}
