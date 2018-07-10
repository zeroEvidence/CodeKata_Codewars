package kata_reducing_by_steps

import (
	"math"
	"math/big"
)

func Gcdi(x, y int) int {
	d := big.NewInt(0)
	u := big.NewInt(int64(math.Abs(float64(x))))
	v := big.NewInt(int64(math.Abs(float64(y))))

	return int(d.GCD(nil, nil, u, v).Int64())
}

func Som(x, y int) int {
	return x + y
}

func Maxi(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Mini(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Lcmu(x, y int) int {
	return int(math.Abs(float64(x*y)) / float64(Gcdi(x, y)))
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	var result []int

	for i, y := range arr {
		result = append(result, f(init, y))
		init = result[i]
	}

	return result
}
