package iteration

func Repeat(character string, repetitions int) string {
	result := ""

	for i := 0; i < repetitions; i++ {
		result += character
	}

	return result
}
