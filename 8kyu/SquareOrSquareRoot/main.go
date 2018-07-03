package kata_square_or_square_root

import "math"

func SquareOrSquareRoot(arr []int) []int {
	var res []int

	for _, v := range arr {
		f := float64(v)
		sqrtf := math.Sqrt(f)

		if math.Mod(f, sqrtf) == 0.0 {
			res = append(res, int(sqrtf))
			continue
		}

		res = append(res, int(f*f))
	}

	return res
}
