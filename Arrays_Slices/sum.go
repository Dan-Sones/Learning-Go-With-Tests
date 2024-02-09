package Arrays_Slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//func SumAll(numbersToSum ...[]int) []int {
//	lengthOfArray := len(numbersToSum)
//
//	// make allows us to create a new slice
//	// Where the second arguement is the initial length
//	sums := make([]int, lengthOfArray)
//
//	for i, numbers := range numbersToSum {
//		sums[i] = Sum(numbers)
//	}
//
//	return sums
//
//}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int                         // create an empty slice
	for _, numbers := range numbersToSum { // For each array passed in

		if len(numbers) == 0 { // if it is empty, then just append 0, there is no point doing a calc
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail)) // append the result of Sum() to the sums array
		}
	}
	return sums
}
