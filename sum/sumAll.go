package sum

func SumAll(numbersToSum ...[]int) []int {
	lenNumbers := len(numbersToSum)
	sums := make([]int, lenNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}