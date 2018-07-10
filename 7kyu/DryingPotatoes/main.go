package kata_drying_potatoes

import "math"

func Potatoes(p0, w0, p1 int) int {
	var dryWeightBefore = 100 - float64(p0)
	var dryWeightAfter = 100 - float64(p1)
	var weight = float64(w0)
	var weightAfter = weight * dryWeightBefore / dryWeightAfter

	return int(math.Floor(weightAfter))
}
