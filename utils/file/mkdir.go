package file

import "os"

const (
	_MODE = 0777
	// _MODE = 0755
)

func MkdirIfNotExsits(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Mkdir(path)
	}
	return nil
}

func Mkdir(path string) error {
	return os.Mkdir(path, _MODE)
}
