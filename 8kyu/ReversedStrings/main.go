package kata_reversed_strings

func ReversedStrings(word string) string {
	var reversed = make([]byte, len(word))

	for i, j := 0, len(word)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = word[j], word[i]
	}

	return string(reversed[:])
}
