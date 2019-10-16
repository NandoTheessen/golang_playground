package iteration

func Repeat(char string, repetitionCount int) string {
	var result string
	for i := 0; i < repetitionCount; i++ {
		result += char
	}
	return result
}
