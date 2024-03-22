package image

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	utils_file "github.com/aamirmousavi/dong/utils/file"
)

func generateFileName(
	addr string,
	extra ...string,
) (string, error) {
	full := _STORAGE + addr
	for _, ex := range extra {
		full += ex
	}
	folders := strings.Split(full, "/")
	for i := range folders {
		if i > 0 {
			folders[i] = folders[i-1] + folders[i]
		}
		if err := utils_file.MkdirIfNotExsits(folders[i]); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s/%s", full, randomName(1000, 9999)), nil
}

func randomName(min, max int) string {
	rand.Seed(time.Now().Unix())
	n := min + rand.Intn(max-min)
	return strconv.Itoa(int(time.Now().UnixNano())) +
		strconv.Itoa(n)
}
