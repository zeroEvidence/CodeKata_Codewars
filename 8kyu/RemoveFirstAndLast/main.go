package kata_remove_first_and_last

func RemoveFirstAndLast(word string) string {
	return string(word[1 : len(word)-1])
}
