package arrays

/**
When you have an array, it is very common to have to iterate over them.
So let's use our new-found knowledge of for to make a Sum function. Sum will take an array of numbers and return the total.
**/
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(input ...[]int) []int {
	var result []int
	for _, numbers := range input {
		result = append(result, Sum(numbers))
	}
	return result
}

/*
Our next requirement is to change SumAll to SumAllTails, where it will calculate the totals of the "tails" of each slice.
The tail of a collection is all items in the collection except the first one (the "head").
*/
func SumAllTails(input ...[]int) []int {
	sumAll := SumAll(input...)
	for i, numbers := range input {
		if(len(numbers)>0){
			sumAll[i] -= numbers[0]
		}
	}
	return sumAll
}
