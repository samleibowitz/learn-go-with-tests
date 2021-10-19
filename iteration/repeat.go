package iteration

// Repeat takes a string and repeats it the number of times specified in the second argument
func Repeat(character string, repetitions int) string {
	result := ""

	for i := 0; i < repetitions; i++ {
		result += character
	}

	return result
}
