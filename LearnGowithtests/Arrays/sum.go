package arrays

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	result := make([]int, length)
	for i, slice := range numbersToSum {
		if len(slice) <= 1 {
			if len(slice) == 1 {
				result[i] = slice[0]
			} else {
				result[i] = 0
			}
		} else {
			tail := slice[1:]
			result[i] = Sum(tail)
		}
	}
	return result
}
