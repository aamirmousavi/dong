package crop

func resizeToBig(ori, dest [2]uint) [2]uint {
	proccesed := [2]uint{}
	originRatio := float64(ori[0]) / float64(ori[1])
	destRatio := float64(dest[0]) / float64(dest[1])
	if originRatio > 1 {
		if destRatio > 1 {
			if originRatio > destRatio {
				proccesed[0] = uint(float64(dest[1]) * originRatio)
				proccesed[1] = dest[1]
			} else {
				proccesed[0] = dest[0]
				proccesed[1] = uint(float64(dest[0]) / originRatio)
			}
		} else {
			proccesed[0] = uint(float64(dest[1]) * originRatio)
			proccesed[1] = dest[1]
		}
	} else if originRatio < 1 {
		if destRatio > 1 {
			proccesed[0] = dest[0]
			proccesed[1] = uint(float64(dest[0]) / originRatio)
		} else {
			if originRatio > destRatio {
				proccesed[0] = uint(float64(dest[1]) * originRatio)
				proccesed[1] = dest[1]
			} else {
				proccesed[0] = dest[0]
				proccesed[1] = uint(float64(dest[0]) / originRatio)
			}
		}
	} else {
		if destRatio > 1 {
			proccesed[0] = dest[0]
			proccesed[1] = uint(float64(dest[0]) / originRatio)
		} else {
			proccesed[0] = uint(float64(dest[1]) * originRatio)
			proccesed[1] = dest[1]
		}
	}
	return proccesed
}
