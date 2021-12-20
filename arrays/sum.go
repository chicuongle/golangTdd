package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(input ...[]int) []int {
	//numberOfInput := len(input)
	var result []int
	//:= make([]int, numberOfInput)
	for _, numbers := range input {
		result = append(result, Sum(numbers))
	}
	return result
}
