package kata_sum_of_positive

func SumOfPositive(values []int) int {
	var sum = 0

	for i := 0; i < len(values); i++ {
		if values[i] > 0 {
			sum += values[i]
		}
	}

	return sum
}
