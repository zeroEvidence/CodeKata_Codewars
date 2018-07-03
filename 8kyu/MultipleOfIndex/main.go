package kata_multiple_of_index

func MultipleOfIndex(ints []int) []int {
	var result []int

	for i, v := range ints {
		if i != 0 && v%i == 0 {
			result = append(result, v)
		}
	}

	return result
}
