package kata_drying_potatoes

import "math"

func Potatoes(p0, w0, p1 int) int {
	return int(math.Floor((float64(w0) * (100 - float64(p1))) / (100 - float64(p0))))
}
